package routes

import (
	"cohort/controllers/students"
	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(route fiber.Router) {
	route.Post("/login", students.Login)
	//to test jwt authentication
	route.Get("/test", students.AuthRequired(), students.Test)
	route.Post("/", students.AddStudents)

}
