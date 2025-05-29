package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var port string = ":8080"
var db *sql.DB

func rootRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGET(w)
	case http.MethodPost:
		handlePOST(w)
	}
}

func handleGET(w http.ResponseWriter) {
	fmt.Fprintf(w, "Hello, world from get method")
}

func handlePOST(w http.ResponseWriter) {

	fmt.Fprintf(w, "Hello, world from POST")
}

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
	mux.HandleFunc("/", rootRequestHandler)

	log.Fatal(http.ListenAndServe(port, mux))
}

func init() {
	fmt.Println("Initalizing server..")
	fmt.Printf("serving at : http://localhost%s\n", port)
}
