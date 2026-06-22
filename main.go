package main

import (
	asciiwebproject "ascii-web-project/handlers"
	"html/template"
	"log"
	"net/http"
)

var err error

func main() {
	asciiwebproject.Temp = template.Must(template.ParseGlob("templates/*.html"))
	if err != nil {
		log.Fatal(err)
		return
	}

	http.HandleFunc("/", asciiwebproject.IndexPage)

	http.ListenAndServe(":8080", nil)

}
