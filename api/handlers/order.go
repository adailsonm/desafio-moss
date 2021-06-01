package handlers

import (
	"encoding/json"
	"github.com/adailsonm/desafio-moss/core/order"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MakeOrderHandlers(r *mux.Router, service *order.Service)  {
	r.Handle("/v1/order", getAllOrder(service)).Methods("GET", "OPTIONS")
	r.Handle("/v1/beer", storeOrder(service)).Methods("POST", "OPTIONS")
}

func getAllOrder(service order.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			getAllOrderJSON(w, service)
		default:
			log.Fatal("Formato inexistente")
		}

	})
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
func getAllOrderJSON(w http.ResponseWriter, service order.UseCase) {
	w.Header().Set("Content-Type", "application/json")
	all, err := service.GetAll()
	if err != nil {
		w.Write(formatJSONError(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		w.Write(formatJSONError("Erro convertendo em JSON"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
