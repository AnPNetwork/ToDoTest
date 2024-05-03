package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s HttpServer) UrlAssing(r *mux.Router) {
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./template/res/")))
	r.PathPrefix("/static/").Handler(staticHandler)

	r.HandleFunc("/", s.MainHandler)
	r.HandleFunc("/api", s.GetHandler).Methods("GET")
	r.HandleFunc("/api", s.PostHandler).Methods("POST")
	r.HandleFunc("/api", s.PutHandler).Methods("PUT")
	r.HandleFunc("/api", s.DeleteHandler).Methods("DELETE")
	s.http.Handler = r
}
