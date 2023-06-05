package apiRoute

import (
	"github.com/gorilla/mux"
	controller "github.com/shwetank0714/jwtmodfile/controllers"
)


func UserRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/users", controller.GetAllUsersList).Methods("GET")
	router.HandleFunc("/api/users/create", controller.CreateUser).Methods("POST")

	return router
}