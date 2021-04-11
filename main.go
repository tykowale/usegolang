package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/controllers"
	"lenslocked.com/models"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "dev"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	userService, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer userService.Close()
	userService.AutoMigrate()

	usersController := controllers.NewUsers(userService)
	staticController := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.Handle("/login", usersController.LoginView).Methods("GET")
	r.HandleFunc("/login", usersController.Login).Methods("POST")
	r.NotFoundHandler = http.Handler(staticController.NotFoundView)
	http.ListenAndServe(":3000", r)
}
