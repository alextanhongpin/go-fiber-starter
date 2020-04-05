package middleware

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
)

func Configure(app *fiber.App) {
	// authorizer := NewAuthorizer([]byte("secret"))
	// app.Use(authorizer.Middleware())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(cors.New())
	{
		cfg := limiter.Config{
			Max:     5,
			Timeout: 1,
			Key: func(c *fiber.Ctx) string {
				return c.IP()
			},
		}
		app.Use(limiter.New(cfg))
	}
}
