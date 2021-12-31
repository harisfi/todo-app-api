package main

import (
	"log"
	"todo-app-api/database"
	"todo-app-api/routers"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	database.SetupDB()
	r := routers.SetupRouter()
	// run server
	log.Println("app running on port 3030")
	log.Fatal(fasthttp.ListenAndServe(":3030", r))
}