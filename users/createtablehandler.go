package users

import (
	"log"

	"github.com/patientplatypus/gorest/config"
)

func Create_Table(finished chan bool, name string) {
	log.Print("Create table started")
	querystring := "CREATE TABLE IF NOT EXISTS users();"
	config.DB.Query(querystring)
	log.Print("Create table finished")
	finished <- true
}
