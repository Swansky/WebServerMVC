package controller

import (
	"awesomeProject1/view"
	"net/http"
)

type ErrorViewData struct {
}

func NotFound(w http.ResponseWriter, http *http.Request) {
	data := ErrorViewData{}

	view.LoadView(w, "NotFoundTemplate", data)
}

func AccessRefused(w http.ResponseWriter, http *http.Request) {
	data := ErrorViewData{}

	view.LoadView(w, "AccessRefusedTemplate", data)
}
