package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/api", HelloServer)
    http.ListenAndServe(":8888", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}
