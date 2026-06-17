package asciiwebproject

import (
	"fmt"
	"html/template"
	"net/http"
)

var Temp *template.Template
var err error



// function to handle index page
func IndexPage(web http.ResponseWriter, request *http.Request) {
	err = Temp.ExecuteTemplate(web, "home.html", nil)
	if err != nil {
		http.Error(web, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Render(web http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Redirect(web, request, "/", http.StatusSeeOther)
		return
	}
	ascii := request.FormValue("inputText")

	d := struct {
		Renderer string
	}{
		Renderer: ascii,
	}
	Temp.ExecuteTemplate(web, "render.html", d)
}
