package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s HttpServer) UrlAssing(r *mux.Router) {
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./template/res/")))
	r.HandleFunc("/", s.HomeHandler)
	r.HandleFunc("/change", s.ChangeGetHandler).Methods("GET")
	r.HandleFunc("/change", s.ChangePostHandler).Methods("POST")
	r.HandleFunc("/change", s.ChangePutHandler).Methods("PUT")
	r.HandleFunc("/change", s.ChangeDeleteHandler).Methods("DELETE")
	s.http.Handler = r
}
