# Networking with HTTP

 - Go Network libraries
     
     The Go standard library has many packages for making web servers:

     That includes:
        - client & server sockets
        - route multiplexing
        - HTTP and HTML, incluing HTML templated
        - JSON and other data formats
        - cryptographic security.
        - SQL database and access
        - compression utilities
        - image generations

     There are also third party packages with improvements.

- Basic HTTP server

    package main

    import (
        "fmt"
        "log"
        "net/http"
    )

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World! from %s", r.URL.Path[1:])
    }

    func main() {

        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":8000", nil))

    }


# Go HTTP design

An HTTP handler functions ia an instance of an interface

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

    type HandlerFunc func(ResponseWriter, *Request)

    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }

    // The HTTP framework can call a method on function type

    func handler(w ResponseWriter,  r *http.Request) {
        fmt.Fprintf(w, "Hello, World! from %s\n", r.URL.Path[1:])
    }