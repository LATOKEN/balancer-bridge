package rlr

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
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
					storage.EventStatusFeeTransferConfirmed, storage.EventStatusFeeTransferInitConfrimed)
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
	//convert other native to corresponding latoken amount
	lpStr, _ := r.GetPriceOfToken("latoken")
	latokenPrice, _ := strconv.ParseFloat(lpStr, 32)
	swappedToken := r.storage.FetchResourceID(strings.ToLower(event.ResourceID))
	spStr, _ := r.GetPriceOfToken(swappedToken.Name)
	otherChainPrice, _ := strconv.ParseFloat(spStr, 32)
	inamount, _ := strconv.ParseFloat(event.InAmount, 32)
	event.OutAmount = uint64(inamount * otherChainPrice / latokenPrice)

	r.logger.Infof("Fee Transfer parameters: outAmount(%d) | recipient(%s) | chainID(%s)\n",
		event.OutAmount, event.ReceiverAddr, worker.GetChainName())
	txHash, err = worker.TransferExtraFee(event.ReceiverAddr, event.OutAmount)
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
