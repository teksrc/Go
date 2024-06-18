package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joncalhoun/lenslocked/controllers"
	"github.com/joncalhoun/lenslocked/views"
)

// Temporarily using one template for all handlers to avoid this being html project for now
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Parse templates
	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", r)
}
