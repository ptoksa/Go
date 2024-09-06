package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// Database connection string
var connStr = "postgres://petri:qwerty@localhost:5432/cruddb"

// DB connection pool
var conn *pgx.Conn

// User struct to represent data
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Initialize database connection
func initDB() {
	var err error
	conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the database.")
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err = conn.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := conn.QueryRow(context.Background(), query, params["id"]).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err = conn.Exec(context.Background(), query, user.Name, user.Email, params["id"])
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := "DELETE FROM users WHERE id = $1"
	_, err := conn.Exec(context.Background(), query, params["id"])
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
    initDB()
    defer conn.Close(context.Background()) // Updated to pass context

    r := mux.NewRouter()

    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users", getUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
