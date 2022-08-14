package api

import (
	"net/http"
	"net/http"
	"github.com/cacing69/go-echo/api/const"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// log.Println(*r)
	w.Write([]byte(const.Test))
}

// func GetRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api", Handler).Methods(http.MethodGet)
// 	return r
// }
