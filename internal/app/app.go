package app

import (
	"context"
	"github.com/gofiber/template/html/v2"
	"real_time_health_map/internal/db"
	"real_time_health_map/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RunApp(ctx context.Context) (*fiber.App, error) {
	database, err := db.ConnectSqlLite()
	if err != nil {
		return nil, err
	}

	err = db.UpMigrations(ctx, database)
	if err != nil {
		return nil, err
	}

	router := handlers.Handlers{
		DB: database,
	}

	app := ConfigServer()

	router.SetupRoutes(app)

	return app, nil
}

func ConfigServer() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			//DisableStartupMessage: true,
			Views: engine,
		})

	app.Static("/static", "./views")
	return app
}
