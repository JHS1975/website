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
		"Title": title,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template error:", err)
	}
}

// func renderTemplate(w http.ResponseWriter, tmpl string, title string) {
// 	log.Printf("Rendering template: %s with title: %s", tmpl, title)
// 	err := templates.ExecuteTemplate(w, "base.html", map[string]interface{}{
// 		"Title": title,
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Println("Template error:", err)
// 	}
// }

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
			notFoundHandler(w, r)
		}
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 - Page Not Found: %s", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}
