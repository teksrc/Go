package main

// `go install github.com/cortesi/modd/cmd/modd@latest` for Modd at global level

import (
	"fmt"
	"net/http"
	// "os"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// Content Types: text/html, application/json, image/png, & video/mp4 for example.
	// https://pkg.go.dev/net/http
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	fmt.Fprint(w, "<h1>Welcome to an awesome Go Web API!</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	// fmt.Fprint(os.Stdout, "Hello World!")
	fmt.Println("Starting server on http://localhost:3000...")
	http.ListenAndServe(":3000", nil)
}
