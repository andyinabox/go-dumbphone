// download directions

package directions

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"googlemaps.github.io/maps"
)

const (
	travelMode = "driving"
	imgURLStr  = "https://maps.googleapis.com/maps/api/staticmap?size=%dx%d&path=enc:%s&key=%s"
	imgW, imgH = 230, 230
)

// GoogleMapsData data for google maps directions
type GoogleMapsData struct {
	Origin      string
	Destination string
	Mode        string
	Time        string
}

// Settings configuration struct
type Settings struct {
	APIKey      string
	HomeAddress string
}

var config *Settings

// Configure configure the module
func Configure(c *Settings) error {

	config = c

	return nil
}

// GetRoutes get directions for a set of directions data
func GetRoutes(data GoogleMapsData) ([]maps.Route, error) {
	if config == nil {
		return nil, errors.New("Module has not been configured")
	}

	c, err := maps.NewClient(maps.WithAPIKey(config.APIKey))
	if err != nil {
		return nil, err
	}
	r := &maps.DirectionsRequest{
		Origin:        data.Origin,
		Destination:   data.Destination,
		Mode:          maps.Mode(data.Mode),
		DepartureTime: data.Time,
	}
	routes, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

// RouteSummary summary of a google route
type RouteSummary struct {
	Duration    time.Duration
	Description string
}

// GetTransitRouteSummary get a summary of transit options
func GetTransitRouteSummary(route maps.Route) RouteSummary {

	duration := route.Legs[0].Duration
	lines := make([]string, 0)

	for _, step := range route.Legs[0].Steps {
		if step.TransitDetails != nil {
			line := step.TransitDetails.Line
			var name string

			if line.ShortName != "" {
				name = line.ShortName
			} else {
				name = line.Name
			}
			lines = append(lines, name)
		}
	}

	summary := RouteSummary{duration, strings.Join(lines, " > ")}

	return summary
}

// GetRouteSummaries get summaries of each route
func GetRouteSummaries(mode maps.Mode, routes []maps.Route) []RouteSummary {
	summaries := make([]RouteSummary, 0)

	for _, route := range routes {
		if mode == maps.TravelModeTransit {
			summaries = append(summaries, GetTransitRouteSummary(route))
		} else {
			summaries = append(summaries, RouteSummary{route.Legs[0].Duration, route.Summary})
		}
	}

	return summaries
}

// RenderRoute render a given route
func RenderRoute(route maps.Route, f *os.File) (*os.File, error) {

	t, err := template.ParseFiles("directions.html")

	if err != nil {
		return f, err
	}

	err = t.Execute(f, route)

	if err != nil {
		return f, err
	}

	return f, nil
}

// MapImageBase64 take a map polyline and encode as Base64 image
func MapImageBase64(polyline maps.Polyline) (string, error) {
	if config == nil {
		return "", errors.New("Module has not been configured")
	}

	imageURL := fmt.Sprintf(imgURLStr, imgW, imgH, polyline.Points, config.APIKey)

	// get image url
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body into bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(body)

	return str, nil
}
