package main

import (
	"fmt"
	"net/http"

	"github.com/anotherparsa/goProjects/router"
)

func main() {
	http.HandleFunc("/", router.Routing)
	fmt.Println("running server on port 6060")
	http.ListenAndServe(":6060", nil)
}
