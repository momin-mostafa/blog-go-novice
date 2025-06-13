package main

import (
	coursemodel "backend/course_model"
	dbhandler "backend/db_handler"
	rootrequesthandler "backend/root_request_handler"
	userModel "backend/user"

	"fmt"
	"log"
	"net/http"
)

var port string = ":8080"

func main() {
	dbhandler.InitalizeDB()

	mux := http.NewServeMux()
	mux.Handle("/", &rootrequesthandler.RootRequestHandler{})
	mux.Handle("/user", &userModel.UserRequestHandler{})
	mux.Handle("/course", &coursemodel.CourseRequestHandler{})

	fmt.Println("Auto Migrating User")

	var err = dbhandler.GetDBPointer().AutoMigrate(&userModel.User{}, &coursemodel.Course{})
	if err != nil {
		panic("AUTO MIGRATION FAILED FOR USER")
	}

	fmt.Println("DB IS UPTO DATE AND RUNNING")

	log.Fatal(http.ListenAndServe(port, mux))
}

func init() {
	fmt.Println("Initalizing server..")
	fmt.Printf("serving at : http://localhost%s\n", port)
}
