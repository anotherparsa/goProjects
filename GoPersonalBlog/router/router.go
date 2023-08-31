package router

import (
	"fmt"
	"net/http"

	"github.com/anotherparsa/goProjects/admin"
	"github.com/anotherparsa/goProjects/home"
)

func Routing(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		home.HomePage(w, r)
	case "/about":
		fmt.Fprint(w, "welcome to about")
	case "/admin/loginpage":
		admin.AdminLoginPageHandler(w, r)
	case "/admin/login":
		admin.AdminLoginHandler(w, r)
	case "/admin/home":
		admin.AdminHomePageHandler(w, r)
	}
}
