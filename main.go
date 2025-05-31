package main

import (
	dbhandler "backend/db_handler"
	rootrequesthandler "backend/root_request_handler"

	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var port string = ":8080"

func main() {
	dbhandler.InitalizeDB()

	mux := http.NewServeMux()
	mux.Handle("/", &rootrequesthandler.RootRequestHandler{})

	log.Fatal(http.ListenAndServe(port, mux))
}

func init() {
	fmt.Println("Initalizing server..")
	fmt.Printf("serving at : http://localhost%s\n", port)
}
