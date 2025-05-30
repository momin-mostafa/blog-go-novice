package main

import (
	rootrequesthandler "backend/root_request_handler"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var port string = ":8080"
var db *sql.DB

func initalizeDB() *sql.DB {
	connStr := "host=localhost port=5432 user=tamim password=tamim dbname=backend sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return db
}

func main() {
	db = initalizeDB()

	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/", &rootrequesthandler.RootRequestHandler{})

	log.Fatal(http.ListenAndServe(port, mux))
}

func init() {
	fmt.Println("Initalizing server..")
	fmt.Printf("serving at : http://localhost%s\n", port)
}
