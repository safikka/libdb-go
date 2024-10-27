package main

import (
	"fmt"
	"log"

	"github.com/safikka/libdb-go/db"
)

func main() {
	adapter, err := db.NewDatabaseAdapter("mysql", "username:password@tcp(localhost:3306)/example_db")
	if err != nil {
		log.Fatalf("Error initializing adapter: %v", err)
	}

	_, err = adapter.CreateTable("users", "id INT PRIMARY KEY, name VARCHAR(100)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table 'users' created successfully")
}
