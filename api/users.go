package api

import (
	"net/http"
//   "github.com/gorilla/mux"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	// log.Println(*r)
	w.Write([]byte("api/user"))
}

// func GetRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api", Handler).Methods(http.MethodGet)
// 	return r
// }
