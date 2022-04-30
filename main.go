package main

import (
	"WebDevelopment/views"
	"github.com/gorilla/mux"
	"net/http"
)

//====================================================================
//1. Create the global Template variable
//2. Parse our template file at and assign it to the variable
//3. Update our Template function to use the template variable
//====================================================================

var HomeView *views.View
var ContactView *views.View
var SignView *views.View

func main() {
	HomeView = views.NewView("bootstrap", "views/home.gohtml")
	ContactView = views.NewView("bootstrap", "views/contact.gohtml")
	SignView = views.NewView("bootstrap", "views/signup.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}

// Home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(HomeView.Render(w, nil))
}

// Contact handler
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(ContactView.Render(w, nil))
}

// Sign handler
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(SignView.Render(w, nil))
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}
