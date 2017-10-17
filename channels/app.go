package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App - contains application core components, such as connection to database
// and router
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize - creates connect to a database and prepares all required
// routes for the application
func (a *App) Initialize(user, password, dbname string) {}

// Run - starts an application
func (a *App) Run(addr string) {}
