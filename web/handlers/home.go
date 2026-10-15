package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

type SightingsLister interface {
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

func Home(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return listing(logger, sightings, "sightings", pages.Home)
}
