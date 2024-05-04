package main

import (
	"fmt"
	"net/http"

	"efficient/server"
)

const (
	PORT = 9000
	HOST = "127.0.0.1"
)

func main() {
	url := fmt.Sprintf("%s:%d", HOST, PORT)
	fmt.Printf("Server listening on %s", url)

	// serve static files
	static := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	// routes
	http.HandleFunc("/", server.Index)

	http.ListenAndServe(url, nil)
}
