// download directions

package directions

import (
	"context"
	"log"
	"strconv"
	"text/template"
	"time"

	"googlemaps.github.io/maps"
)

const (
	travelMode = "driving"
	imgTpl     = "https://maps.googleapis.com/maps/api/staticmap?size=%dx%d&path=enc:%s&key="
	imgW, imgH = 230, 230
)

// GoogleMapsData data for google maps directions
type GoogleMapsData struct {
	APIKey      string
	Origin      string
	Destination string
	Mode        string
	Time        string
}

// Directions get directions for a set of directions data
func Directions(data GoogleMapsData) []maps.Route {
	c, err := maps.NewClient(maps.WithAPIKey(data.APIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        data.Origin,
		Destination:   data.Destination,
		Mode:          maps.Mode(data.Mode),
		DepartureTime: data.Time,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return route
}
