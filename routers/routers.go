package routers

import (
	"todo-app-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to API TODO"})
	})

	activity := app.Group("/activity-groups")
	activity.Get("/", controllers.GetAllActivity)
	activity.Post("/", controllers.CreateActivity)
	activity.Get("/:id", controllers.GetOneActivity)
	activity.Delete("/:id", controllers.DeleteActivity)
	activity.Patch("/:id", controllers.UpdateActivity)

	todo := app.Group("/todo-items")
	todo.Get("/", controllers.GetAllTodo)
	todo.Post("/", controllers.CreateTodo)
	todo.Get("/:id", controllers.GetOneTodo)
	todo.Delete("/:id", controllers.DeleteTodo)
	todo.Patch("/:id", controllers.UpdateTodo)
}