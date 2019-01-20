package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

type EMPLOYEE struct {
    ID     string
    NUMBER string
}

func main() {
    fmt.Println("aaaa")
    db, err := sql.Open("postgres", "host=172.23.0.1 port=5123 user=root password=password dbname=testdb sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

    // INSERT
    var empID string
    id := 3
    number := 4444
    err = db.QueryRow("INSERT INTO employee(emp_id, emp_number) VALUES($1,$2) RETURNING emp_id", id, number).Scan(&empID)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(empID)

    // SELECT
    rows, err := db.Query("SELECT * FROM employee")

    if err != nil {
        fmt.Println(err)
    }

    var es []EMPLOYEE
    for rows.Next() {
        var e EMPLOYEE
        rows.Scan(&e.ID, &e.NUMBER)
        es = append(es, e)
    }
    fmt.Printf("%v", es)
}