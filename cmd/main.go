package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", HelloServer)
	http.HandleFunc("/abc", HelloServer)
	http.HandleFunc("/api/cba", HelloServer)
	http.ListenAndServe(":8888", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
