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

//ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
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
	// Clear old tasks

	_, err = ExecDb(db, "delete from tasks")
	_, err = ExecDb(db, "delete from workorders")
	taskBySite("Minto", 4)
	taskBySite("Tomago", 5)
	taskBySite("Chinderah", 6)
	taskBySite("Newcastle", 3)
	taskBySite("Edinbur", 2)
	taskBySite("USA", 10)
	taskBySite("Thermoloc", 8)
}

func taskBySite(sitename string, siteid int) {

	sqlResult, _err := SQLMap(db,
		`select taskname,description,to_char(startdate,'YYYY-MM-DD') as startdate,to_char(starttime,'HH:MI:SS') as starttime,
			duration,labour_cost,material_cost,othercost,
			priority,category,class,taskfrequency,instructions,
			lineno 
		from fm_task where taskname like '%`+sitename+`%' order by taskname`)

	if _err != nil {
		log.Println("Error:", _err.Error())
	}
	gotRows := 1
	for i, row := range sqlResult {
		if i%2 == 0 {
			// Strip the \b\ out of the taskname
			x := strings.Replace(row["taskname"], `\b\`, ``, 1)
			log.Println("Task", gotRows, x, row["startdate"], row["starttime"])
			//log.Println("Task", gotRows, x, row["startdate"])
			row["taskname"] = x
			gotRows++

			var TaskID int

			insertErr := db.QueryRow(`insert into tasks (
					taskname,
					description,
					startdate,
					starttime,
					duration,
					labour_cost,
					material_cost,
					othercost,
					priority,
					category,
					class,
					taskfrequency,
					instructions,
					lineno,
					site) 
				values (
					$1,
					$2,
					to_date($3, 'YYYY-MM-DD'),
					$4,
					$5,
					$6,
					$7,
					$8,
					$9,
					$10,
					$11,
					$12,
					$13,
					$14,
					$15) returning id`,
				row["taskname"],
				row["description"],
				ToNullString(row["startdate"]),
				ToNullString(row["starttime"]),
				ToNullString(row["duration"]),
				ToNullString(row["labour_cost"]),
				ToNullString(row["material_cost"]),
				ToNullString(row["othercost"]),
				row["priority"],
				row["category"],
				row["class"],
				row["taskfrequency"],
				row["instructions"],
				ToNullString(row["lineno"]),
				siteid).Scan(&TaskID)

			if insertErr != nil {
				log.Fatal(insertErr.Error())
			}

			log.Println("Task ID is ", TaskID)

			// Now have a look at workorders that are related to this task
			workorders, selErr := SQLMap(db, `select 
				wow,taskname,descr,
				to_char(planned_date,'YYYY-MM-DD'),
				to_char(start_time, 'HH:MI:SS'),
				duration,labour_cost,material_cost,other_cost,
				status,has_problem,
				completed_date,
				actual_duration,actual_labour_cost,actual_material_cost,actual_other_cost,
				feedback,priority,category,class,frequency
				from fm_workorder where taskname=$1`, row["taskname"])
			if selErr != nil {
				log.Fatal(selErr.Error())
			}
			for wi, wrow := range workorders {
				log.Println("WorkOrder:", wi, wrow)

				_, insErr := ExecDb(db,
					`insert into workorders (
					wow,
					taskname,
					descr,
					planned_date,
					start_time,
					duration,
					labour_cost,
					material_cost,
					other_cost,
					status,
					has_problem,
					completed_date,
					actual_duration,
					actual_labour_cost,
					actual_material_cost,
					actual_other_cost,
					feedback,
					priority,
					category,
					class,
					frequency,
					task_id,
					site_id) 
				values (
					$1,
					$2,
					$3,
					to_date($4, 'YYYY-MM-DD'),
					$5,
					$6,
					$7,
					$8,
					$9,
					$10,
					$11,
					to_date($12,'YYYY-MM-DD'),
					$13,
					$14,
					$15,
					$16,
					$17,
					$18,
					$19,
					$20,
					$21,
					$22,
					$23)`,
					wrow["wow"],
					wrow["taskname"],
					wrow["description"],
					ToNullString(wrow["planned_date"]),
					ToNullString(wrow["starttime"]),
					ToNullString(wrow["duration"]),
					ToNullString(wrow["labour_cost"]),
					ToNullString(wrow["material_cost"]),
					ToNullString(wrow["other_cost"]),
					wrow["status"],
					wrow["has_problem"],
					ToNullString(wrow["completed_date"]),
					ToNullString(wrow["actual_duration"]),
					ToNullString(wrow["actual_labour_cost"]),
					ToNullString(wrow["actual_material_cost"]),
					ToNullString(wrow["actual_other_cost"]),
					wrow["feedback"],
					wrow["priority"],
					wrow["category"],
					wrow["class"],
					wrow["frequency"],
					TaskID, siteid)

				if insErr != nil {
					log.Fatal(insErr.Error())
				}

			}

		}
	}
}
