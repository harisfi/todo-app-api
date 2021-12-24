package main

import (
	"log"
	"todo-app-api/routers"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func main() {
	r := routers.SetupRouter()
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	// run server
	log.Fatal(fasthttp.ListenAndServe(":3030", r))
}