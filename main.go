package main

import (
	asciiwebproject "ascii-web-project/handlers"
	"fmt"
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
	fmt.Println("server running at http://localhost:8080/")

	http.HandleFunc("/", asciiwebproject.IndexPage)
	http.HandleFunc("/reload", asciiwebproject.ReLoad)

	http.ListenAndServe(":8080", nil)

}
