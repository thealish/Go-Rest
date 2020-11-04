package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tasks []Task

type Task struct {
	ID          string `json:"id"`
	Task        string `json:"task"`
	Finished    bool   `json:"is_finished"`
	DateCreated *Date  `json:"date_created"`
}
type Date struct {
	YearCreated  int `json:"year_created"`
	MonthCreated int `json:"month_created"`
	DayCreated   int `json:"day_created"`
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000000))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = strconv.Itoa(rand.Intn(1000000))
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(tasks)

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)

}

func main() {
	// Initia;ize the Mux router
	r := mux.NewRouter()

	tasks = append(tasks, Task{ID: "1", Task: "read book", Finished: false, DateCreated: &Date{YearCreated: 2020, MonthCreated: 11, DayCreated: 2}})
	tasks = append(tasks, Task{ID: "2", Task: "read book", Finished: false, DateCreated: &Date{YearCreated: 2020, MonthCreated: 11, DayCreated: 2}})
	tasks = append(tasks, Task{ID: "3", Task: "read book", Finished: false, DateCreated: &Date{YearCreated: 2020, MonthCreated: 11, DayCreated: 2}})

	// Create handlers

	r.HandleFunc("/api/task/{id}", getTask).Methods("GET")
	r.HandleFunc("/api/task-list", getTasks).Methods("GET")
	r.HandleFunc("/api/task-create", createTask).Methods("POST")
	r.HandleFunc("/api/task-update/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/api/task-delete/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
