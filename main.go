package main

import (
	"WebDevelopment/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

//====================================================================
//1. Create the global Template variable
//2. Parse our template file at and assign it to the variable
//3. Update our Template function to use the template variable
//====================================================================

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
