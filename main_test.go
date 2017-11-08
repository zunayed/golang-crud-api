package main

import (
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.InitializeDb("postgres", "", "crud_api", 5432)
	a.InitializeRouter()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")

}
