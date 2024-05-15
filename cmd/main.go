package main

import (
	"context"
	"log"

	"github.com/drizzleent/scraper/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to init app %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Failed to run app %s", err.Error())
	}
}
