package main

import (
	"log"
	"todo-app-api/database"
	"todo-app-api/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	database.SetupDB()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	routers.SetupRouter(app)
	// run server
	log.Fatal(app.Listen(":3030"))
}