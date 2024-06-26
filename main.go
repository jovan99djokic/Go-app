package main

import (
    "fmt"
    "log"
    "net/http"
)

// controller definitions
func ping(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/ping endpoint was invoked")
    fmt.Fprintf(w, "Ping pong dong")
}

func test(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/test endpoint was invoked")
    fmt.Fprintf(w, "Wot")
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/hello endpoint was invoked")
    fmt.Fprintf(w, "Hi Hay Ho")
}

func main() {
    // register the endpoints
    http.HandleFunc("/ping", ping)
    http.HandleFunc("/test", test)
    http.HandleFunc("/hello", hello)

    // start the server, listen on localhost port 8000
    fmt.Println("Jovan's server is now running extra super duper good")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
