package config

import (
	"database/sql"
)

// var (
// 	DB *sql.DB
// )

var DB *sql.DB

// var DB = Sql_config()

// func Sql_connect() (db *sql.DB) {
//
// 	sqlConfig := reflect.ValueOf(Sql_config())
// 	host := reflect.Indirect(sqlConfig).FieldByName("host")
// 	port := reflect.Indirect(sqlConfig).FieldByName("port")
// 	user := reflect.Indirect(sqlConfig).FieldByName("user")
// 	password := reflect.Indirect(sqlConfig).FieldByName("password")
// 	dbname := reflect.Indirect(sqlConfig).FieldByName("dbname")
//
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer db.Close()
//
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println("Successfully connected~!")
// 	return db
// }
