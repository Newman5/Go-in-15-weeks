// internal/api/routes.go
package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
        })
    })

    app.Get("/hello/:name", func(c *fiber.Ctx) error {
        name := c.Params("name")

        return c.JSON(fiber.Map{
            "message": "hello " + name,
        })
    })
}