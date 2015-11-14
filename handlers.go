package main

import (
	"database/sql"
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
	e.Get("/people/:id", getPerson)
	e.Post("/people/:id", savePerson)

	e.Get("/site", getSites)
	e.Get("/site/:id", getSite)
	e.Post("/site/:id", saveSite)

	e.Get("/roles", getRoles)
	e.Get("/vendors", getVendors)

	// Equipment Related functions
	e.Get("/equipment", getAllEquipment)
	e.Get("/equipment/:id", getEquipment)
	e.Post("/equipment/:id", saveEquipment)
	e.Get("/subparts/:id", subParts)
	e.Get("/spares", getAllSpares)
	e.Get("/spares/:id", getEquipment)
	e.Post("/spares/:id", saveEquipment)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Some type conversion helper functions

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
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

type PeopleType struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	HourlyRate string `json:"hourlyrate"`
	Comments   string `json:"comments"`
	Alternate  string `json:"alternate"`
	Role       string `json:"role"`
	Location   string `json:"location"`
	Calendar   string `json:"calendar"`
}

func getPeople(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from person order by name")
	return c.JSON(http.StatusOK, sqlResult)
}

func getPerson(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	sqlResult, err := SQLMapOne(db, `select * from person where id=$1`, id)
	if err != nil {
		log.Println(err.Error())
	}
	return c.JSON(http.StatusOK, sqlResult)
}

func savePerson(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	person := new(PeopleType)
	if binderr := c.Bind(person); binderr != nil {
		log.Println(binderr.Error())
		return binderr
	}
	log.Println(person)

	_, err := ExecDb(db,
		`update person
			set user_id=$2,
			    name=$3,
			    email=$4,
			    phone=$5,
			    hourlyrate=$6,
			    comments=$7,
			    alternate=$8,
			    role=$9,
			    location=$10
			where id=$1`,
		id,
		ToNullInt64(person.UserID),
		person.Name,
		person.Email,
		person.Phone,
		person.HourlyRate,
		person.Comments,
		person.Alternate,
		ToNullInt64(person.Role),
		ToNullInt64(person.Location))

	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

/////////////////////////////////////////////////////////////////////////////////////////////////
// Logic for handling Site requests

type SiteType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	ContactName string `json:"contactname"`
}

func getSites(c *echo.Context) error {
	sqlResult, _ := SQLMap(db,
		"select * from site order by name")
	return c.JSON(http.StatusOK, sqlResult)
}

func getSite(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	sqlResult, err := SQLMapOne(db, `select * from site where id=$1`, id)
	if err != nil {
		log.Println(err.Error())
	}
	return c.JSON(http.StatusOK, sqlResult)
}

func saveSite(c *echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.String(http.StatusNotAcceptable, "Invalid ID")
	}

	site := new(SiteType)
	if binderr := c.Bind(site); binderr != nil {
		log.Println(binderr.Error())
		return binderr
	}
	log.Println(site)

	_, err := ExecDb(db,
		`update site
			set name=$2,
			    address=$3,
			    phone=$4,
			    contactname=$5
			where id=$1`,
		id,
		site.Name,
		site.Address,
		site.Phone,
		site.ContactName)

	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(http.StatusOK, site)
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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Descr       string `json:"descr"`
	Comments    string `json:"comments"`
	Modelno     string `json:"modelno"`
	Serialno    string `json:"serialno"`
	Location    string `json:"location"`
	Parent_id   string `json:"parent_id"`
	Category    string `json:"category"`
	Vendor      string `json:"vendor"`
	Parent_name string `json:"parent_name"`
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

func getAllSpares(c *echo.Context) error {
	sqlResult, err := SQLMap(db,
		`select e.*,
			p.name as parent_name,
			l.name as location_name
		from equipment e 
			left outer join equipment p on (p.id=e.parent_id)
			left outer join site l on (l.id=e.location)
		where e.category=3
		order by location_name,e.name`)
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

	_, err := ExecDb(db,
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

	return c.JSON(http.StatusOK, eq)
}
