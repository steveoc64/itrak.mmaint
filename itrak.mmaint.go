package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowCredentials: true,
		Debug:            true,
	})
	e.Use(c.Handler)

	loadHandlers(e)

	// Start the web server
	if itrak.Debug {
		log.Printf("Starting Web Server of port %d ...", itrak.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", itrak.WebPort))
}
