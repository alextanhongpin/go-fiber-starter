package main

import (
	"fmt"

	"github.com/gofiber/fiber"

	"github.com/alextanhongpin/go-fiber/middleware"
)

func handler(c *fiber.Ctx) {
	// fmt.Printf("hello world: %s", "hi")
	c.JSON(fiber.Map{
		"hello": "world",
	})
}

func main() {
	app := fiber.New()

	middleware.Configure(app)

	v1 := app.Group("/v1")
	v1.Get("/list", handler)

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
