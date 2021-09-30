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

	insertStmt1 := `insert into tbl_employee(salary,code,name,additionalinformation,status,createdby,createddate,updatedby,updateddate) 
	values(11111.11, 'emp001', 'Employee 001', '{"key1":"value1"}', '1','SYSTEM', current_timestamp, 'SYSTEM', current_timestamp)`
	_, e1 := db.Exec(insertStmt1)
	CheckError(e1)

	insertStmt2 := `insert into tbl_employee(salary,code,name,additionalinformation,status,createdby,createddate,updatedby,updateddate) 
	values(22222.22, 'emp002', 'Employee 002', '{"key2":"value2"}', '1','SYSTEM', current_timestamp, 'SYSTEM', current_timestamp)`
	_, e2 := db.Exec(insertStmt2)
	CheckError(e2)

	insertStmt3 := `insert into tbl_employee(salary,code,name,additionalinformation,status,createdby,createddate,updatedby,updateddate) 
	values($1,$2,$3,$4,$5,$6,current_timestamp,$7,current_timestamp)`
	_, e3 := db.Exec(insertStmt3, 33333.33, "emp003", "Employee 003", "{\"key3\":\"value3\"}", "1", "SYSTEM", "SYSTEM")
	CheckError(e3)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Created successfully!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
