package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/thoas/stats"
	"log"
)

var db *sql.DB
var server_stats *stats.Stats

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
	server_stats = stats.New()
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(server_stats.Handler)

	// Expose some Routes for testing
	e.Index("public/index.html")
	e.Get("/stats", getStats)
	e.Get("/test1", getTestData)
	e.Get("/equipment", getEquipment)
	e.Get("/part", getPartsList)

	// Start the web server
	if itrak.Debug {
		log.Printf("Starting Web Server of port %d ...", itrak.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", itrak.WebPort))
}
