package main

import (
	"net/http"

	routes "api/router"
)

func main() {

	r := routes.NewRouter()
	http.ListenAndServe(":8080", r)

}
