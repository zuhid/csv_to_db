package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "csv_to_db"
)

func connectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer db.Close()

	// create table
	createTableSQL := `CREATE TABLE IF NOT EXISTS contact (
        first_name TEXT,
        last_name TEXT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully!")

	// add data
	insertData := `insert into contact (first_name, last_name) values ('John', 'Doe')`
	_, err = db.Exec(insertData)
	if err != nil {
		panic(err)
	}

	fmt.Println("data inserted successfully!")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	// records := readCsvFile("input/contact.csv")
	// fmt.Println(records)
	connectToDB()
}
