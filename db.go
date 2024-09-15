package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := "user=file_user password=file dbname=file_management sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}
