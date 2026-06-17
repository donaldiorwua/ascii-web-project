package main

import (
	asciiwebproject "ascii-web-project/functions"
	"html/template"
	"net/http"
)

//var err error

func main() {
	asciiwebproject.Temp = template.Must(template.ParseGlob("templates/*.html"))
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	http.HandleFunc("/", asciiwebproject.IndexPage)
	http.HandleFunc("/render", asciiwebproject.Render)

	http.ListenAndServe(":8080", nil)

}
