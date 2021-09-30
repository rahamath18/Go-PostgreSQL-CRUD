# Go Programming Language & PostgreSQL Database - Examples (CRUD)

The Go programming language is an open-source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

https://golang.org/doc/

## Used technology stack in this example

1. go1.17.windows-amd64.msi - Download - https://golang.org/dl/
2. go get github.com/lib/pq - (installing the sql package)

Note: Make sure GOPATH & GOROOT are in place in env for github.com/lib/pq.

## Run the following Go example one by one by for CRUD operation 

### DDL for new table: tbl_employee
ddl.aql

### Check DB Connection
$ go run .\DBConn.go

### Create Employee Data
$ go run .\Create.go

### Retrieve Employee Data
$ go run .\Retrieve.go

### Update Employee Data
$ go run .\Update.go

### Delete Employee Data
$ go run .\Delete.go
