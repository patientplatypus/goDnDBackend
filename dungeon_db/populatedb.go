package dungeon

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func PopulateDB(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Instance(r)
	if session.Values["authenticated"] == false {
		if session.Values["username"] != "theGreateDM" {
			log.Print("Verboten!")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	}
	RaceType()
	ClassType()
	BackgroundType()
}
