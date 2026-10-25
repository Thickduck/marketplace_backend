package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "hi guys backend is connecteddddddd",
		})
	})

	app.Post("/api/login", func(c *fiber.Ctx) error {
		payload := c.Body()
		fmt.Println("API Gateway reached")
		fmt.Println(payload)

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "endpoint working successfully",
			"payload": string(payload),
		})
	})

	app.Listen(":8080")
}
