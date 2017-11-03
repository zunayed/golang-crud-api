package main

import (
	//"database/sql"
	//"errors"
	"log"
	"os"
	"testing"

	//_ "github.com/mattn/go-sqlite3"
)

var a App

func TestMain(m *testing.M) {
	// err := errors.New("")

	// a.DB, err = sql.Open(
	// 	"sqlite3", ":memory:",
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	a.InitializeDb("postgres", "", "crud_api")
	a.InitializeRouter()
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")

}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
	id SERIAL,
	name VARCHAR(50) NOT NULL,
	age INT NOT NULL
)`
