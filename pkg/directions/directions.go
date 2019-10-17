// download directions

package directions

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	APIKey string
}

var config *Settings

// Configure configure the module
func Configure(c *Settings) error {

	config = c

	if config.APIKey == "" {
		return errors.New("API Key has not been configured")
	}

	return nil
}

func prepRouteForTemplate(route maps.Route) (maps.Route, error) {
	return route, nil
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
		Alternatives:  true,
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

// ToString output string version of route summary
func (r RouteSummary) ToString() string {
	return fmt.Sprintf("%s (%v)", r.Description, r.Duration)
}

// GetTransitRouteSummary get a summary of transit options
func GetTransitRouteSummary(route maps.Route) (RouteSummary, error) {

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

	return summary, nil
}

// GetRouteSummaries get summaries of each route
func GetRouteSummaries(mode maps.Mode, routes []maps.Route) ([]RouteSummary, error) {
	summaries := make([]RouteSummary, 0)

	for _, route := range routes {
		if mode == maps.TravelModeTransit {
			s, _ := GetTransitRouteSummary(route)
			summaries = append(summaries, s)
		} else {
			s := RouteSummary{route.Legs[0].Duration, route.Summary}
			summaries = append(summaries, s)
		}
	}

	return summaries, nil
}

// RenderRoute render a given route
func RenderRoute(wr io.Writer, route maps.Route) error {

	// timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// filename := fmt.Sprintf("%s%s.html", os.TempDir(), timestamp)

	t, err := template.ParseFiles("./directions.html")
	if err != nil {
		return err
	}

	// f, err := os.Create(filename)
	// if err != nil {
	// 	return err
	// }

	// defer f.Close()

	prepped, err := prepRouteForTemplate(route)

	err = t.Execute(wr, prepped)
	if err != nil {
		return err
	}

	return nil
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
