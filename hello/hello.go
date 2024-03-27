package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Deadline string `json:"deadline"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	// Do something with the Person struct...
	fmt.Fprintf(w, "Task: %+v", task)
	stmt, err := db.Prepare("INSERT INTO task(title, content, deadline) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(task.Title, task.Content, task.Deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, id)
}

func readTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, deadline FROM task")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Deadline)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the update operation
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the delete operation
}

func main() {
	http.HandleFunc("/task", createTask)
	http.HandleFunc("/tasks", readTasks)
	http.HandleFunc("/updatetask/", updateTask)
	http.HandleFunc("/deletetask/", deleteTask)
	http.ListenAndServe(":8080", nil)
}
