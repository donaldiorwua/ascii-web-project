package asciiwebproject

import (
	"fmt"
	"html/template"
	"net/http"
)

var Temp *template.Template
var err error

// function to render Hello World
func HelloApp(web http.ResponseWriter, request *http.Request) {
	
	fmt.Fprintf(web, "Hello Word")
}

// function to render Hello World
func About(web http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(web, "This is ASCII Art Web Golang Poject!\nIts an improvement form the ASCII Art Project")
}

// function to handle index page
func IndexPage(web http.ResponseWriter, request *http.Request) {
	err = Temp.ExecuteTemplate(web, "home.html", nil)
	if err != nil {
		http.Error(web, err.Error(), http.StatusInternalServerError)
		return
	}
}
