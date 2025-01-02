package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on http://localhost:8080")

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

	// ... (rest of your code) ...

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
