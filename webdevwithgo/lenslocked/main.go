package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joncalhoun/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username") // 👈 getting path param
		_, err := w.Write([]byte("Hello " + username))
		if err != nil {
			fmt.Println(err)
		}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.html")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.html")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.html")
	executeTemplate(w, tplPath)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

// Temporarily using one template for all handlers to avoid this being html project for now
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/user/{username}", userHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notFoundHandler)
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", r)
}
