package app

import (
	"context"
	"log"

	"github.com/drizzleent/scraper/internal/api/scraper"
	"github.com/drizzleent/scraper/internal/config"
	"github.com/drizzleent/scraper/internal/config/env"
	"github.com/drizzleent/scraper/internal/service"
	scraperService "github.com/drizzleent/scraper/internal/service/scraper"
	"github.com/gocolly/colly"
)

const (
	userAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	collector *colly.Collector

	service service.Scraper

	implementation *scraper.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if nil == s.grpcConfig {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) Collector() *colly.Collector {
	if nil == s.collector {
		s.collector = colly.NewCollector(colly.UserAgent(userAgent))
	}

	return s.collector
}

func (s *serviceProvider) Service(ctx context.Context) service.Scraper {
	if nil == s.service {
		s.service = scraperService.NewScraper(s.Collector())
	}

	return s.service
}
func (s *serviceProvider) Implementation(ctx context.Context) *scraper.Implementation {
	if nil == s.implementation {
		s.implementation = scraper.NewImplementaion(s.Service(ctx))
	}

	return s.implementation
}
