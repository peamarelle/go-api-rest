package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Id error")
		return
	}

	for _, task := range tasks {
		if task.Id == taskId {
			fmt.Printf("%#v\n", task)
			json.NewEncoder(w).Encode(task)
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Id error")
		return
	}

	for index, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Printf("Deleted task %#v\n", task)
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var taskUpdate Task
	json.NewDecoder(r.Body).Decode(&taskUpdate)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Id error")
		return
	}

	for index, task := range tasks {
		if task.Id == taskId {
			tasks[index] = taskUpdate
			fmt.Printf("Updated task %#v\n", taskUpdate)
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
