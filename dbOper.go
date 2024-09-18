package main

import (
	"log"
	"net/http"
	"os"
)

// func NewDB (path string )(*DB , error){

// }

func (db *DB) CreateChirps(w http.ResponseWriter, r *http.Request) {
	data, err := ValidateJSON(w, r)
	if err != nil {
		return
	}
	log.Print(string(data))
}

func (db *DB) ensureDB() error {
	_, err := os.Stat("database.json")

	if os.IsNotExist(err) {
		log.Println("Database.json does not exist , creating a new file")
		os.WriteFile("database.json", []byte{}, 0666)
	}

	return nil
}
