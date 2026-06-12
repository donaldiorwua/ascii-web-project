package asciiwebproject

import (
	"fmt"
	"net/http"
	"strings"
)

// function to render Hello World
func HelloApp(web http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("server running...")

	// range over the request form and print its content to the terminal
	for key, value := range request.Form{
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", strings.Join(value, " "))
	}
	
	fmt.Fprintf(web, "Hello World")
}
