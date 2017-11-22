package createcharacter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

var Username string
var Checkok bool

type Character struct {
	Charactername string
	Id            int
}

type ErrorReturn struct {
	Status string `json:"status"`
}

// type Charactername string

func SessionsCheck(w http.ResponseWriter, r *http.Request) (username string, checkok bool) {
	session, _ := config.Instance(r)
	// log.Print("inside sessionscheck...what is the value of stuff....")
	// log.Print("value of err: ", err)
	// log.Print("session: ", session)
	// log.Print("session.Values: ", session.Values)
	// log.Print("username: ", session.Values["username"])
	// log.Print("username string: ", session.Values["username"].(string))
	// log.Print("authenticated: ", session.Values["authenticated"])

	if session.Values["username"] == nil {
		log.Print("inside if statement 1)")
		fmt.Printf("User not logged in!")
		var errorreturn ErrorReturn
		errorreturn.Status = "user not logged in!"
		json.NewEncoder(w).Encode(errorreturn)
		return "nil", false
	} else {
		log.Print("inside if statement 2)")
		if session.Values["username"] == nil {
			log.Print("inside if statement 3)")
			if session.Values["authenticated"] == false {
				log.Print("inside if statement 4)")
				log.Print("Verboten!")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return "nil", false
			}
		}
		log.Print("inside if statement 4)")
		return session.Values["username"].(string), true
	}
	log.Print("inside if statement 5)")
	return "nil", false
}

func NewCharacter(w http.ResponseWriter, r *http.Request) {
	Username, Checkok = SessionsCheck(w, r)
	// charactername := r.FormValue("charactername")
	if Checkok == false {
		return
	}
	log.Print("Inside NewCharacter")
	var incomingjson Character
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	log.Print("value of r.Body: ", r.Body)
	decoder := json.NewDecoder(r.Body)
	log.Print("after json decoder")
	err := decoder.Decode(&incomingjson)
	log.Print("after json decoder2")
	if err != nil {
		panic(err)
	}
	log.Print("after panic err")
	log.Print("value of incomingjson: ", incomingjson)

	charactername := incomingjson.Charactername
	log.Print("charactername: ", charactername)

	// log.Print("username: ", Username)

	RegisterCharacter(charactername, w, r)
}

func SavedCharacters(w http.ResponseWriter, r *http.Request) {
	log.Print("Inside SavedCharacters")
	Username, Checkok = SessionsCheck(w, r)
	if Checkok != false {
		RetrieveSavedCharacters(Username, w, r)
	}
}
