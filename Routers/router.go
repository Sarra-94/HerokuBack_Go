package router

import (
	controllers "template/Controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/authentification", controllers.Isauthenticated).Methods("POST")
	return r
}
