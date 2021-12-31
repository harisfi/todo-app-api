package controllers

import (
	"todo-app-api/database"
	"todo-app-api/database/models"
	"todo-app-api/services"

	"github.com/valyala/fasthttp"
)

func getActivityRequests(ctx *fasthttp.RequestCtx) (string, string) {
	req := services.ExtractRequests(ctx)

	var title, email string

	if req["title"] == nil {
		ctx.SetStatusCode(400)
		services.SendJSONResponse(ctx, nil, "Bad Request", "title cannot be null")
		return "", ""
	}
	title = req["title"].(string)

	if req["email"] == nil {
		email = ""
	} else {
		email = req["email"].(string)
	}

	return title, email
}

func findOneActivity(ctx *fasthttp.RequestCtx) (models.Activity, interface{}) {
	id := ctx.UserValue("id")
	activity := models.Activity{}

	database.GetDB().Find(&activity, id)
	return activity, id
}

func GetAllActivity(ctx *fasthttp.RequestCtx) {
	activities := []models.Activity{}
	database.GetDB().Find(&activities)
	services.SendJSONResponse(ctx, activities, "", "")
}

func GetOneActivity(ctx *fasthttp.RequestCtx) {
	activity, id := findOneActivity(ctx)
	if activity.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Activity with ID " + id.(string) + " Not Found")
	} else {
		services.SendJSONResponse(ctx, activity, "", "")
	}
}

func CreateActivity(ctx *fasthttp.RequestCtx) {
	title, email := getActivityRequests(ctx)

	if title != "" {
		a := models.Activity{
			Title: title,
			Email: email,
		}
	
		database.GetDB().Create(&a)
		services.SendJSONResponse(ctx, a, "", "")
		ctx.SetStatusCode(201)
	}
}

func DeleteActivity(ctx *fasthttp.RequestCtx) {
	activity, id := findOneActivity(ctx)
	if activity.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Activity with ID " + id.(string) + " Not Found")
	} else {
		database.GetDB().Delete(&activity)
		services.SendJSONResponse(ctx, nil, "", "")
	}
}

func UpdateActivity(ctx *fasthttp.RequestCtx) {
	activity, id := findOneActivity(ctx)
	if activity.ID == 0 {
		ctx.SetStatusCode(404)
		services.SendJSONResponse(ctx, nil, "Not Found", "Activity with ID " + id.(string) + " Not Found")
	} else {
		title, email := getActivityRequests(ctx)

		if title != "" {
			activity.Title = title
			activity.Email = email

			database.GetDB().Updates(&activity)
			services.SendJSONResponse(ctx, activity, "", "")
		}
	}
}
