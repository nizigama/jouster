package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
				"Message":   err.Error(),
				"Timestamp": time.Now().Format(time.DateTime),
			})
		},
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/analyze", analyze)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
