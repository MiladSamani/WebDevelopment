package controllers

import (
	"WebDevelopment/models"
	"WebDevelopment/views"
	"fmt"
	"net/http"
)

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//NewUsers function that will be used to construct and return a Users object.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}

}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user tries to create a new user account. POST /signup
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
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err := fmt.Fprintln(w, "User is", user)
	if err != nil {
		return
	}
}

// Login is used to process the login form when a user tries to log in as an existing user (via email & pw). POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		_, err2 := fmt.Fprintln(w, "Invalid email address.")
		if err2 != nil {
			return
		}
	case models.ErrInvalidPassword:
		_, err2 := fmt.Fprintln(w, "Invalid password provided.")
		if err2 != nil {
			return
		}
	case nil:
		_, err2 := fmt.Fprintln(w, user)
		if err2 != nil {
			return
		}
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
