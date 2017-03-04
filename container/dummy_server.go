package main

import (
    "fmt"
    "net/http"
    "os"
)

const VERSION = "0"

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
	panic(err)
    }
    fmt.Fprintf(w, "<h1>Hello</h1><p>I'm running on <b>%s</b>.<br /> App version %v</p>", hostname, VERSION)
    fmt.Printf("%s\n", r.URL.Path)
}

func main() {
    fmt.Printf("server started - version %v\n", VERSION)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}
