package main

import (
	"github.com/labstack/echo"
	"github.com/thoas/stats"

	"log"
	"net/http"
)

var server_stats *stats.Stats

/////////////////////////////////////////////////////////////////////////////////////////////////
// Add Handlers to the web server

func loadHandlers(e *echo.Echo) {

	if itrak.Debug {
		e.SetDebug(true)
	}

	// Point to the client application generated inside the webapp dir
	e.Index("./webapp/build/index.html")
	e.ServeDir("/", "./webapp/build/")

	server_stats = stats.New()
	e.Use(server_stats.Handler)

	e.Get("/stats", getStats)
	e.Get("/test1", getTestData)
	e.Get("/part", getPartsList)
	e.Get("/task", getTaskList)
	e.Get("/jtask", getJTaskList)

	e.Post("/login", login)
	e.Delete("/login", logout)

	e.Get("/people", getPeople)
	e.Post("/people", createPeople)
	e.Patch("/people/:id", updatePeople)
	e.Delete("/people/:id", deletePeople)

	e.Get("/site", getSite)
	e.Post("/site", createSite)
	e.Patch("/site/:id", updateSite)
	e.Delete("/site/:id", deleteSite)

	e.Get("/roles", getRoles)
	e.Get("/equipment", getEquipment)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Some simple test case handlers

func getStats(c *echo.Context) error {
	return c.JSON(http.StatusOK, server_stats.Data())
}

func getTestData(c *echo.Context) error {
	res, _ := SQLMap(db, "select * from test1")
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

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling logins

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

	sqlResult, _ := SQLMap(db, "select username,role from users where username=$1 and passwd=$2",
		l.Username,
		l.Password)
	log.Println("SQLResult", sqlResult)

	if len(sqlResult) == 1 {
		r := new(loginResponse)
		r.Username = l.Username
		r.Role = sqlResult[0]["role"]
		r.Token = "98023840238402840"
		return c.JSON(http.StatusOK, r)
	} else {
		return c.String(http.StatusUnauthorized, "invalid")
	}
}

func logout(c *echo.Context) error {
	log.Println("Logging out ...")
	return c.String(http.StatusOK, "bye")
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling people requests

func getPeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from person order by name")
	return c.JSON(http.StatusOK, sqlResult)
}
func createPeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}
func updatePeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}
func deletePeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling Site requests

func getSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from site order by name")
	return c.JSON(http.StatusOK, sqlResult)
}
func createSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from site")
	return c.JSON(http.StatusOK, sqlResult)
}
func updateSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from site")
	return c.JSON(http.StatusOK, sqlResult)
}
func deleteSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from site order by name")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling the Roles table

func getRoles(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from roles order by id")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling the Equipment table

func getEquipment(c *echo.Context) error {
	sqlResult, _ := SQLMap(db, "select * from equipment order by name")
	return c.JSON(http.StatusOK, sqlResult)
}
