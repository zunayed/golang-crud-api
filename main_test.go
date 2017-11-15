package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	dbtype := os.Getenv("DBTYPE")

	if dbtype == "postgres" {
		log.Printf("Using real postgres backend")
		a.InitializeDb("postgres", "", "postgres", 5432, "")
	} else {
		log.Printf("Using SQL lite backend")
		a.InitializeDb("", "", "", 5432, "test")
	}
	a.InitializeRouter()
	code := m.Run()
	a.DB.DropTable(&Product{})
	os.Exit(code)
}
