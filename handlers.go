package main

import (
	"github.com/labstack/echo"
	"github.com/thoas/stats"
	"log"
	"net/http"
	"strconv"
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
	e.Get("/vendors", getVendors)

	e.Get("/equipment", getAllEquipment)
	e.Get("/equipment/:id", getEquipment)
	e.Post("/equipment/:id", saveEquipment)
	e.Get("/subparts/:id", subParts)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Some simple test case handlers

func getStats(c *echo.Context) error {
	return c.JSON(http.StatusOK, server_stats.Data())
}

func getTestData(c *echo.Context) error {
	res, _ := SQLMap(db,
		"select * from test1")
	return c.JSON(http.StatusOK, res)
}

func getPartsList(c *echo.Context) error {
	res, _ := SQLMap(db,
		"select * from fm_part")
	return c.JSON(http.StatusOK, res)
}

func getTaskList(c *echo.Context) error {
	res, _ := SQLMap(db,
		"select lineno,instructions from fm_task order by lineno limit 300")
	return c.JSON(http.StatusOK, res)
}

func getJTaskList(c *echo.Context) error {
	res, _ := SQLJMap(db,
		"select lineno,instructions from fm_task order by lineno limit 300")
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
	//log.Println("Login Credentials", l)

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
	sqlResult, _ := SQLMap(db,
		"select * from person order by name")
	return c.JSON(http.StatusOK, sqlResult)
}
func createPeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}
func updatePeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}
func deletePeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from person")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling Site requests

func getSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from site order by name")
	return c.JSON(http.StatusOK, sqlResult)
}
func createSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from site")
	return c.JSON(http.StatusOK, sqlResult)
}
func updateSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from site")
	return c.JSON(http.StatusOK, sqlResult)
}
func deleteSite(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from site order by name")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling the Roles table

func getRoles(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from roles order by id")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling the Vendors table

func getVendors(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from vendor order by id")
	return c.JSON(http.StatusOK, sqlResult)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling the Equipment table

type equipmentType struct {
	ID        string
	Name      string
	Descr     string
	Comments  string
	Modelno   string
	Serialno  string
	Location  string
	Parent_id string
	Category  string
	Vendor    string
}

func getAllEquipment(c *echo.Context) error {
	sqlResult, err := SQLMap(db,
		`select e.*,
			p.name as parent_name,
			l.name as location_name
		from equipment e 
			left outer join equipment p on (p.id=e.parent_id)
			left outer join site l on (l.id=e.location)
		order by location_name,e.name`)
	if err != nil {
		log.Println(err.Error())
	}
	return c.JSON(http.StatusOK, sqlResult)
}

func subParts(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}
	sqlResult, err := SQLMap(db,
		`select e.*,
			p.name as parent_name,
			l.name as location_name
		from equipment e 
			left outer join equipment p on (p.id=e.parent_id)
			left outer join site l on (l.id=e.location)
		where e.parent_id=$1
		order by location_name,e.name`, id)
	if err != nil {
		log.Println(err.Error())
	}
	return c.JSON(http.StatusOK, sqlResult)
}

func getEquipment(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	sqlResult, err := SQLMapOne(db,
		`select e.*,
			p.name as parent_name,
			l.name as location_name
		from equipment e 
			left outer join equipment p on (p.id=e.parent_id)
			left outer join site l on (l.id=e.location)
		where e.id=$1
		order by location_name,e.name`, id)
	if err != nil {
		log.Println(err.Error())
	}
	return c.JSON(http.StatusOK, sqlResult)
}

func saveEquipment(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	eq := new(equipmentType)
	if binderr := c.Bind(eq); binderr != nil {
		log.Println(binderr.Error())
		return binderr
	}
	log.Println(eq)

	sqlResult, err := ExecDb(db,
		`update equipment 
			set name=$2,
			    descr=$3,
			    comments=$4,
			    modelno=$5,
			    serialno=$6,
			    location=$7,
			    vendor=$8,
			    category=$9
			where id=$1`,
		id,
		eq.Name,
		eq.Descr,
		eq.Comments,
		eq.Modelno,
		eq.Serialno,
		eq.Location,
		eq.Vendor,
		eq.Category)

	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(http.StatusOK, sqlResult)
}
