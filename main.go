package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type allTasks []Task

var tasks = allTasks{{Id: 1, Name: "Task 1"}, {Id: 2, Name: "Task 2"}, {Id: 3, Name: "Task 3"}}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)

	log.Fatal(http.ListenAndServe(":8080", router))
}
