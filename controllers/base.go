package controllers

import "todo-app-api/database/models"

type baseOutput struct {
	Status	string			`json:"status"`
	Message	string			`json:"message"`
	Data	interface{}		`json:"data"`
}

var activitiesCache []models.Activity = []models.Activity{}
var activityChanged bool = true

var todoItemsCache []models.Todo = []models.Todo{}
var todoItemChanged bool = true