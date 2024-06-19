package view

import (
	"html/template"
	"io"
)

func Send(w io.Writer, name string, data *any) bool {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return false
	}
	tmpl.Execute(w, data)
	return true
}
