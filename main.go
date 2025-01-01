package main

import (
	"html/template"
	"log"
	"net/http"
)

// var templates = template.Must(template.ParseGlob("templates/*.html"))

var templates = template.Must(template.New("").ParseFiles(
	"templates/base.html",
	"templates/index.html",
	"templates/reunion.html",
	"templates/passed_away.html",
	"templates/contact.html",
))

func init() {
	log.Println("Templates parsed successfully")
}

func renderTemplate(w http.ResponseWriter, tmpl string, title string) {
	log.Printf("Rendering template: %s with title: %s", tmpl, title)
	err := templates.ExecuteTemplate(w, "base.html", map[string]interface{}{
		"Title":           title,
		"ContentTemplate": tmpl,
	})
	if err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	log.Println("Server started on http://localhost:8080")

	// Static files (CSS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			log.Println("Handling Home page")
			renderTemplate(w, "index.html", "Home")
		case "/reunion":
			log.Println("Handling 50th Class Reunion page")
			renderTemplate(w, "reunion.html", "50th Class Reunion")
		case "/passed_away":
			log.Println("Handling Passed Away page")
			renderTemplate(w, "passed_away.html", "Passed Away")
		case "/contact":
			log.Println("Handling Contact Us page")
			renderTemplate(w, "contact.html", "Contact Us")
		default:
			log.Println("404 - Page Not Found")
			http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		}
	})

	http.HandleFunc("/send_message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")
			log.Printf("Message received from %s (%s): %s", name, email, message)
			// Add logic to send an email or save the message
		} else {
			log.Println("Invalid request method")
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Example success message
		renderTemplate(w, "contact.html", "Message Sent!")
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 - Page Not Found: %s", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}
