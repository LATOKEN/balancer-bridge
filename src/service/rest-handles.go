package rlr

import (
	"github.com/latoken/bridge-balancer-service/src/models"
)

// Status ...
func (r *BridgeSRV) StatusOfWorkers() (map[string]*models.WorkerStatus, error) {
	// get blockchain heights from workers and from database
	workers := make(map[string]*models.WorkerStatus)
	for _, w := range r.Workers {
		status, err := w.GetStatus()
		if err != nil {
			r.logger.Errorf("While get status for worker = %s, err = %v", w.GetChainName(), err)
			return nil, err
		}
		workers[w.GetChainName()] = status
	}

	for name, w := range workers {
		blocks := r.storage.GetCurrentBlockLog(name)
		w.SyncHeight = blocks.Height
	}

	return workers, nil
}

//GetPriceOfToken
func (r *BridgeSRV) GetPriceOfToken(name string) (string, error) {
	priceLog, err := r.storage.GetPriceLog(name)
	if err != nil {
		return "", err
	}
	return priceLog.Price, nil
}

//Create signature and hash
func (r *BridgeSRV) CreateSignature(amount, recipientAddress, destinationChainID string) (signature string, err error) {
	messageHash, err := r.laWorker.CreateMessageHash(amount, recipientAddress, destinationChainID)
	signature, err = r.laWorker.CreateSignature(messageHash, destinationChainID)
	if err != nil {
		return "", err
	}
	return signature, nil
}
