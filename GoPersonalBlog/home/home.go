package home

import (
	"html/template"
	"net/http"

	"github.com/anotherparsa/goProjects/dbtools"
)

type Context struct {
	Posts map[int]dbtools.Posts
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./home/template/home.html")
	Context := dbtools.GetPosts()
	template.Execute(w, Context)
}
