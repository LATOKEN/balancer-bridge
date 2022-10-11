package blcr_srv

import (
	"fmt"
	"time"

	"github.com/latoken/bridge-balancer-service/src/service/storage"
	"github.com/latoken/bridge-balancer-service/src/service/workers"
	"github.com/latoken/bridge-balancer-service/src/service/workers/utils"
)

func (r *BridgeSRV) emitFeeReversal(wrkr workers.IWorker) {
	for {
		events := r.Storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusFeeTransferFailed, storage.EventStatusFeeReversalInit, storage.EventStatusFeeReversalSentFailed, storage.EventStatusFeeReversalSent})
		for _, event := range events {
			if event.Status == storage.EventStatusFeeReversalInit && event.OriginChainID == wrkr.GetDestinationID() {
				r.logger.Infoln("attempting to send fee reversal")
				if _, err := r.sendFeeReversal(wrkr, event); err != nil {
					r.logger.Errorf("fee reversal failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypeFeeReversal,
					storage.EventStatusFeeReversalInit, storage.EventStatusFeeReversalFailed, storage.EventStatusFeeReversalSentConfirmed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

func (r *BridgeSRV) sendFeeReversal(wrkr workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      wrkr.GetChainName(),
		Type:       storage.TxTypeFeeReversal,
		SwapID:     event.SwapID,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("Fee Reversal parameters: outAmount(%s) | recipient(%s)\n",
		event.InAmount, event.ReceiverAddr)
	txHash, nonce, err := wrkr.ReversalTx(utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID),
		event.DepositNonce, utils.StringToBytes32(event.ResourceID), event.ReceiverAddr, event.InAmount, event.Data)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusNotFound
		r.Storage.UpdateEventStatus(event, storage.EventStatusFeeReversalSentFailed)
		r.Storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send fee reversal tx: %w", err)
	}
	txSent.TxHash = txHash
	txSent.Nonce = nonce
	r.Storage.UpdateEventStatus(event, storage.EventStatusFeeReversalSent)
	r.logger.Infof("send fee reversal tx success | recipient=%s, tx_hash=%s", event.ReceiverAddr, txSent.TxHash)
	// create new tx(claimed)
	r.Storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
