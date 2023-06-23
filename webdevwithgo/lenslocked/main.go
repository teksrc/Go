package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func userHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username") // 👈 getting path param
		_, err := w.Write([]byte("Hello " + username))
		if err != nil {
			fmt.Println(err)
		}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Welcome to an awesome Go Web API!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:frankacarv@gmail.com\">frankacarv@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>
			<b>Is there a free version?</b>
			Yes! We offer a free trial for 30 days on any paid plans.
		</li>
		<li>
			<b>What are your support hours?</b>
			We have support staff answering emails 24/7, though response
			times may be a bit slower on weekends.
		</li>
		<li>
			<b>How do I contact support?</b>
			Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
		</li>
	</ul>`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}



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
