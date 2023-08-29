package home

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./home/template/index.html")
	template.Execute(w, nil)
}
