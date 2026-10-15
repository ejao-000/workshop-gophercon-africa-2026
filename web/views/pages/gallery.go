package pages

import (
	"cmp"
	"strings"

	"workshop/internal/domain/sighting"
	"workshop/web/views/components"
)

// GalleryItem is one photograph in the gallery, flattened for the template.
type GalleryItem struct {
	URL     string
	Alt     string
	Caption string
	Animal  string
	Place   string
}

// galleryItems collects every still image from the sightings given.
func galleryItems(sightings []sighting.Sighting) []GalleryItem {
	items := make([]GalleryItem, 0)
	for _, s := range sightings {
		animal := components.Animal(s)
		place := components.Place(s.Location)
		for _, m := range s.Media {
			if m.Type != "StillImage" || m.URL == "" {
				continue
			}
			items = append(items, GalleryItem{
				URL:     m.URL,
				Alt:     cmp.Or(m.Title, m.Description, strings.Join([]string{animal, place}, ", ")),
				Caption: cmp.Or(m.Creator, m.RightsHolder, m.License),
				Animal:  animal,
				Place:   place,
			})
		}
	}
	return items
}
