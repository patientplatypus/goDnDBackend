package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	//super duper dangerous - Axe in production!
	// log.Print("Waiting for create table")
	// finished := make(chan bool)
	// go Create_Table(finished, "classes")
	// <-finished
	// log.Print("Continuing with the rest of the program")

	DeleteAll()
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var incomingjson User
	err := decoder.Decode(&incomingjson)

	if err != nil {
		panic(err)
	}

	username := incomingjson.Username
	password := incomingjson.Password

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished, "classes")
	<-finished
	log.Print("Continuing with the rest of the program")

	if username != "" && password != "" {
		incomingjson.Register(w, r)
	} else {
		fmt.Fprintln(w, "error username or password is blank!")
	}
}

type User struct {
	Username string
	Password string
	Id       int
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var incomingjson User
	err := decoder.Decode(&incomingjson)

	if err != nil {
		panic(err)
	}

	username := incomingjson.Username
	password := incomingjson.Password

	log.Print("username: ", username)
	log.Print("password: ", password)
	if username != "" && password != "" {
		incomingjson.Login(w, r)
	} else {
		fmt.Fprintln(w, "error username or password is blank!")
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	Logout(w, r)
}

func QueryAllUsers(w http.ResponseWriter, r *http.Request) {
	QueryAll()
}
