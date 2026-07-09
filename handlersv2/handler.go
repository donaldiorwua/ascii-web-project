package asciiwebproject

import (
	"html/template"
	"log"
	"net/http"
)

var Temp *template.Template
var err error

type PageData struct {
	InputText string
	Banner    string
	Result    string
	Error     string
}

func IndexPage(web http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		err = Temp.ExecuteTemplate(web, "index.html", nil)
		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}
	case "POST":

		data := PageData{
			InputText: request.FormValue("inputText"),
			Banner:    request.FormValue("banner_file"),
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
			if err := Temp.ExecuteTemplate(web, "index.html", data); err != nil {
				log.Println(err)
			}
			return
		}

		err = Temp.ExecuteTemplate(web, "asciiart.html", data)

		if err != nil {
			log.Println(err)
			http.Error(web, "internal server error", http.StatusInternalServerError)
			return
		}
	case "default":
		http.Error(web, "method not allowed", http.StatusMethodNotAllowed)

	}
}

func ReLoad(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		InputText: r.URL.Query().Get("inputText"),
		Banner:    r.URL.Query().Get("banner_file"),
	}

	filedata, err := LoadBanner(data.Banner)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	data.Result, err = RenderArt(data.InputText, filedata)
	if err != nil {
		data.Error = err.Error()
		return
	}

	err = Temp.ExecuteTemplate(w, "asciiart.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
