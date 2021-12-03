package view

import (
	"io"
	"text/template"
)

func LoadView(w io.Writer, templateName string, data interface{}) {

	files, err := template.ParseFiles("template/"+templateName+".html", "template/layout/menu.layout.tmpl",
		"template/layout/globalcss.layout.tmpl")
	if err != nil {
		panic(err)
		return
	}
	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
