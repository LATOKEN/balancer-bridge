package rlr

import (
	"fmt"
	"time"

	"github.com/latoken/bridge-balancer-service/src/service/storage"
	workers "github.com/latoken/bridge-balancer-service/src/service/workers"
	"github.com/latoken/bridge-balancer-service/src/service/workers/utils"
)

// !!! TODO !!!

// emitRegistreted ...
func (r *BridgeSRV) emitDepositRedeemAnchor(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.TxType{storage.TxTypeTokenTransfer}, []storage.EventStatus{storage.EventStatusFeeTransferInitConfrimed, storage.EventStatusFeeTransferSentFailed})
		for _, event := range events {
			if event.Status == storage.EventStatusFeeTransferInitConfrimed &&
				worker.GetDestinationID() == event.DestinationChainID {
				r.logger.Infoln("attempting to send depositRedeemAnchor")
				if _, err := r.sendDepositRedeemAnchor(worker, event); err != nil {
					r.logger.Errorf("depositRedeemAnchor failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypeTokenTransfer,
					storage.EventStatusFeeTransferInitConfrimed, storage.EventStatusFeeTransferSentFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendDepositRedeemAnchor(worker workers.IWorker, event *storage.Event) (txHash string, err error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypeTokenTransfer,
		CreateTime: time.Now().Unix(),
	}

	workerCfg := worker.GetWorkerConfig()
	if event.ReceiverAddr == workerCfg.USTContractAddress.String() {
		redeemTx := r.storage.GetRedeemedAmount(event.InAmount)

		swap := r.storage.GetConfirmedSwapByResourceIDsAndAmount()

	}

	r.logger.Infof("Fee Transfer parameters: outAmount(%s) | recipient(%s) | chainID(%s)\n",
		amount, event.ReceiverAddr, worker.GetChainName())
	txHash, err = worker.TransferExtraFee(utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID),
		event.DepositNonce, utils.StringToBytes32(event.ResourceID), event.ReceiverAddr, amount)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusFeeTransferSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send fee transfer tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusFeeTransferSent)
	r.logger.Infof("send fee transfer tx success | recipient=%s, tx_hash=%s", event.ReceiverAddr, txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil

}
