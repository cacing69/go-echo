package api

import (
	"net/http"
  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// log.Println(*r)
	w.Write([]byte("lorem ipsum"))
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api", Index).Methods(http.MethodGet)
	return r
}
