package main

import (
	asciiwebproject "ascii-web-project/functions"
	"html/template"
	"net/http"
)

var temp *template.Template

func main() {
	// http.HandleFunc("/hello", asciiwebproject.HelloApp)
	// http.HandleFunc("/about", asciiwebproject.About)
	
	http.HandleFunc("/", asciiwebproject.IndexPage)

	http.ListenAndServe(":8080", nil)

}
