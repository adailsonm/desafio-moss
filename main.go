package main

import (
	"github.com/adailsonm/desafio-moss/api/handlers"
	"github.com/adailsonm/desafio-moss/core/order"
	"github.com/adailsonm/desafio-moss/core/pizza"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	db, err := gorm.Open(postgres.Open("postgres://postgres:desafiomoss@db/postgres"))
	if err != nil {
		log.Fatal(err)
	}
	orderService := order.NewService(db)
	pizzaService := pizza.NewService(db)

	//handlers
	handlers.MakeOrderHandlers(r, orderService)
	handlers.MakePizzzaHandlers(r, pizzaService)

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

