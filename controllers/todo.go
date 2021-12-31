package controllers

import (
	"fmt"
	"strconv"
	"todo-app-api/database"
	"todo-app-api/database/models"
	"todo-app-api/services"

	"github.com/valyala/fasthttp"
)

func getTodoRequests(ctx *fasthttp.RequestCtx) (string, uint, string) {
	req := services.ExtractRequests(ctx)

	var (
		activityGroupId uint
		title, priority string
	)

	if req["title"] == nil {
		ctx.SetStatusCode(400)
		services.SendJSONResponse(ctx, nil, "Bad Request", "title cannot be null")
		return "", 0, ""
	}
	title = req["title"].(string)

	if req["activity_group_id"] == nil {
		ctx.SetStatusCode(400)
		services.SendJSONResponse(ctx, nil, "Bad Request", "activity_group_id cannot be null")
		return "", 0, ""
	}
	activityGroupId = uint(req["activity_group_id"].(float64))

	if req["priority"] == nil {
		priority = "very-high"
	} else {
		priority = req["priority"].(string)
	}

	return title, activityGroupId, priority
}

func findOneTodo(ctx *fasthttp.RequestCtx) (models.Todo, interface{}) {
	id := ctx.UserValue("id")
	todo := models.Todo{}

	database.GetDB().Find(&todo, id)
	return todo, id
}

func GetAllTodo(ctx *fasthttp.RequestCtx) {
	todoItems := []models.Todo{}
	if activityGroupId := ctx.FormValue("activity_group_id"); activityGroupId != nil {
		activityGroupIdx, _ := strconv.Atoi(string(activityGroupId))
		database.GetDB().Find(&todoItems, models.Todo{ActivityGroupId: uint(activityGroupIdx)})
	} else {
		database.GetDB().Find(&todoItems)
	}
	services.SendJSONResponse(ctx, todoItems, "", "")
}

func GetOneTodo(ctx *fasthttp.RequestCtx) {
	todo, id := findOneTodo(ctx)
	if todo.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Todo with ID " + id.(string) + " Not Found")
	} else {
		services.SendJSONResponse(ctx, todo, "", "")
	}
}

func CreateTodo(ctx *fasthttp.RequestCtx) {
	title, activityGroupId, priority := getTodoRequests(ctx)

	if title != "" && activityGroupId != 0 && priority != "" {
		a := models.Activity{}
		database.GetDB().Find(&a, activityGroupId)
		if a.ID == 0 {
			ctx.SetStatusCode(404)
			services.SendJSONResponse(ctx, nil, "Not Found", fmt.Sprintf("Activity with activity_group_id %d Not Found", activityGroupId))
		} else {
			t := models.Todo{
				Title: title,
				ActivityGroupId: activityGroupId,
				Priority: priority,
				IsActive: true,
			}

			database.GetDB().Create(&t)
			services.SendJSONResponse(ctx, t, "", "")
			ctx.SetStatusCode(201)
		}
	}
}

func DeleteTodo(ctx *fasthttp.RequestCtx) {
	todo, id := findOneTodo(ctx)
	if todo.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Todo with ID " + id.(string) + " Not Found")
	} else {
		database.GetDB().Delete(&todo)
		services.SendJSONResponse(ctx, nil, "", "")
	}
}

func UpdateTodo(ctx *fasthttp.RequestCtx) {
	todo, id := findOneTodo(ctx)
	if todo.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Todo with ID " + id.(string) + " Not Found")
	} else {
		req := services.ExtractRequests(ctx)

		if req["title"] != nil {
			todo.Title = req["title"].(string)
		}

		if req["is_active"] != nil {
			todo.IsActive = req["is_active"].(bool)
		}

		if req["priority"] != nil {
			todo.Priority = req["priority"].(string)
		}

		database.GetDB().Updates(&todo)
		services.SendJSONResponse(ctx, todo, "", "")
	}
}
