package routers

import (
	"encoding/json"
	"todo-app-api/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func SetupRouter() func(ctx *fasthttp.RequestCtx) {
	var (
		route1 = "/activity-groups"
		route2 = "/todo-items"
	)

	router := fasthttprouter.New()

	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		m, _ := json.Marshal(map[string]string{"message": "Welcome to API TODO"})
		ctx.SetBody(m)
	})

	router.GET(route1, controllers.GetAllActivity)
	router.POST(route1, controllers.CreateActivity)
	router.GET(route1 + "/:id", controllers.GetOneActivity)
	router.DELETE(route1 + "/:id", controllers.DeleteActivity)
	router.PATCH(route1 + "/:id", controllers.UpdateActivity)

	router.GET(route2, controllers.GetAllTodo)
	router.POST(route2, controllers.CreateTodo)
	router.GET(route2 + "/:id", controllers.GetOneTodo)
	router.DELETE(route2 + "/:id", controllers.DeleteTodo)
	router.PATCH(route2 + "/:id", controllers.UpdateTodo)

	return router.Handler
}