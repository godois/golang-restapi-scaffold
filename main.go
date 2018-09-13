package main

import (
	"net/http"
	"runtime"

	"github.com/godois/golang-restapi-scaffold/config"
	"github.com/godois/golang-restapi-scaffold/route"
)

func init() {
	config.LoadConfig()
}

func main() {
	if numprocs := config.C.GetInt("cpu.numprocs"); numprocs != 0 {
		runtime.GOMAXPROCS(numprocs)
	}
	api := route.InitRouter()
	strPort := config.C.GetString("server.port")
	println("Listening at " + strPort + " port")
	http.ListenAndServe(strPort, api.MakeHandler())
}
