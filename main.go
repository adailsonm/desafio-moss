package main

import (
	"database/sql"
	"github.com/adailsonm/desafio-moss/api/handlers"
	"github.com/adailsonm/desafio-moss/core/order"
	"github.com/adailsonm/desafio-moss/core/pizza"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	db, err := sql.Open("postgres","postgres://postgres:desafiomoss@db/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

