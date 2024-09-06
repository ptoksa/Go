package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Define a simple struct to hold form data
type FormData struct {
	Name    string
	Message string
}

// Parse templates for HTML rendering
var tmpl = template.Must(template.ParseFiles("index.html"))

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the form
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// Handle form submission
		r.ParseForm()
		formData := FormData{
			Name:    r.FormValue("name"),
			Message: r.FormValue("message"),
		}
		// Display submitted data
		tmpl.Execute(w, formData)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Handler for serving static files like CSS
func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	// Serve static files (e.g., CSS)
	http.HandleFunc("/static/", staticHandler)

	// Route for the home page
	http.HandleFunc("/", homeHandler)

	// Start the server
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
