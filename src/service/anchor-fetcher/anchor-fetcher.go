package anchor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"
	"github.com/sirupsen/logrus"
)

//AnchorFetcherSrv
type AnchorFetcherSrv struct {
	logger       *logrus.Entry
	storage      *storage.DataBase
	anchorFetCfg *models.AnchorFetcherConfig
}

//CreateNewAnchorFetcherSrv
func CreateNewAnchorFetcherSrv(logger *logrus.Logger, db *storage.DataBase, anchorFetCfg *models.AnchorFetcherConfig) *AnchorFetcherSrv {
	return &AnchorFetcherSrv{
		logger:       logger.WithField("layer", "anchor-fetcher"),
		storage:      db,
		anchorFetCfg: anchorFetCfg,
	}
}

func (f *AnchorFetcherSrv) Run() {
	f.logger.Infoln("Anchor fetcher srv started")
	go f.collector()
}

func (f *AnchorFetcherSrv) collector() {
	for {
		f.getNewTxs()
		time.Sleep(60 * time.Second)
	}
}

func (f *AnchorFetcherSrv) getNewTxs() {
	lastTxId := f.storage.GetLastTxId()
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	txs, err := f.makeReq(f.anchorFetCfg.TxListURL, httpClient)
	if err != nil {
		logrus.Warnf("Anchor fetch error = %s", err)
		return
	}

	for _, tx := range (*txs)["txs"].([]interface{}) {
		tx := tx.(map[string]interface{})

		txId := tx["id"].(int64)
		if txId <= lastTxId {
			break
		}

		for _, event := range tx["logs"].([]interface{})[0].(map[string]interface{})["events"].([]interface{}) {
			event := event.(map[string]interface{})

			if event["type"].(string) == "from_contract" {
				attributes := event["attributes"].([]interface{})

				for i, attribute := range attributes {
					attribute := attribute.(map[string]interface{})
					txType := attribute["value"].(string)

					if txType == "redeem_stable" || txType == "deposit_stable" {
						f.storage.SaveAnchorTx(&storage.AnchorTx{
							ID:         txId,
							Type:       txType,
							AUSTAmount: attributes[i+2].(map[string]interface{})["value"].(string),
							USTAmount:  attributes[i+3].(map[string]interface{})["value"].(string),
						})

						break
					}
				}

				break
			}
		}
	}

	f.logger.Infoln("New anchor txs fetched")
}

// MakeReq HTTP request helper
func (f *AnchorFetcherSrv) makeReq(url string, c *http.Client) (*map[string]interface{}, error) {
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
func (f *AnchorFetcherSrv) doReq(req *http.Request, client *http.Client) ([]byte, error) {
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
