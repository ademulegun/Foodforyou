package Persistence

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetUpDatabase(){
	log.Println("Connecting to the database...")
	var err error
	DbConn, err = sql.Open("mysql", "sa:P@ssw0rd$123@tcp(127.0.0.1:3306)/foodforyou")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the data successfully")
}