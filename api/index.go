package handler

import (
	"net/http"
  "github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)
	w.Write("lorem")
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods(http.MethodGet)
	return r
}
