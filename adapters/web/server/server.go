package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/adapters/web/handler"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func CreateNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	// Roteador que ira manipular todas as nossas requests
	r := mux.NewRouter()
	// Middleware que irá logar todas as nossas requests
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
