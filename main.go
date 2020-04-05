package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()
	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) {
		fmt.Printf("hello %s!", c.Params("name"))
	})

	app.Get("/:name/:age?", func(c *fiber.Ctx) {
		fmt.Printf("Name: %s, Age: %s",
			c.Params("name"),
			c.Params("age"),
		)
	})

	app.Get("/api*", func(c *fiber.Ctx) {
		fmt.Printf("/api%s", c.Params("*"))
	})

	app.Listen(3000)
}
