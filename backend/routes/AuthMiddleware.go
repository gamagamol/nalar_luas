package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware() fiber.Handler{
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess,err:=store.Get(c)
	if err!=nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":fiber.StatusUnauthorized,
			"Message":"Unauthorized",

		})

		
	}

	if c.Cookies("sessionId") == ""{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":fiber.StatusUnauthorized,
			"Message":"Unauthorized",

		})
	}


	if sess.Get("sessionId") == nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":fiber.StatusUnauthorized,
			"Message":"Unauthorized",

		})
		}

	return c.Next()
}