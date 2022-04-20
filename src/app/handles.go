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
			"/status",
			"signature/{amount}/{recipientAddress}/{destinationChainID}",
			// "/failed_swaps/{page}",
			// "/resend_tx/{id}",
			// "/set_mode/{mode}",
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

	if msg == "" {
		a.logger.Errorf("Empty request(price/{token})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
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
