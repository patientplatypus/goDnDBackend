package users

import (
	"log"

	"github.com/patientplatypus/gorest/config"
)

func DeleteAll() {
	log.Print("inside DeleteAll() in users")
	config.DB.Query("TRUNCATE users")
}
