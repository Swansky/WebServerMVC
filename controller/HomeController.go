package controller

import (
	"net/http"
	"text/template"
)

type HomeViewData struct {
	Title string

}

func Home(w http.ResponseWriter, http *http.Request) {
	data := HomeViewData{Title: "Je suis un titre"}
	files, err := template.ParseFiles("template/HomeTemplate.html")
	if err != nil {
		panic(err)
		return
	}
	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


