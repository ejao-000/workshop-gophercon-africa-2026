package handlers

import (
	"log/slog"
	"net/http"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"

	"github.com/a-h/templ"
)

// Map shows every sighting pinned on a map.
func Map(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return listing(logger, sightings, "map", pages.Map)
}

// Gallery shows every photograph recorded with the sightings.
func Gallery(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return listing(logger, sightings, "gallery", pages.Gallery)
}

// Stats shows the tallies behind the log.
func Stats(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return listing(logger, sightings, "stats", pages.Stats)
}

// listing renders a page built from every sighting in the log.
func listing(logger *slog.Logger, sightings SightingsLister, name string, page func([]sighting.Sighting) templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		items, err := sightings.List(ctx, sighting.ListFilter{})
		if err != nil {
			logger.ErrorContext(ctx, "handlers: listing sightings for "+name+": "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to load sightings.")
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err = page(items).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering "+name+": "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to render page.")
		}
	}
}
