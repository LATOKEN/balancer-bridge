package rlr

import (
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
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
func (r *BridgeSRV) GetPriceOfToken(name string) (price string) {
	priceLog := r.storage.GetPriceLog(name)
	return priceLog.Price
}
