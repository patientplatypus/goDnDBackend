package createcharacter

import (
	"log"

	"github.com/patientplatypus/gorest/config"
)

func Create_Table(finished chan bool) {
	log.Print("Create table started")
	querystring := "CREATE TABLE IF NOT EXISTS usercharacters(usernames text not null, characters json);"
	config.DB.Query(querystring)
	log.Print("Create table finished")
	finished <- true
}
