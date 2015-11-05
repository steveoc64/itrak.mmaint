package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/thoas/stats"

	// Utility routines for talking to SQLServer
	"github.com/steveoc64/itrak/mssql"

	//	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	//	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
	//"time"
)

// Runtime variables, held in external file config.json
type itrakConfig struct {
	Debug     bool
	SQLServer string
	SQLPort   int
	WebPort   int
	User      string
	Password  string
}

var itrak itrakConfig

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
	flag.IntVar(&itrak.SQLPort, "sqlport", itrak.SQLPort, "Port Number for SQLServer")
	flag.StringVar(&itrak.User, "user", itrak.User, "Username for SQLServer")
	flag.StringVar(&itrak.Password, "password", itrak.Password, "Password for SQLServer")
	flag.IntVar(&itrak.WebPort, "webport", itrak.WebPort, "Port Number for Web Server")
	flag.Parse()

	log.Printf("Starting\n\tDebug: \t\t%t\n\tSQLServer: \t%s:%d\n\tSQLUser: \t%s [%s]\n\tWeb Port: \t%d\n",
		itrak.Debug,
		itrak.SQLServer, itrak.SQLPort,
		itrak.User, itrak.Password,
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
	e.Get("/skills", func(c *echo.Context) error {
		res, _ := mssql.Query("select * from SBSiTrak.dbo.DR_SkillLevels")
		return c.JSON(http.StatusOK, res)
	})
	e.Get("/countries", func(c *echo.Context) error {
		res, _ := mssql.Query("exec SBSiTrak.dbo.CTY_GetCountries")
		return c.JSON(http.StatusOK, res)

	})

	// Connect to the SQLServer
	mssql.SetDebug(itrak.Debug)
	err := mssql.Dial(itrak.SQLServer, itrak.SQLPort, itrak.User, itrak.Password)
	defer mssql.Close()
	if err != nil {
		log.Fatalln("Exiting ..")
	}

	if itrak.Debug {
		log.Println("Starting Web Server ...")
	}
	e.Run(fmt.Sprintf(":%d", itrak.WebPort))
}
