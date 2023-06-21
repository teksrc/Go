package main

// `go install github.com/cortesi/modd/cmd/modd@latest` for Modd at global level

import (
	"fmt"
	"net/http"
	// "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Content Types: text/html, application/json, image/png, & video/mp4 for example.
	// https://pkg.go.dev/net/http
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Welcome to an awesome Go Web API!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:frankacarv@gmail.com\">frankacarv@gmail.com</a></p>")
}

func main() {
	// fmt.Fprint(os.Stdout, "Hello World!")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", nil)
}
