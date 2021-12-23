package main

import (
	"log"
	"todo-app-api/routers"

	"github.com/valyala/fasthttp"
)

func main() {
	r := routers.SetupRouter()

	// run server
	log.Fatal(fasthttp.ListenAndServe(":3030", r))
}