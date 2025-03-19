package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/produk_db")
    if err != nil {
        panic(err.Error())
    }

    err = DB.Ping()
    if err != nil {
        panic(err.Error())
    }
}