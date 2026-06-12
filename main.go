package main

import (
	asciiwebproject "ascii-web-project/functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", asciiwebproject.HelloApp)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Not found", err)
	}
}
