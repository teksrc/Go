package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Welcome to an awesome Go Web API!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:frankacarv@gmail.com\">frankacarv@gmail.com</a></p>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, r.URL.Path) // ability to get the path off the request
// 	// fmt.Fprint(w, r.URL.Path)
// 	// fmt.Fprint(w, r.URL.RawPath)

// 	// use a switch instead
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		notFoundHandler(w, r)
// 	}
// }

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, r.URL.Path) // ability to get the path off the request
	// fmt.Fprint(w, r.URL.Path)
	// fmt.Fprint(w, r.URL.RawPath)

	// use a switch instead of if else if
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	var router Router
	// var router http.HandlerFunc = pathHandler
	// Why does Go allow both types and conversion, makes life easier for us to not have to convert over.
	// http.Handler - interface with the ServeHTTP method
	// http.HandlerFunc - function that accepts same args as ServeHTTP method, but implements http.Handler
	// http.Handle - function that takes in a pattern and takes in an http.Handler
	// http.HandleFunc - takes in a pattern and takes in a function that has underlying type looks like http.handlerFunc
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", router)
}
