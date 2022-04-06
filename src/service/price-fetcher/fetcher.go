package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"

	"github.com/sirupsen/logrus"
	gecko "github.com/superoo7/go-gecko/v3"
)

//FetcherSrv
type FetcherSrv struct {
	logger    *logrus.Entry
	storage   *storage.DataBase
	AllTokens []string
}

//CreateNewFetcherSrv
func CreateNewFetcherSrv(logger *logrus.Logger, db *storage.DataBase, cfg *models.FetcherConfig) *FetcherSrv {
	return &FetcherSrv{
		logger:    logger.WithField("layer", "fetcher"),
		storage:   db,
		AllTokens: cfg.AllTokens,
	}
}

func (f *FetcherSrv) Run() {
	f.logger.Infoln("Fetcher srv started")
	go f.collector()
}

func (f *FetcherSrv) collector() {
	for {
		f.getPriceInfo()
		time.Sleep(30 * time.Second)
	}
}

func (f *FetcherSrv) getPriceInfo() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg := gecko.NewClient(httpClient)

	ids := f.AllTokens
	vc := []string{"usd"}

	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		f.logger.Warn("fetch timeout exceeded")
		return
	}

	priceLog := make([]*storage.PriceLog, len(ids))
	for index, name := range ids {
		price := (*sp)[name]["usd"]

		if name == "weowns" {
			resp, err := f.makeReq("https://nomics.com/data/currencies-ticker?filter=any&interval=1d&quote-currency=USD&symbols=WEOWNS", httpClient)
			if err != nil {
				logrus.Warnf("fetch WEOWNS price error = %s", err)
				priceLog[index] = &storage.PriceLog{}

				continue
			} else {
				price64, _ := strconv.ParseFloat((*resp)["items"].([]interface{})[0].(map[string]interface{})["price"].(string), 32)
				price = float32(price64)
			}
		}

		priceLog[index] = &storage.PriceLog{
			Name:       name,
			Price:      fmt.Sprintf("%f", price),
			UpdateTime: time.Now().Unix(),
		}
	}
	f.storage.SavePriceInformation(priceLog)
	f.logger.Infoln("new prices fetched at", time.Now().Unix())
}

// MakeReq HTTP request helper
func (f *FetcherSrv) makeReq(url string, c *http.Client) (*map[string]interface{}, error) {
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
func (f *FetcherSrv) doReq(req *http.Request, client *http.Client) ([]byte, error) {
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
