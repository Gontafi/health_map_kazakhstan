package main

import (
	"context"
	"log"
	"real_time_health_map/internal/app"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	srv, err := app.RunApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Listen(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
