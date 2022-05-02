package controllers

import (
	"WebDevelopment/views"
	"fmt"
	"net/http"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//NewUsers function that will be used to construct and return a Users object.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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
	_, err := fmt.Fprintln(w, "Name is", form.Name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(w, "Email is", form.Email)
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(w, "Password is", form.Password)
	if err != nil {
		return
	}
}
