package main

import (
	"database/sql"
	"log"
	//"net/http"
	"fmt"
)
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "1234"
	dbname = "onedb"
)
var db *sql.DB
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)

func initDB() error {
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error in connection to pg server: %s\n", err)
		return err
	}

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Printf("Error in connection to DB: %s\n", err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error in starting transaction: %s\n", err)
		return err
	}
	err = tx.Commit()
	if err != nil{
		tx.Rollback()
		log.Printf("Error in commiting transaction: %s\n", err)
		return err
	}
	return nil
}