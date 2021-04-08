package main

import (
	"log"
	"os"

	"github.com/rav1L/book-machine/modules/db"
)

const (
	dbDriver = "postgres"
)

var (
	myDB   db.ISQL
	dbPath = os.Getenv("DATABASE_URL")
)

func main() {
	myDB = &db.Handler{}
	err := myDB.Init(dbDriver, dbPath)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer myDB.Disconnect()
	falseID := "dee44255-bf77-4e3c-8a2e-ca35ae05f861"
	trueID := "dee44255-bf77-4e3c-8a2e-ca35ae05f860"
	falseBook, err := myDB.Books().Get(falseID)
	if err != nil {
		log.Printf("%+v", err)
	}
	trueBook, err := myDB.Books().Get(trueID)
	if err != nil {
		log.Printf("%+v", err)
	}
	log.Println(falseBook)
	log.Println(trueBook)
}
