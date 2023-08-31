package admin

import (
	"html/template"
	"net/http"
)

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./admin/template/adminlogin.html")
	template.Execute(w, nil)
}
