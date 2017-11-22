package createcharacter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func RetrieveSavedCharacters(username string, w http.ResponseWriter, r *http.Request) {
	log.Print("Inside RetrieveSavedCharacters function")
	var databasecharacters UserCharacters
	var datarows []byte
	config.DB.QueryRow("SELECT characters FROM usercharacters where usernames = $1", Username).Scan(&datarows)
	json.Unmarshal(datarows, &databasecharacters)
	json.NewEncoder(w).Encode(databasecharacters)
}
