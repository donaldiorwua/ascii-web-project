package asciiwebproject

import (
	"fmt"
	"net/http"
)

func HelloApp(web http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("server running...")
	fmt.Fprintf(web, "Hello World")
}
