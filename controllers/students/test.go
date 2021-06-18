package students

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Test(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)
	name := claims["auth"].(string)
	log.Println(name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ping": "pong"})
}
