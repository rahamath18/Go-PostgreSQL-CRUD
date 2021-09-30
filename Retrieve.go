package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	rows, err := db.Query("SELECT id,salary,code,name,additionalinformation,status,createdby,createddate,updatedby,updateddate from tbl_employee")
	CheckError(err)
	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		var id int
		var salary float32
		var code string
		var name string
		var additionalInformation string
		var status string
		var createdBy string
		var createdDate string
		var updatedBy string
		var updatedDate string

		err = rows.Scan(&id, &salary, &code, &name, &additionalInformation, &status, &createdBy, &createdDate, &updatedBy, &updatedDate)
		CheckError(err)

		fmt.Println(id, salary, code, name, additionalInformation, status, createdBy, createdDate, updatedBy, updatedDate)
	}

	CheckError(err)
	err = db.Ping()
	CheckError(err)

	fmt.Println("Retrieved successfully!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
