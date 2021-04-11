package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

type Users struct {
	NewView     *views.View
	LoginView   *views.View
	userService *models.UserService
}

func NewUsers(userService *models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "users/new"),
		LoginView:   views.NewView("bootstrap", "users/login"),
		userService: userService,
	}
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)

	if err != nil {
		panic(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.userService.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid credentials")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid credentials")
	case nil:
		fmt.Fprintln(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
