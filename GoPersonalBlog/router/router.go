package router

import (
	"fmt"
	"net/http"
)

func Routing(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		fmt.Fprint(w, "welcome to home")
	case "/about":
		fmt.Fprint(w, "welcome to about")
	}
}
