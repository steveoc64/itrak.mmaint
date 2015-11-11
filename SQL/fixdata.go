// Quick and Dirty program to create all the relations in the (new) database

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

// Runtime variables, held in external file config.json
type iTrakMMaintConfig struct {
	Debug          bool
	DataSourceName string
	WebPort        int
}

var itrak iTrakMMaintConfig

// Load the config.json file, and override with runtime flags
func LoadConfig() {
	cf, err := os.Open("../config.json")
	if err != nil {
		log.Println("Could not open ../config.json :", err.Error())
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&itrak); err != nil {
		log.Fatalln("Failed to load config.json :", err.Error())
	}

	log.Printf("Starting\n\tDebug: \t\t%t\n\tSQLServer: \t%s\n\tWeb Port: \t%d\n",
		itrak.Debug,
		itrak.DataSourceName,
		itrak.WebPort)
}

func SQLMap(db *sql.DB, sqlStatement string, sqlParams ...interface{}) ([]map[string]string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var results []map[string]string
	rows, err := db.Query(sqlStatement, sqlParams...)
	if err != nil {
		fmt.Println("Error executing: ", sqlStatement)
		return results, err
	}
	cols, _ := rows.Columns()
	colsLower := make([]string, len(cols))

	for i, v := range cols {
		colsLower[i] = strings.ToLower(v)
	}

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		result := make(map[string]string, len(cols))
		rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[cols[i]] = ""
			} else {
				result[cols[i]] = string(raw)
			}
		}
		results = append(results, result)
	}
	return results, nil
}

func ExecDb(db *sql.DB, sqlStatement string, sqlParams ...interface{}) (int64, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	result, err := db.Exec(sqlStatement, sqlParams...)
	if err != nil {
		fmt.Println("Error executing: ", sqlStatement)
		fmt.Println(err)
		return 0, err
	}
	return result.RowsAffected()
}

func main() {

	// Connect to the SQLServer
	var err error

	LoadConfig()

	db, err = sql.Open("postgres", itrak.DataSourceName)
	defer db.Close()
	if err != nil {
		log.Fatalln("Exiting ..", err.Error())
	}
	log.Println("DSN:", itrak.DataSourceName)

	///////////////////////////////////////////////////////////////////////////////////////
	// Person table
	//   - fix location
	sqlResult, _err := SQLMap(db, "select * from fm_person")
	if _err != nil {
		log.Println("Error:", _err.Error())
	}
	for _, row := range sqlResult {
		// Lookup the location id
		_res, _err := SQLMap(db, "select id from site where name=$1", row["location"])
		if _err != nil {
			log.Fatal(_err.Error())
		} else {
			siteId := _res[0]["id"]
			personName := row["personname"]
			updatedRows, _ := ExecDb(db, "update person set location=$2 where name=$1", personName, siteId)
			log.Println("Updated", updatedRows, "persons setting location =", siteId, "where name =", personName)

		}
	}
	ExecDb(db, "update person set location=null where location=0")
	ExecDb(db, "update person set role=null where role=0")
	ExecDb(db, "alter table person add foreign key(location) references site")
	ExecDb(db, "alter table person add foreign key(role) references roles")

	///////////////////////////////////////////////////////////////////////////////////////
	// Equipment table
	//		- fix location
	//		- fix the category
	// 		- fix the vendor
	//		- fix the parent ID
	sqlResult, _err = SQLMap(db, "select * from fm_equipment")
	if _err != nil {
		log.Println("Error:", _err.Error())
	}
	for _, row := range sqlResult {
		// Lookup the location id
		_res, _err := SQLMap(db, "select id from site where name=$1", row["location"])
		if _err != nil {
			log.Fatal(_err.Error())
		}
		var siteId int
		if len(_res) > 0 {
			siteId, _err = strconv.Atoi(_res[0]["id"])
		}

		// now get the category ID
		var catId int
		_res, _err = SQLMap(db, "select id from equip_type where name=$1", row["category"])
		if _err != nil {
			log.Fatal(_err.Error())
		}
		if len(_res) > 0 {
			catId, _err = strconv.Atoi(_res[0]["id"])
		}
		log.Println("Equip Category = ", catId)

		// get the parentId
		var parentId int
		if row["partof"] != "" {
			_res, _err = SQLMap(db, "select id from equipment where name=$1", row["partof"])
			if _err != nil {
				log.Fatal(err.Error())
			}
			if len(_res) > 0 {
				parentId, _err = strconv.Atoi(_res[0]["id"])
			}
		}

		// Now get the Vendor
		vendorId := 0
		if row["vendor"] != "" {
			_res, _err = SQLMap(db, "select id from vendor where name=$1", row["vendor"])
			if _err != nil {
				log.Fatal(err.Error())
			}
			if len(_res) > 0 {
				vendorId, _err = strconv.Atoi(_res[0]["id"])
			}
		}

		updatedRows, _eqErr := ExecDb(db, "update equipment set location=$2,category=$3,vendor=$4,parent_id=$5 where name=$1",
			row["name"],
			siteId,
			catId,
			vendorId,
			parentId)

		if _eqErr != nil {
			log.Fatal(_eqErr.Error())
		}

		log.Println("Updated", updatedRows, "equipments setting location =", siteId,
			", category =", catId,
			", vendor =", vendorId,
			", parent part =", parentId,
			"where name =", row["name"])
	}
	ExecDb(db, "update equipment set location=null where location=0")
	ExecDb(db, "update equipment set category=null where category=0")
	ExecDb(db, "update equipment set vendor=null where vendor=0")
	ExecDb(db, "update equipment set parent_id=null where parent_id=0")
	ExecDb(db, "alter table equipment add foreign key(location) references site")
	ExecDb(db, "alter table equipment add foreign key(category) references equip_type")
	ExecDb(db, "alter table equipment add foreign key(vendor) references vendor")
	ExecDb(db, "alter table equipment add foreign key(parent_id) references equipment")

	///////////////////////////////////////////////////////////////////////////////////////
	// Vendor table
	//		- fix rating
	sqlResult, _err = SQLMap(db, "select * from fm_vendor")
	if _err != nil {
		log.Println("Error:", _err.Error())
	}
	for _, row := range sqlResult {
		// Lookup the location id
		var ratingId int
		_res, _err := SQLMap(db, "select id from vendor_rating where name=$1", row["vendorrating"])
		if _err != nil {
			log.Fatal(_err.Error())
		}
		if len(_res) > 0 {
			ratingId, _err = strconv.Atoi(_res[0]["id"])
		}

		updatedRows, _u := ExecDb(db, "update vendor set rating=$2 where name=$1",
			row["name"],
			ratingId)
		if _u != nil {
			log.Fatal(_u.Error())
		}
		ExecDb(db, "update vendor set rating=null where rating=0")
		log.Println("Updated", updatedRows, "vendor setting rating =", ratingId,
			"where name =", row["name"])
	}
	ExecDb(db, "alter table vendor add foreign key(rating) references vendor_rating")
}
