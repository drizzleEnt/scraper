package scraper

import (
	"context"

	desc "github.com/drizzleent/scraper/pkg/scraper"
)

func (i *Implementation) Scrap(ctx context.Context, req *desc.ScrapUrl) (*desc.ScrapData, error) {
	err := i.service.Scrap()
	if err != nil {
		return &desc.ScrapData{
			Text: "failed",
		}, err
	}
	return &desc.ScrapData{
		Text: "secsess",
	}, nil
}
