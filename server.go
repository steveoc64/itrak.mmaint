package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/thoas/stats"

	pgsql "github.com/steveoc64/itrak.mmaint/pgsql"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// Runtime variables, held in external file config.json
type itrakMMaintConfig struct {
	Debug     bool
	SQLServer string
	WebPort   int
}

var itrak itrakMMaintConfig

// Load the config.json file, and override with runtime flags
func loadConfig() {
	cf, err := os.Open("config.json")
	if err != nil {
		log.Println("Could not open config.json :", err.Error())
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&itrak); err != nil {
		log.Fatalln("Failed to load config.json :", err.Error())
	}

	flag.BoolVar(&itrak.Debug, "debug", itrak.Debug, "Enable Debugging")
	flag.StringVar(&itrak.SQLServer, "sqlserver", itrak.SQLServer, "Name of SQLServer")
	flag.IntVar(&itrak.WebPort, "webport", itrak.WebPort, "Port Number for Web Server")
	flag.Parse()

	log.Printf("Starting\n\tDebug: \t\t%t\n\tSQLServer: \t%s\n\tWeb Port: \t%d\n",
		itrak.Debug,
		itrak.SQLServer,
		itrak.WebPort)
}

// Run the MicroServer
func main() {

	loadConfig()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	s := stats.New()
	e.Use(s.Handler)

	// Expose some Routes for testing
	e.Index("public/index.html")
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	// Connect to the SQLServer
	pgsql.SetDebug(itrak.Debug)
	err := pgsql.Dial(itrak.SQLServer)
	defer pgsql.Close()
	if err != nil {
		log.Fatalln("Exiting ..")
	}

	if itrak.Debug {
		log.Printf("Starting Web Server of port %d ...", itrak.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", itrak.WebPort))
}
