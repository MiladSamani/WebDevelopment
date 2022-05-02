package main

import (
	"WebDevelopment/controllers"
	"WebDevelopment/models"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Milad1377007"
	dbname   = "Milad"
)

func main() {
	// Create a DB connection string and then use it to create our model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer func(us *models.UserService) {
		err := us.Close()
		if err != nil {
			panic(err)
		}
	}(us)
	err = us.AutoMigrate()
	if err != nil {
		return
	}
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
