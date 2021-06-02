package handlers

import (
	"encoding/json"
	"github.com/adailsonm/desafio-moss/core/order"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeOrderHandlers(r *mux.Router, service *order.Service)  {
	r.Handle("/v1/order", storeOrder(service)).Methods("POST", "OPTIONS")
}

func storeOrder(service order.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var o order.Order
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		err = service.Store(&o)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}