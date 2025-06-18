package main

import (
	"log"
	"database/sql"
)


func main()  {
	var err error
	db, err := sql.Open("sqlite", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	initDB()

}
