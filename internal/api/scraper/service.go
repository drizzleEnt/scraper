package scraper

import (
	"github.com/drizzleent/scraper/internal/service"
	desc "github.com/drizzleent/scraper/pkg/scraper"
)

type Implementation struct {
	desc.UnimplementedScraperServer

	service service.Scraper
}

func NewImplementaion(service service.Scraper) *Implementation {
	return &Implementation{
		service: service,
	}
}
