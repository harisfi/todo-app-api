package main

import (
	"todo-app-api/routers"

	"github.com/valyala/fasthttp"
)

func main() {
	r := routers.SetupRouter()

	// run server
	fasthttp.ListenAndServe(":3030", r)
}