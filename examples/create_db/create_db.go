package main

import (
	"fmt"
	"log"

	"github.com/safikka/libdb-go/db" // Sesuaikan dengan path libdb-go Anda
)

func main() {
	adapter, err := db.NewDatabaseAdapter("mysql", "username:password@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalf("Error initializing adapter: %v", err)
	}

	_, err = adapter.CreateDatabase("example_db")
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	fmt.Println("Database 'example_db' created successfully")
}
