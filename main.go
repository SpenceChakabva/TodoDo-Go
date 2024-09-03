package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    _ "modernc.org/sqlite" // Import the modernc.org/sqlite driver
)

var db *gorm.DB
var err error

// Todo represents a task
type Todo struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Status string `json:"status"`
}

func main() {
    // Initialize the database
    db, err = gorm.Open(sqlite.Open("file:todos.db?cache=shared&mode=rwc"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate the schema
    db.AutoMigrate(&Todo{})

    // Initialize the router
    router := mux.NewRouter()

    // Define the routes
    router.HandleFunc("/todos", getTodos).Methods("GET")
    router.HandleFunc("/todos", createTodo).Methods("POST")
    router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
    router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

    // Start the server
    log.Fatal(http.ListenAndServe(":8000", router))
}

// Handlers

func getTodos(w http.ResponseWriter, r *http.Request) {
    var todos []Todo
    db.Find(&todos)
    json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    json.NewDecoder(r.Body).Decode(&todo)
    db.Create(&todo)
    json.NewEncoder(w).Encode(todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var todo Todo
    db.First(&todo, params["id"])
    json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var todo Todo
    db.First(&todo, params["id"])
    json.NewDecoder(r.Body).Decode(&todo)
    db.Save(&todo)
    json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var todo Todo
    db.Delete(&todo, params["id"])
    json.NewEncoder(w).Encode("Todo deleted")
}