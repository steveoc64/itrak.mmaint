package pgsql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var debug = false

func SetDebug(flag bool) {
	debug = flag
}

func Dial(dataSourceName string) error {

	if debug {
		log.Printf("SQL: Connection: %s", dataSourceName)
	}

	tempdb, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		if debug {
			log.Printf("SQL: Open Connection to %s Failed : %s\n", dataSourceName, err.Error())
		}
		return err
	}
	if debug {
		log.Println("SQL: Connected")
	}
	db = tempdb
	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

type QueryResult struct {
	Columns []string
	Rows    [][]interface{}
}

func Query(cmd string, args ...interface{}) (*QueryResult, error) {

	if debug {
		log.Println("SQL: Query:", cmd, ",", args)
	}

	rows, err := db.Query(cmd, args...)

	if err != nil {
		if debug {
			log.Println("ERROR:", err.Error())
		}
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, nil
	}

	// Create the QueryResult return value
	var retval QueryResult
	copy(retval.Columns, cols)

	// We dont know how many rows there are yet, but we know the size of each row
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		retval.Columns = append(retval.Columns, cols[i])
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Extract the Values into a row
		fields := make([]interface{}, len(cols))
		for i := 0; i < len(fields); i++ {
			fields[i] = *vals[i].(*interface{})
		}
		retval.Rows = append(retval.Rows, fields)
	}

	if rows.Err() != nil {
		return &retval, rows.Err()
	}
	return &retval, nil
}
