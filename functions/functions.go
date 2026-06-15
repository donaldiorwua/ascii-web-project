package asciiwebproject

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp *template.Template

// function to render Hello World
func HelloApp(web http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("server running...")

	fmt.Fprintf(web, "Hello Word")
}

// function to render Hello World
func About(web http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(web, "This is ASCII Art Web Golang Poject!\nIts an improvement form the ASCII Art Project")
}

// function to handle index page
func IndexPage(web http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(web, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(web, nil)
}
