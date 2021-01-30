// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("html/index.html")
	logFatal(err)
	err = html.Execute(writer, nil)
	logFatal(err)
}

func main() {
	// cssを適用するために必要な設定
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html/"))))
	http.HandleFunc("/view", viewHandler)
	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
