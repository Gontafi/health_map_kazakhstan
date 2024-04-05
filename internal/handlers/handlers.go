package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"real_time_health_map/internal/service"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) MainPage(ctx *fiber.Ctx) error {

	stats, _ := service.GetStats(ctx.Context(), h.DB, nil)

	return ctx.Render("index", fiber.Map{"data": stats})
}
