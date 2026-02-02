package main

import (
	"github.com/gimbalabs/cardano-go-fiber-api/internal/api"
	"github.com/gofiber/fiber/v2"
)

// This import works because the module name in go.mod is
// github.com/gimbalabs/cardano-go-fiber-api and internal/api
// is inside the same module

func main() {
    app := fiber.New()

    api.RegisterRoutes(app)

    app.Listen(":3000")
}