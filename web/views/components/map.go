package components

import "workshop/internal/domain/sighting"

// MapPoint is one marker, handed to map.js as JSON.
type MapPoint struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Title string  `json:"title"`
	Date  string  `json:"date"`
	Place string  `json:"place"`
	Count string  `json:"count,omitzero"`
	URL   string  `json:"url,omitzero"`
}

// MapPoints turns sightings into markers, leaving out any with no position.
func MapPoints(sightings []sighting.Sighting) []MapPoint {
	points := make([]MapPoint, 0, len(sightings))
	for _, s := range sightings {
		if s.Location.Latitude == 0 && s.Location.Longitude == 0 {
			continue
		}
		points = append(points, MapPoint{
			Lat:   s.Location.Latitude,
			Lng:   s.Location.Longitude,
			Title: displayName(s.Species),
			Date:  s.HappenedAt.Format("2 Jan 2006"),
			Place: place(s.Location),
			Count: count(s.IndividualCount),
			URL:   SightingURL(s),
		})
	}
	return points
}

// GlobePoints is the ["lat","lng"] pairs the 3D globe plots, one per sighting.
func GlobePoints(sightings []sighting.Sighting) [][]float64 {
	points := make([][]float64, 0, len(sightings))
	for _, s := range sightings {
		if s.Location.Latitude == 0 && s.Location.Longitude == 0 {
			continue
		}
		points = append(points, []float64{s.Location.Latitude, s.Location.Longitude})
	}
	return points
}

// count renders how many were seen, or nothing when GBIF didn't say.
func count(n *int) string {
	if n == nil {
		return ""
	}
	return individuals(*n)
}
