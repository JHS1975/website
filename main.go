package main

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	log.Println("Templates parsed successfully")
}

func main() {
	log.Println("Server started on http://localhost:8080")

	// Parse base template
	baseTemplate, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal("Error parsing base template:", err)
	}

	// Parse and clone templates
	indexTemplate, err := baseTemplate.Clone().ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error parsing index template:", err)
	}

	reunionTemplate, err := baseTemplate.Clone().ParseFiles("templates/reunion.html")
	if err != nil {
		log.Fatal("Error parsing reunion template:", err)
	}

	passedAwayTemplate, err := baseTemplate.Clone().ParseFiles("templates/passed_away.html")
	if err != nil {
		log.Fatal("Error parsing passed away template:", err)
	}

	contactTemplate, err := baseTemplate.Clone().ParseFiles("templates/contact.html")
	if err != nil {
		log.Fatal("Error parsing contact template:", err)
	}

	// Static files (CSS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle menu items
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, indexTemplate, "Home")
	})

	http.HandleFunc("/reunion", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, reunionTemplate, "50th Class Reunion")
	})

	http.HandleFunc("/passed_away", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, passedAwayTemplate, "Passed Away")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, contactTemplate, "Contact Us")
	})

	// Handle form submissions
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
		err := contactTemplate.Execute(w, "Message Sent!")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil)) // Use log.Fatal for graceful shutdown on errors
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, title string) {
	err := tmpl.Execute(w, map[string]interface{}{
		"Title":      title,
		"ActivePage": tmpl.Name(),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
