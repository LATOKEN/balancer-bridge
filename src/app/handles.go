package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/common"
)

const numPerPage = 100

// Endpoints ...
func (a *App) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/price/{token}",
			"/status",
			"signature/{amount}/{recipientAddress}/{originChainID}",
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
	originChainID := mux.Vars(r)["originChainID"]

	if amount == "" || recipientAddress == "" || originChainID == "" {
		a.logger.Errorf("Empty Request (/signature/{amount}/{recipient}/{originChainID})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}
	messagehash, signature, err := a.relayer.CreateSignature(amount, recipientAddress, originChainID)
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := SigAndHash{
		Hash:      messagehash,
		Signature: signature,
	}

	common.ResponJSON(w, http.StatusOK, res)
}
