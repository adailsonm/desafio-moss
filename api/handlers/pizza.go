package handlers

import (
	"encoding/json"
	"github.com/adailsonm/desafio-moss/core/pizza"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func MakePizzzaHandlers(r *mux.Router, service *pizza.Service) {
	r.Handle("/v1/pizza", getAllPizza(service)).Methods("GET", "OPTIONS")
	r.Handle("/v1/pizza/find", getPizzaByName(service)).Methods("GET", "OPTIONS")
	r.Handle("/v1/pizza", storePizza(service)).Methods("POST", "OPTIONS")
	r.Handle("/v1/pizza/{id:[0-9]+}", UpdatePizza(service)).Methods("PUT", "OPTIONS")
	r.Handle("/v1/pizza/{id:[0-9]+}", DeletePizza(service)).Methods("DELETE", "OPTIONS")

}

func getAllPizza(service pizza.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			getAllPizzaJSON(w, service)
		default:
			log.Fatal("Formato inexistente")
		}

	})
}

func DeletePizza(service pizza.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			DeletePizzaJSON(r, w, service)
		default:
			log.Fatal("Formato inexistente")
		}

	})
}

func UpdatePizza(service pizza.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		var p pizza.Pizza
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.Update(id, &p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func DeletePizzaJSON(r *http.Request, w http.ResponseWriter, service pizza.UseCase) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.Write(formatJSONError(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = service.Remove(id)
	if err != nil {
		w.Write(formatJSONError(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func storePizza(service pizza.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p pizza.Pizza
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.Store(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func getPizzaByName(service pizza.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			getPizzaByNameJSON(r, w, service)
		default:
			log.Fatal("Formato inexistente")
		}

	})
}

func getPizzaByNameJSON(r *http.Request, w http.ResponseWriter, service pizza.UseCase) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	all, err := service.GetByName(name)
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

func getAllPizzaJSON(w http.ResponseWriter, service pizza.UseCase) {
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
