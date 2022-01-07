package controllers

import (
	"todo-app-api/database"
	"todo-app-api/database/models"

	"github.com/gofiber/fiber/v2"
)

func getActivityRequests(c *fiber.Ctx) (string, string) {
	type reqActivity struct {
		Title	string	`json:"title" form:"title"`
		Email	string	`json:"email" form:"email"`
	}

	ra := new(reqActivity)
	if err := c.BodyParser(ra); err != nil {
		return "", ""
	}

	return ra.Title, ra.Email
}

func findOneActivity(id string) (models.Activity, string) {
	activity := models.Activity{}
	if activityChanged {
		database.GetDB().Find(&activity, id)
	} else {
		activity = activitiesCache[activity.ID]
	}

	return activity, id
}

func fetchAllActivity() {
	database.GetDB().Find(&activitiesCache)
}

func GetAllActivity(c *fiber.Ctx) error {
	if activityChanged {
		fetchAllActivity()
		activityChanged = false
	}

	return c.JSON(&baseOutput{
		Status: "Success",
		Message: "Success",
		Data: activitiesCache,
	})
}

func GetOneActivity(c *fiber.Ctx) error {
	activity, id := findOneActivity(c.Params("id"))
	if activity.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Activity with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		return c.JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: activity,
		})
	}
}

func CreateActivity(c *fiber.Ctx) error {
	title, email := getActivityRequests(c)

	if title != "" {
		a := models.Activity{
			Title: title,
			Email: email,
		}

		database.GetDB().Create(&a)
		activityChanged = true

		return c.Status(fiber.StatusCreated).JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: a,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(&baseOutput{
			Status: "Bad Request",
			Message: "title cannot be null",
			Data: map[int]int{},
		})
	}
}

func DeleteActivity(c *fiber.Ctx) error {
	activity, id := findOneActivity(c.Params("id"))
	if activity.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Activity with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		activityChanged = true
		database.GetDB().Delete(&activity)
		return c.JSON(&baseOutput{
			Status: "Success",
			Message: "Success",
			Data: map[int]int{},
		})
	}
}

func UpdateActivity(c *fiber.Ctx) error {
	activity, id := findOneActivity(c.Params("id"))
	if activity.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&baseOutput{
			Status: "Not Found",
			Message: "Activity with ID " + id + " Not Found",
			Data: map[int]int{},
		})
	} else {
		title, _ := getActivityRequests(c)

		if title != "" {
			activity.Title = title

			database.GetDB().Updates(&activity)
			return c.JSON(&baseOutput{
				Status: "Success",
				Message: "Success",
				Data: activity,
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(&baseOutput{
				Status: "Bad Request",
				Message: "title cannot be null",
				Data: map[int]int{},
			})
		}
	}
}
