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
func (r *BridgeSRV) emitFeeTransfer(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusFeeTransferInitConfrimed, storage.EventStatusFeeTransferSentFailed})
		for _, event := range events {
			if event.Status == storage.EventStatusFeeTransferInitConfrimed {
				r.logger.Infoln("attempting to send fee transfer")
				if _, err := r.sendFeeTransfer(worker, event); err != nil {
					r.logger.Errorf("fee transfer failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypeFeeTransfer,
					storage.EventStatusFeeTransferInitConfrimed, storage.EventStatusFeeTransferSentFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendFeeTransfer(worker workers.IWorker, event *storage.Event) (txHash string, err error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypeFeeTransfer,
		CreateTime: time.Now().Unix(),
	}
	// for BSC-USDT decimal conversion
	tetherRID := r.storage.FetchResourceIDByName("tether").ID

	bscDestID := ""
	if worker, ok := r.Workers["BSC"]; ok {
		bscDestID = worker.GetDestinationID()
	}
	htDestID := ""
	if worker, ok := r.Workers["HT"]; ok {
		htDestID = worker.GetDestinationID()
	}

	var amount string
	if (event.OriginChainID == bscDestID || event.OriginChainID == htDestID) && event.ResourceID == tetherRID {
		amount = utils.Convertto6Decimals(event.InAmount)
	} else if (event.DestinationChainID == bscDestID || event.DestinationChainID == htDestID) && event.ResourceID == tetherRID {
		amount = utils.Convertto18Decimals(event.InAmount)
	} else {
		amount = event.InAmount
	}

	r.logger.Infof("Fee Transfer parameters: outAmount(%s) | recipient(%s) | chainID(%s)\n",
		amount, event.ReceiverAddr, worker.GetChainName())
	txHash, err = worker.TransferExtraFee(utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID),
		event.DepositNonce, utils.StringToBytes32(event.ResourceID), event.ReceiverAddr, amount)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusNotFound
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
