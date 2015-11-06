package main

import (
	jsn "github.com/elgs/gosqljson"
	"github.com/labstack/echo"

	"net/http"
)

func getStats(c *echo.Context) error {
	return c.JSON(http.StatusOK, server_stats.Data())
}

func getTestData(c *echo.Context) error {
	res, _ := jsn.QueryDbToMap(db, "camel", "select * from test1")
	return c.JSON(http.StatusOK, res)
}

func getEquipment(c *echo.Context) error {
	res, _ := jsn.QueryDbToMap(db, "camel", "select * from fm_equipment order by name")
	return c.JSON(http.StatusOK, res)
}

func getPartsList(c *echo.Context) error {
	res, _ := jsn.QueryDbToMap(db, "canel", "select * from fm_part")
	return c.JSON(http.StatusOK, res)
}
