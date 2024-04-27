package controller

import (
	"fmt"
	"net/http"
	"test/domain"
	"test/repository"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	http *http.Server
	db   *repository.PostgresHandler
}

func NewServerHttp(settings domain.HTTPSettings, db *repository.PostgresHandler) *HttpServer {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", settings.Host, settings.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	server := &HttpServer{
		http: srv,
		db:   db,
	}

	server.UrlAssing(r)

	return server
}

func (srv *HttpServer) Start() error {
	return srv.http.ListenAndServe()
}
