package main

// `go install github.com/cortesi/modd/cmd/modd@latest` for Modd at global level

import (
	"fmt"
	"net/http"
	// "os"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, r.URL.Path) // ability to get the path off the request
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// 	return
	// } else if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	// 	return
	// }
	// use a switch instead
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: handle page not found error - 404
		notFoundHandler(w, r)
	}
}

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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("404, Page Not Found."))
	return
}

func main() {
	// fmt.Fprint(os.Stdout, "Hello World!")
	http.HandleFunc("/", pathHandler) // testing r.URL.Path
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", nil)
}
