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
	CheckError(err)
	defer db.Close()

	updateStmt := `update tbl_employee set salary = $1,code = $2,name = $3,additionalinformation = $4,status = $5,updatedby = $6,updateddate = current_timestamp where id = $7`
	_, e := db.Exec(updateStmt, 111111.11, "emp001", "Employee 001", "{\"key1\":\"value1-updated\"}", "1", "SYSTEM", 1)
	CheckError(e)
	err = db.Ping()
	CheckError(err)
	fmt.Println("Updated successfully!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
