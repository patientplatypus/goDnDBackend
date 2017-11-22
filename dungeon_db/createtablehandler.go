package dungeon

import (
	"log"

	"github.com/patientplatypus/gorest/config"
)

func Create_Table(finished chan bool, name string) {
	log.Print("Create table started")
	namestring := name + "name"
	tablename := "dungeon_" + name
	dataname := "data" + name
	querystring := "CREATE TABLE IF NOT EXISTS " + tablename + "(" + namestring + " text not null, " + dataname + " json);"
	config.DB.Query(querystring)
	log.Print("Create table finished")
	finished <- true
}
