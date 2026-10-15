package handlers

import (
	"log/slog"
	"net/http"
	"workshop/web/views/pages"
)

// NotFound answers any address that isn't one of the site's pages.
func NotFound(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)

		if err := pages.NotFound().Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering not found page: "+err.Error())
		}
	}
}
