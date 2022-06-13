package rlr

import (
	"fmt"
	"math"
	"strconv"

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

// GetFarmsInfo
func (r *BridgeSRV) GetUserFarmBalance(farmId, userBalance string) (map[string]string, error) {
	userFarmBalance := make(map[string]string)

	farmCfg := r.Farmer.FarmCfgs[farmId]

	farmInfo, err := r.storage.GetFarm(farmId)
	if err != nil {
		return nil, err
	}

	pricePerFullShareFloat0, err := strconv.ParseFloat(farmInfo.PricePerFullShare0, 64)
	if err != nil {
		return nil, err
	}

	userBalanceFloat, err := strconv.ParseFloat(userBalance, 64)
	if err != nil {
		return nil, err
	}
	userBalanceFloat -= userBalanceFloat * farmCfg.WithdrawalFee

	userFarmBalance[farmCfg.Token0] = fmt.Sprintf("%d", int(pricePerFullShareFloat0*userBalanceFloat/math.Pow(10, 18)))
	if farmCfg.Type != "SINGLE" {
		pricePerFullShareFloat1, err := strconv.ParseFloat(farmInfo.PricePerFullShare1, 64)
		if err != nil {
			return nil, err
		}

		userFarmBalance[farmCfg.Token1] = fmt.Sprintf("%d", int(pricePerFullShareFloat1*userBalanceFloat/math.Pow(10, 18)))
	}

	return userFarmBalance, nil
}

// GetFarmsInfo
func (r *BridgeSRV) GetFarmsInfo() ([]*models.FarmInfo, error) {
	farmCfgs := r.Farmer.FarmCfgs
	farmInfos := make([]*models.FarmInfo, 0, len(farmCfgs))

	for _, farmCfg := range farmCfgs {
		farmInfo, err := r.storage.GetFarm(farmCfg.ID)
		if err != nil {
			return nil, err
		}

		farmInfos = append(farmInfos, &models.FarmInfo{
			ID:                  farmCfg.ID,
			FarmId:              farmCfg.FarmId,
			TVL:                 farmInfo.TVL,
			APY:                 farmInfo.APY,
			ChainId:             farmCfg.ChainId,
			Name:                farmCfg.Name,
			Protocol:            farmCfg.Protocol,
			DepositToken:        farmCfg.DepositToken,
			Logo0:               farmCfg.Logo0,
			Logo1:               farmCfg.Logo1,
			DepositResourceID:   farmCfg.DepositResourceID,
			WithdrawResourceID:  farmCfg.WithdrawResourceID,
			Gas:                 farmCfg.Gas,
			WrappedTokenAddress: farmCfg.WrappedTokenAddress,
		})
	}

	return farmInfos, nil
}
