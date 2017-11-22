package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

type LoginResponse struct {
	Status string
}

func (incomingjson *User) Login(w http.ResponseWriter, r *http.Request) {

	log.Print("Inside Login Function")
	session, err := config.Instance(r)
	// log.Print(w.Header()["Set-Cookie"])
	// log.Print("After session assignment in login function")
	// log.Print("What is the value of session, err?")
	// log.Print("Session: ", session)
	// log.Print("Err: ", err)
	// log.Print("Set Value and Test")
	// session.Values["testing"] = "hellotheresailor"
	// log.Print("session.Values[testing]: ", session.Values["testing"])
	// log.Print("Test with previous values: ")
	// if session.Values["testingappend"] == nil {
	// 	log.Print("inside testingappen nil if statement")
	// 	session.Values["testingappend"] = "hohoho"
	// 	log.Print(session.Values["testingappend"])
	// } else {
	// 	log.Print("inside testingappend not nil else statement")
	// 	session.Values["testingappend"] = session.Values["testingappend"].(string) + session.Values["testing"].(string)
	// 	log.Print("Value of session.Values[testingappend]: ", session.Values["testingappend"])
	// }
	//
	// session.Save(r, w)
	// log.Print(w.Header()["Set-Cookie"])

	// log.Print("After assigning session in login")
	log.Print("loginjson: ", incomingjson)
	// log.Print("&loginjson.Username: ", &incomingjson.Username)
	// log.Print("loginjson.Username: ", incomingjson.Username)
	// marshalledusername, error := json.Marshal(loginjson.Username)
	// if error != nil {
	// 	panic(error)
	// }
	var tempvar string

	err = config.DB.QueryRow("SELECT username FROM users WHERE username=$1;", incomingjson.Username).Scan(&tempvar)
	log.Print("err: ", err)
	if err == nil {
		// 1 row
		log.Print("Found username")
		var passwordindatabase string
		config.DB.QueryRow("SELECT password FROM users WHERE username=$1;", &incomingjson.Username).Scan(&passwordindatabase)
		if passwordindatabase == incomingjson.Password {
			log.Print("username and password match!")
			session.Values["authenticated"] = true
			session.Values["username"] = incomingjson.Username
			// session.Values["Status"] = "Success, user logged in"
			session.Save(r, w)
			sessionusername := session.Values["username"]
			log.Print("Value of session.Values[authenticated], session.values[username]: ",
				session.Values["authenticated"], " ", sessionusername)
			response := LoginResponse{Status: "Success, user logged in"}
			json.NewEncoder(w).Encode(response)
		} else {
			log.Print("username and password don't match!")
			session.Values["authenticated"] = false
			session.Values["username"] = ""
			session.Save(r, w)
			response := LoginResponse{Status: "Failure, username and password don't match"}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		//empty result or error
		log.Print("Username not found or there was an error: ", err)
		response := LoginResponse{Status: "User not found!"}
		json.NewEncoder(w).Encode(response)
	}
}
