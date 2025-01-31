package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func csvToDb(tableName string, filePath string, psqlInfo string) {
	// read csv file
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(`Error opening csv file: `, err)
	}
	defer f.Close()

	// create db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(`Error creating connection: `, err)
	}
	defer db.Close()

	// read csv header line
	reader := csv.NewReader(f)

	// create table
	line, err := reader.Read()
	if err != nil {
		fmt.Println(`Error reading csv: `, err)
	}

	_, err = db.Exec(`create table if not exists ` + tableName + `("` + strings.Join(line, `" text, "`) + `" text);`)
	if err != nil {
		fmt.Println(`Error creating table: `, err)
	}
	_, err = db.Exec(`delete from ` + tableName + `;`)
	if err != nil {
		fmt.Println(`Error deleting table: `, err)
	}

	// load the data
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading data: ", err)
			return
		}
		// insert data
		for i, str := range record {
			record[i] = strings.Replace(str, `'`, `&#39;`, -1)
		}
		script := `insert into ` + tableName + ` values ('` + strings.Join(record, `','`) + `')`

		_, err = db.Exec(script)
		if err != nil {
			fmt.Println(`Error inserting data: `+script, err)
		}
	}
}

func main() {
	psqlInfo := `host=localhost port=5432 user=postgres password=P@ssw0rd dbname=csv_to_db sslmode=disable`
	// csvToDb(`snyk`, `input/snyk_issues_detail_01_30_2025_ee2b4003-d196-4160-9a5a-d70781c04519.csv`, psqlInfo)
	csvToDb(`okta`, `input/syslog_query.csv`, psqlInfo)
}
