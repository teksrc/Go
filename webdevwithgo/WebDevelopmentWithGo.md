# Web Development with Golang

https://courses.calhoun.io/courses/cor_wdv2

Common Causes and Fixes for Errors in Go
https://errorsingo.com/compliation-no-new-variables-on-left-side-of/

`go install github.com/cortesi/modd/cmd/modd@latest` for Modd at global level

Content Types: text/html, application/json, image/png, & video/mp4 for example.
https://pkg.go.dev/net/http
https://pkg.go.dev/net/url#URL encoding.decoding raw path
https://meyerweb.com/eric/tools/dencoder/ helpful
https://pkg.go.dev/net/http#pkg-constants

There are two main types in the net/http package:

1. http.Handler - an interface with a ServeHTTP method
2. http.HandlerFunc - a function that accepts the same args as a ServeHTTP method

There are also two functions in the http package that look similar, and are used to register their respectively similar types for a web server:

1. http.Handle - a function that accepts a pattern and an http.Handler as its arguments.
2. http.HandleFunc - a function that accepts a pattern and a function that looks like a http.HandlerFunc as its arguments.
