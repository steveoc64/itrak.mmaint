package main

import (
	"github.com/labstack/echo"

	"net/http"
)

func loadHandlers(e *echo.Echo) {
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
