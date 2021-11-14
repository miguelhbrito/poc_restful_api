package main

import (
    "fmt"
    "net/http"

	db "github.com/stone_assignment/db_connect"
)

func main() {

	dbconnection := db.InitDB()
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	addrHttp := fmt.Sprintf(":%d", 5000)

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))git st
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}