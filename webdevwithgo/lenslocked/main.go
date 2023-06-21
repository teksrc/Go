package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, r.URL.Path) // ability to get the path off the request
	// fmt.Fprint(w, r.URL.Path)
	// fmt.Fprint(w, r.URL.RawPath)

	// use a switch instead
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404, Page Not Found."))
	return
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", nil)
}
