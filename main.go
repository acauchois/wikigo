package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Contributor struct {
	ID   int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("Connecting to database...")

	db, err := sql.Open("mysql", "wikigo:wikigopwd@tcp(db:3306)/image_wikigo")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT id, firstname, lastname, email FROM contributor")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var contributor Contributor
		err = results.Scan(&contributor.ID, &contributor.Firstname, &contributor.Lastname, &contributor.Email)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(contributor.Firstname)
	}
}