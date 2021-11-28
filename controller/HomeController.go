package controller

import (
	"awesomeProject1/view"
	"net/http"
)

type HomeViewData struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := HomeViewData{Title: "Je suis un titre"}

	view.LoadView(w, "HomeTemplate", data)
}
