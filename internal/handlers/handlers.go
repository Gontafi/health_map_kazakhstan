package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"real_time_health_map/internal/service"
	"real_time_health_map/internal/utils"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) MainPage(ctx *fiber.Ctx) error {

	stats, err := service.GetStats(ctx.Context(), h.DB, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(stats)
	return ctx.Render("index", fiber.Map{"data": stats, "meta": utils.OblastMap})
}
