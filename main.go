package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":8087" // Public port: 8087

func main() {
	fmt.Printf("Server running on http://localhost%s\n", port)

	// Parse base template
	baseTemplate, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal("Error parsing base template:", err)
	}

	// Parse and clone templates
	indexTemplate, err := baseTemplate.Clone()
	if err != nil {
		log.Fatal("Error cloning base template:", err)
	}
	indexTemplate, err = indexTemplate.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error parsing index template:", err)
	}

	reunionTemplate, err := baseTemplate.Clone()
	if err != nil {
		log.Fatal("Error cloning base template:", err)
	}
	reunionTemplate, err = reunionTemplate.ParseFiles("templates/reunion.html")
	if err != nil {
		log.Fatal("Error parsing reunion template:", err)
	}

	passedAwayTemplate, err := baseTemplate.Clone()
	if err != nil {
		log.Fatal("Error cloning base template:", err)
	}
	passedAwayTemplate, err = passedAwayTemplate.ParseFiles("templates/passed_away.html")
	if err != nil {
		log.Fatal("Error parsing passed away template:", err)
	}

	contactTemplate, err := baseTemplate.Clone()
	if err != nil {
		log.Fatal("Error cloning base template:", err)
	}
	contactTemplate, err = contactTemplate.ParseFiles("templates/contact.html")
	if err != nil {
		log.Fatal("Error parsing contact template:", err)
	}

	// Static files (CSS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle menu items
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, indexTemplate, "Home", "home")
	})

	http.HandleFunc("/reunion", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, reunionTemplate, "50th Class Reunion", "reunion")
	})

	http.HandleFunc("/passed_away", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, passedAwayTemplate, "Passed Away", "passed_away")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, contactTemplate, "Contact Us", "contact")
	})

	// Handle form submissions
	http.HandleFunc("/send_message", func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("Request received: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

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
	log.Fatal(http.ListenAndServe(port, nil)) // Use log.Fatal for graceful shutdown on errors
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, title string, activePage string) {
	// log.Printf("Rendering template: %s with title: %s and active page: %s", tmpl.Name(), title, activePage)
	err := tmpl.Execute(w, map[string]interface{}{
		"Title":      title,
		"ActivePage": activePage,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Template rendering error:", err)
	}
}
