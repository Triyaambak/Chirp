package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func NewDB(path string) (*DB, error) {
	err := os.WriteFile(path, []byte{}, 0666)
	if err != nil {
		return nil, err
	}
	newDB := &DB{path: path}
	return newDB, nil
}

func (db *DB) ensureDB() error {
	_, err := os.Stat(db.path)

	if os.IsNotExist(err) {
		log.Println("Database.json does not exist, creating a new file")

		newDB, err := NewDB(db.path)
		if err != nil {
			return err
		}

		*db = *newDB
	}

	return nil
}

func (db *DB) CreateChirps(w http.ResponseWriter, r *http.Request) {
	err := db.ensureDB()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errResp, _ := json.Marshal(struct{ Error string }{Error: "Something went wrong while ensuring / creating db"})
		w.Write(errResp)
	}
	data, err := ValidateJSON(w, r)
	if err != nil {
		return
	}
	log.Print(string(data))
}
