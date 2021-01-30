// main.go
package main

import (
	"fmt"
	"githubapi-golang/rest"
	customTMPL "githubapi-golang/template"
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

func createHandler(writer http.ResponseWriter, request *http.Request) {
	userName := request.FormValue("username")
	gr := rest.NewGetRequest("https://api.github.com/users/")
	user := gr.GetUser(userName)
	repos := gr.GetRepos(userName)
	customTMPL.OutPutFile(user, repos)
	html, err := template.ParseFiles("html/result.html")
	logFatal(err)
	err = html.Execute(writer, nil)
	logFatal(err)
}

func main() {
	// cssを適用するために必要な設定
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html/"))))
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/create", createHandler)
	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
