package createcharacter

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func DeleteAllUsersCharacters(w http.ResponseWriter, r *http.Request) {
	log.Print("inside Deleteallusercharacters() in dungeon")
	config.DB.Query("DROP TABLE userscharacters")
}
