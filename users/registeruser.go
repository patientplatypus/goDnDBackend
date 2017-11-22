package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

type RegisterResponse struct {
	Status string
}

func (incomingjson *User) Register(w http.ResponseWriter, r *http.Request) {
	if err := config.DB.QueryRow("SELECT username FROM users WHERE username=$1;", &incomingjson.Username).Scan(&incomingjson.Username); err == nil {
		// 1 row
		fmt.Println("there was 1 or more rows returned and err is: ", err)
		fmt.Println("You cannot add a user with that name as a user with that name already exists! oh no!")
		response := LoginResponse{Status: "Failure, other user with that name already exists!"}
		json.NewEncoder(w).Encode(response)
		// QueryAll()
	} else if err == sql.ErrNoRows {
		// empty result
		fmt.Println("no rows from sql and err is: ", err)
		config.DB.QueryRow("insert into users (username, password, id) values ($1, $2, $3);", incomingjson.Username, incomingjson.Password, rand.Intn(999999))
		fmt.Println("User inserted and rows is equal to: ", err)
		incomingjson.Login(w, r)
	} else {
		// error
		response := LoginResponse{Status: "Failure, Some other error!"}
		json.NewEncoder(w).Encode(response)
		fmt.Println("Something has gone wrong. err is: ", err)
	}
}
