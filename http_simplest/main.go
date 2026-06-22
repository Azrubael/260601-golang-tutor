package main

import (
	"fmt"
	"net/http"
)

func httpRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User list")
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Post list")
}

func main() {
	http.HandleFunc("/root", httpRootHandler)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/posts", handlePosts)
	http.ListenAndServe(":8080", nil)
}
