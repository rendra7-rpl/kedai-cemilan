package main

import (
    "kedai_cemilan/database"
    "kedai_cemilan/handlers"
    "net/http"
)

func main() {
    database.InitDB()
    defer database.DB.Close()

    http.HandleFunc("/", handlers.IndexHandler(database.DB))

    http.ListenAndServe(":8080", nil)
}