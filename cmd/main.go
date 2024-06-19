package main

import (
	"log"
	"os"
	"path"

	"github.com/oduortoni/odrserver/src/server"
	"github.com/oduortoni/odrserver/src/server/config"
)

const (
	HOST = "localhost"
	PORT = 13001
	STATIC_DIR = "static"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get current working directory\n")
	}

	configFileName := path.Join(rootDir, "config.json")
	configOk, config := config.Read(configFileName)
	host := HOST
	port := PORT
	staticDir := STATIC_DIR
	if configOk {
		host = config.Host
		port = config.Port
		staticDir = config.Static
	}
	staticDir = path.Join(rootDir, staticDir)

	err = server.Run(host, port, staticDir)
	if err != nil {
		log.Fatalf("[server error] %s\n", err.Error())
	} else {
		println("Running")
	}
}
