package asciiwebproject

import (
	"html/template"
	"log"
	"net/http"
)

var Temp *template.Template
var err error

func IndexPage(web http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		err = Temp.Execute(web, nil)
		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}
	case "POST":
		err = request.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}

		inputText := request.FormValue("inputText")
		banner_file := request.FormValue("banner_file")

		data := struct {
			InputText string
			Banner    string
			Result    string
			Error     string
		}{
			InputText: inputText,
			Banner:    banner_file,
			Result:    "asciarts",
		}

		filedata, err := LoadBanner(data.Banner)
		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}
		data.Result, err = RenderArt(data.InputText, filedata)
		if err != nil {
			data.Error = err.Error()
		}

		err = Temp.Execute(web, data)
		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}
	case "default":
		http.Error(web, "method not allowed", http.StatusMethodNotAllowed)

	}
}

// {func ReLoad(w http.ResponseWriter, r *http.Request)  {
// 	banner := r.URL.Path
// }}