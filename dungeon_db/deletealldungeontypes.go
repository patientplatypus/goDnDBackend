package dungeon

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func DeleteAllDungeonTypes(w http.ResponseWriter, r *http.Request) {
	log.Print("inside DeleteAllDungeonClasses() in dungeon")
	config.DB.Query("DROP TABLE dungeon_classes")
	config.DB.Query("DROP TABLE dungeon_races")
	config.DB.Query("DROP TABLE dungeon_backgrounds")
}
