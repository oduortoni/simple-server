package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	content, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(rw, "%s", "<html><head><title>404</title><body><h1>404</h1><p>An error occured</p></body></html>")
	} else {
		content.Execute(rw, nil)
	}
}
