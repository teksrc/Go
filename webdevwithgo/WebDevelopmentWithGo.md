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

Naming:
https://go.dev/talks/2014/names.slide#1

There are two main types in the net/http package:

1. http.Handler - an interface with a ServeHTTP method
2. http.HandlerFunc - a function that accepts the same args as a ServeHTTP method

There are also two functions in the http package that look similar, and are used to register their respectively similar types for a web server:

1. http.Handle - a function that accepts a pattern and an http.Handler as its arguments.
2. http.HandleFunc - a function that accepts a pattern and a function that looks like a http.HandlerFunc as its arguments.

https://go-chi.io/#/ or alternative of gorilla mux used in v1 course
https://github.com/gorilla/mux
Why use gorilla mux instead of ServeMux type?
A lot of things aren't easy to do directly, 3rd party library can handle these complications. For ex: dynamic parameters.GET /galleries/:galleryId for example Amazon does this with Id's in their paths for items where the rest of the path stays the same. So the ID Is apart of the URL. Common to have dynami variables, plus, paths can get further complex. Again, simplifies the work.

https://github.com/go-chi/chi#url-parameters Chi url parameters

Helpful little article for Restful routing with Chi in Go:
https://thedevelopercafe.com/articles/restful-routing-with-chi-in-go-d05a2f952b3d

An Introduction to Templates in Go
https://www.calhoun.io/intro-to-templates/
https://go.dev/doc/articles/wiki/
https://pkg.go.dev/text/template#hdr-Actions
https://html.spec.whatwg.org/multipage/named-characters.html
https://www.rapidtables.com/web/html/html-codes.html

Alternative structures to MVC in Go
https://www.gobeyond.dev
Kat Zien's 2018 gophercon talk on YT https://www.youtube.com/watch?v=oL6JBUk6tj0
https://blog.jetbrains.com/go/2023/04/11/catching-up-with-kat-zien-on-the-structure-of-go-apps-in-2023/

Use cases for the Go language include:

- Cloud applications.
- Web development.
- Database implementations.
- Distributed networking services.
- Utilities.
- IoT devices.
