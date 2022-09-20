package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/latoken/bridge-balancer-service/src/common"
)

// Endpoints ...
func (a *App) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/price/{token}",
			"/price/{resourceID}?v=2",
			"/status",
			"/signature/{amount}/{recipientAddress}/{destinationChainID}",
			"/swapstatus/{txHash}",
		},
	}
	common.ResponJSON(w, http.StatusOK, endpoints)
}

func (a *App) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := a.relayer.StatusOfWorkers()
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.ResponJSON(w, http.StatusOK, status)
}

func (a *App) PriceHandler(w http.ResponseWriter, r *http.Request) {
	// msg := models.PriceConfig{
	// 	Name: mux.Vars(r)["token"],
	// }
	msg := mux.Vars(r)["token"]
	v := r.URL.Query().Get("v")

	if msg == "" {
		a.logger.Errorf("Empty request(price)")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}

	if v == "2" {
		rID := a.relayer.Storage.FetchResourceID(msg)
		msg = rID.Name
	}
	price, err := a.relayer.GetPriceOfToken(msg)
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.ResponJSON(w, http.StatusOK, price)
}

func (a *App) SignatureHandler(w http.ResponseWriter, r *http.Request) {
	amount := mux.Vars(r)["amount"]
	recipientAddress := mux.Vars(r)["recipientAddress"]
	destinationChainID := mux.Vars(r)["destinationChainID"]

	if amount == "" || recipientAddress == "" || destinationChainID == "" {
		a.logger.Errorf("Empty Request (/signature/{amount}/{recipient}/{destinationChainID})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}
	signature, err := a.relayer.CreateSignature(amount, recipientAddress, destinationChainID)
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.ResponJSON(w, http.StatusOK, signature)
}

func (a *App) SwapStatusHandler(w http.ResponseWriter, r *http.Request) {
	txHash := mux.Vars(r)["txHash"]

	if txHash == "" {
		a.logger.Errorf("Empty Request (/swapstatus/{txHash})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}
	status, err := a.relayer.GetSwapStatusByTxHash(txHash)
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.ResponJSON(w, http.StatusOK, status)
}
