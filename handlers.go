package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"

	"net/http"
)

var server_stats *stats.Stats

func loadHandlers(e *echo.Echo) {

	// Setup the web server

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(cors.Default().Handler)

	server_stats = stats.New()
	e.Use(server_stats.Handler)

	// Expose some Routes for testing
	e.Index("public/index.html")
	e.Get("/stats", getStats)
	e.Get("/test1", getTestData)
	e.Get("/equipment", getEquipment)
	e.Get("/part", getPartsList)
	e.Get("/task", getTaskList)
	e.Get("/jtask", getJTaskList)

}

func getStats(c *echo.Context) error {
	return c.JSON(http.StatusOK, server_stats.Data())
}

func getTestData(c *echo.Context) error {
	res, _ := SQLMap(db, "select * from test1")
	return c.JSON(http.StatusOK, res)
}

func getEquipment(c *echo.Context) error {
	res, _ := SQLMap(db, "select * from fm_equipment order by name")
	return c.JSON(http.StatusOK, res)
}

func getPartsList(c *echo.Context) error {
	res, _ := SQLMap(db, "select * from fm_part")
	return c.JSON(http.StatusOK, res)
}

func getTaskList(c *echo.Context) error {
	res, _ := SQLMap(db, "select lineno,instructions from fm_task order by lineno limit 300")
	return c.JSON(http.StatusOK, res)
}

func getJTaskList(c *echo.Context) error {
	res, _ := SQLJMap(db, "select lineno,instructions from fm_task order by lineno limit 300")
	return c.JSON(http.StatusOK, res)
}
