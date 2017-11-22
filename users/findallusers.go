package users

import (
	"fmt"
	"log"

	"github.com/patientplatypus/gorest/config"
)

func QueryAll() {
	rows, err := config.DB.Query("SELECT username, id FROM users")
	defer rows.Close()
	log.Print("Here are the usernames in the database:")
	for rows.Next() {
		var username string
		var id int
		err = rows.Scan(&username, &id)
		// err = rows.Scan(&username)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(username, id)
		// log.Print(username)
	}
}
