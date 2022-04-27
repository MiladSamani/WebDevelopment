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

func main() {
	HomeView = views.NewView("bootstrap", "views/home.gohtml")
	ContactView = views.NewView("bootstrap", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := HomeView.Template.ExecuteTemplate(w, HomeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := ContactView.Template.ExecuteTemplate(w, ContactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
