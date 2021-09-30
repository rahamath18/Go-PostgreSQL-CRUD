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

	deleteStmt := `delete from tbl_employee where id = $1`
	_, e := db.Exec(deleteStmt, 1)
	CheckError(e)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Deleted successfully!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
