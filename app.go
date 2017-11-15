package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) InitializeDb(user, password, dbname string, dbport int, dbtype string) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable port=%v",
		user,
		password,
		dbname,
		dbport,
	)

	var err error
	if dbtype == "test" {
		a.DB, err = gorm.Open("sqlite3", ":memory:")
	} else {
		a.DB, err = gorm.Open("postgres", connectionString)
		if err != nil {
			panic("failed to connect database")
		}
	}

	// Migrate the schema
	a.DB.AutoMigrate(&Product{})
}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Printf("Running server on %v", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
