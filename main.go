package main

import (
	"net/http"

	"github.com/godois/golang-restapi-scaffold/route"
)

func main() {
	api := route.InitRouter()
	http.ListenAndServe(":8080", api.MakeHandler())
}
