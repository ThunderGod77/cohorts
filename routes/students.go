package routes

import (
	"cohort/controllers/students"
	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(route fiber.Router) {
	route.Post("/", students.AddStudents)
	route.Post("/", students.Login)

	route.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"ping": "pong"})
	})

}
