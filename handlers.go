package main

import (
	"github.com/labstack/echo"
	"github.com/thoas/stats"

	"log"
	"net/http"
)

var server_stats *stats.Stats

func loadHandlers(e *echo.Echo) {

	e.SetDebug(true)

	// Point to the client application generated inside the webapp dir
	e.Index("./webapp/build/index.html")
	e.ServeDir("/", "./webapp/build/")

	server_stats = stats.New()
	e.Use(server_stats.Handler)

	e.Get("/stats", getStats)
	e.Get("/test1", getTestData)
	e.Get("/equipment", getEquipment)
	e.Get("/part", getPartsList)
	e.Get("/task", getTaskList)
	e.Get("/jtask", getJTaskList)
	e.Post("/login", login)
	e.Delete("/login", logout)
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

type loginCreds struct {
	Username string
	Password string
}

type loginResponse struct {
	Username string
	Role     string
	Token    string
}

func login(c *echo.Context) error {
	l := new(loginCreds)
	err := c.Bind(&l)
	if err != nil {
		log.Println("Bind Error:", err.Error())
	}
	log.Println("Login Credentials", l)

	sqlResult, _ := SQLMap(db, "select username from users where username=$1 and passwd=$2",
		l.Username,
		l.Password)
	log.Println("SQLResult", sqlResult)

	if len(sqlResult) == 1 {
		r := new(loginResponse)
		r.Username = l.Username
		r.Role = "Worker"
		r.Token = "98023840238402840"
		return c.JSON(http.StatusOK, r)
	} else {
		return c.String(http.StatusNotFound, "invalid")
	}
}

func logout(c *echo.Context) error {
	log.Println("Logging out ...")
	return c.String(http.StatusOK, "bye")
}
