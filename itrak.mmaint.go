package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

// Run the MicroServer
func main() {

	LoadConfig()

	// Connect to the SQLServer
	var err error

	db, err = sql.Open("postgres", itrak.DataSourceName)
	defer db.Close()
	if err != nil {
		log.Fatalln("Exiting ..")
	}

	// Setup the web server
	e := echo.New()
	loadHandlers(e)

	// Start the web server
	if itrak.Debug {
		log.Printf("Starting Web Server of port %d ...", itrak.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", itrak.WebPort))
}
