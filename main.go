package desafio_moss

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	//middlewares - código que vai ser executado em todas as requests
	//aqui podemos colocar logs, inclusão e validação de cabeçalhos, etc

	//handlers

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

