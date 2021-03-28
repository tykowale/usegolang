package main

import (
	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"net/http"
)

func main() {
	usersController := controllers.NewUsers()
	staticController := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.NotFoundHandler = http.Handler(staticController.NotFoundView)
	http.ListenAndServe(":3000", r)
}
