package server

import (
	"fmt"
	"net/http"
	"github.com/oduortoni/odrserver/src/server/controllers"
)

func Run(host string, port int, staticDir string) error {
	serverUrl := fmt.Sprintf("%s:%d", host, port)

	staticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))

	fmt.Printf("\n\n\t---------------------------------------------------------\n\n\n\t\tserver listening at %q\n\n\n\t---------------------------------------------------------\n", serverUrl)
	
	http.HandleFunc("/", controllers.Index)
	return http.ListenAndServe(serverUrl, nil)
}