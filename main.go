package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/patientplatypus/gorest/config"

	"github.com/gorilla/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "patientplatypus"
	password = "Fvnjty0b"
	dbname   = "dungeon_world"
)

// var DB = config.Sql_connect()

// var DB config.Sqlconfig

// var DB *sql.DB

func main() {
	// DB = config.Sql_config()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "patientplatypus", "Fvnjty0b", "dungeon_world")
	var err error
	config.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = config.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected~!")

	log.Print("Instatiate Cookiestore")
	config.Configure_Sessions()
	log.Print("Instantiated")

	router := NewRouter()
	// os.Setenv("ORIGIN_ALLOWED", "*")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	credentialsOK := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk, credentialsOK)(router)))

}
