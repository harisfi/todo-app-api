package controllers

import (
	"fmt"
	"strconv"
	"todo-app-api/database"
	"todo-app-api/database/models"

	"github.com/gofiber/fiber/v2"
)

func getTodoRequests(c *fiber.Ctx) (string, uint, string, bool) {
	type reqTodo struct {
		Title				string	`json:"title" form:"title"`
		ActivityGroupId		uint	`json:"activity_group_id" form:"activity_group_id"`
		Priority			string	`json:"priority" form:"priority"`
		IsActive			bool	`json:"is_active" form:"is_active"`
	}
	
	rt := new(reqTodo)
	if err := c.BodyParser(rt); err != nil {
		return "", 0, "", true
	}

	if rt.Priority == "" {
		rt.Priority = "very-high"
	}

	return rt.Title, rt.ActivityGroupId, rt.Priority, rt.IsActive
}

func findOneTodo(c *fiber.Ctx) (models.Todo, string) {
	id := c.Params("id")
	todo := models.Todo{}
	if todoItemChanged {
		database.GetDB().Find(&todo, id)
	} else {
		todo = todoItemsCache[todo.ID]
	}

	return todo, id
}

func fetchAllTodo() {
	database.GetDB().Find(&todoItemsCache)
}

func GetAllTodo(c *fiber.Ctx) error {
	if todoItemChanged {
		fetchAllTodo()
		todoItemChanged = false
	}

	todoItems := []models.Todo{}
	if activityGroupId := c.Query("activity_group_id"); activityGroupId != "" {
		activityGroupIdx, _ := strconv.Atoi(string(activityGroupId))
		for _, t := range todoItemsCache {
			if int(t.ActivityGroupId) == activityGroupIdx {
				todoItems = append(todoItems, t)
			}
		}
	} else {
		todoItems = todoItemsCache
	}

	return c.JSON(&baseOutput{
		Status: "Success",
		Message: "Success",
		Data: todoItems,
	})
}

func GetOneTodo(c *fiber.Ctx) error {
	todo, id := findOneTodo(c)
	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Todo with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		return c.JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: todo,
		})
	}
}

func CreateTodo(c *fiber.Ctx) error {
	title, activityGroupId, priority, _ := getTodoRequests(c)

	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&baseOutput{
			Status: "Bad Request",
			Message: "title cannot be null",
			Data: map[int]int{},
		})
	}

	if activityGroupId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&baseOutput{
			Status: "Bad Request",
			Message: "activity_group_id cannot be null",
			Data: map[int]int{},
		})
	}

	a := models.Activity{}
	database.GetDB().Find(&a, activityGroupId)
	if a.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: fmt.Sprintf("Activity with activity_group_id %d Not Found", activityGroupId),
			Data: map[int]int{},
		})
	} else {
		t := models.Todo{
			Title: title,
			ActivityGroupId: activityGroupId,
			Priority: priority,
			IsActive: true,
		}

		database.GetDB().Create(&t)
		todoItemChanged = true

		return c.Status(fiber.StatusCreated).JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: t,
		})
	}
}

func DeleteTodo(c *fiber.Ctx) error {
	todo, id := findOneTodo(c)
	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Todo with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		database.GetDB().Delete(&todo)
		todoItemChanged = true
		return c.JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: map[int]int{},
		})
	}
}

func UpdateTodo(c *fiber.Ctx) error {
	todo, id := findOneTodo(c)
	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Todo with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		title, _, priority, isActive := getTodoRequests(c)

		if title != "" {
			todo.Title = title
		}

		todo.IsActive = isActive

		if priority != "" {
			todo.Priority = priority
		}

		database.GetDB().Updates(&todo)
		todoItemChanged = true
		return c.JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: todo,
		})
	}
}
