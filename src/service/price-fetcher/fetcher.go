package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"

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
		priceLog[index] = &storage.PriceLog{
			Name:       name,
			Price:      fmt.Sprintf("%f", (*sp)[name]["usd"]),
			UpdateTime: time.Now().Unix(),
		}
	}
	f.storage.SavePriceInformation(priceLog)
	f.logger.Infoln("new prices fetched at", time.Now().Unix())
}
