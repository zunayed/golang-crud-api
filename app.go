package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) InitializeDb(user, password, dbname string, dbport int) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable port=%v",
		user,
		password,
		dbname,
		dbport,
	)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
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
