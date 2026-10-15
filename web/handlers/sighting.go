package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"uuid"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

// SightingFinder looks up a single sighting by its ID.
type SightingFinder interface {
	Find(ctx context.Context, id uuid.UUID) (sighting.Sighting, error)
}

// Sighting shows one sighting's record, or the 404 page when there is no
// sighting with that ID.
func Sighting(logger *slog.Logger, sightings SightingFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			renderError(ctx, w, logger, http.StatusNotFound, "There is no sighting with that address.")
			return
		}

		item, err := sightings.Find(ctx, id)
		if errors.Is(err, sighting.ErrNotFound) {
			renderError(ctx, w, logger, http.StatusNotFound, "There is no sighting with that address.")
			return
		} else if err != nil {
			logger.ErrorContext(ctx, "handlers: finding sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to load the sighting.")
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err = pages.Sighting(item).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to render page.")
		}
	}
}
