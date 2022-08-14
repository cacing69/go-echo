package api

import (
	"net/http"
  "github.com/gorilla/mux"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	// log.Println(*r)
	w.Write([]byte("lorem ipsum"))
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api", HandlerIndex).Methods(http.MethodGet)
	return r
}
