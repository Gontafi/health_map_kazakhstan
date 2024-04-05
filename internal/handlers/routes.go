package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) SetupRoutes(app *fiber.App) {
	app.Get("/", h.MainPage)
	//app.Get("/api/:region_id")
	//app.Post("/api/:region_id")
}
