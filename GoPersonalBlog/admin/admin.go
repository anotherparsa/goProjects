package admin

import (
	"html/template"
	"net/http"
)

func AdminLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./admin/template/adminlogin.html")
	template.Execute(w, nil)
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/home", 302)
}

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./admin/template/adminhome.html")
	template.Execute(w, nil)
}
