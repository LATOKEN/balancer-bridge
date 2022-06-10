package farmer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"
	"github.com/latoken/bridge-balancer-service/src/service/workers"

	"github.com/sirupsen/logrus"
)

//FarmerSrv
type FarmerSrv struct {
	logger    *logrus.Entry
	storage   *storage.DataBase
	Workers   map[string]workers.IWorker
	farmerCfg *models.FarmerConfig
	FarmCfgs  map[string]*models.FarmConfig
}

//CreateNewFarmerSrv
func CreateNewFarmerSrv(logger *logrus.Logger, db *storage.DataBase, workers map[string]workers.IWorker,
	farmerCfg *models.FarmerConfig, farmCfgs map[string]*models.FarmConfig) *FarmerSrv {
	return &FarmerSrv{
		logger:    logger.WithField("layer", "farmer"),
		storage:   db,
		Workers:   workers,
		farmerCfg: farmerCfg,
		FarmCfgs:  farmCfgs,
	}
}

func (f *FarmerSrv) Run() {
	f.logger.Infoln("Fetcher srv started")
	go f.collector()
}

func (f *FarmerSrv) collector() {
	for {
		f.getFarmInfo()
		time.Sleep(60 * time.Second)
	}
}

func (f *FarmerSrv) getFarmInfo() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	apy, _ := f.makeReq(f.farmerCfg.ApyURL, httpClient)
	lps, _ := f.makeReq(f.farmerCfg.LpsURL, httpClient)
	prices, _ := f.makeReq(f.farmerCfg.PricesURL, httpClient)

	for _, farmCfg := range f.FarmCfgs {
		worker := f.Workers[farmCfg.ChainName]

		balance, pricePerFullShare, err := worker.GetVaultInfo(farmCfg.VaultAddress)
		if err != nil {
			logrus.Warnf("fetch vault info error = %s", err)
			return
		}
		tvl, _ := strconv.ParseFloat(balance.String(), 64)

		if farmCfg.Type == "SINGLE" {
			f.storage.UpsertFarm(&storage.Farm{
				ID:                 farmCfg.ID,
				TVL:                fmt.Sprintf("%f", (tvl*(*prices)[farmCfg.Oracle].(float64))/math.Pow(10, 18)),
				APY:                fmt.Sprintf("%f", (*apy)[farmCfg.FarmId].(map[string]interface{})["totalApy"].(float64)),
				PricePerFullShare0: pricePerFullShare.String(),
				UpdateTime:         time.Now().Unix(),
			})
		} else {
			totalySupply, reserve0, reserve1, err := worker.GetPairInfo(farmCfg.PairAddress)
			if err != nil {
				logrus.Warnf("fetch pair info error = %s", err)
				return
			}
			var bigInt big.Int

			f.storage.UpsertFarm(&storage.Farm{
				ID:                 farmCfg.ID,
				TVL:                fmt.Sprintf("%f", (tvl*(*lps)[farmCfg.Oracle].(float64))/math.Pow(10, 18)),
				APY:                fmt.Sprintf("%f", (*apy)[farmCfg.FarmId].(map[string]interface{})["totalApy"].(float64)),
				PricePerFullShare0: bigInt.Div(bigInt.Mul(pricePerFullShare, reserve0), totalySupply).String(),
				PricePerFullShare1: bigInt.Div(bigInt.Mul(pricePerFullShare, reserve1), totalySupply).String(),
				UpdateTime:         time.Now().Unix(),
			})
		}
	}
}

// MakeReq HTTP request helper
func (f *FarmerSrv) makeReq(url string, c *http.Client) (*map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := f.doReq(req, c)
	if err != nil {
		return nil, err
	}

	t := make(map[string]interface{})
	er := json.Unmarshal(resp, &t)
	if er != nil {
		return nil, er
	}

	return &t, err
}

// helper
// doReq HTTP client
func (f *FarmerSrv) doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
