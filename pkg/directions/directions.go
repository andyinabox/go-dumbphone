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
	imgURLStr            = "https://maps.googleapis.com/maps/api/staticmap?size=%dx%d&path=enc:%s&key=%s"
	imgW, imgH           = 230, 230
	templateFile         = "./directions.html"
	transitLineDelimiter = " > "
)

// RouteSummary contains summary information about a route
type RouteSummary struct {
	Duration    time.Duration
	Description string
}

// ToString output string version of route summary
func (r RouteSummary) ToString() string {
	return fmt.Sprintf("%s (%v)", r.Description, r.Duration)
}

// Trip creates an instance of a directions request. After creating a new Trip,
// be sure to use the `Fetch` method to retrieve possible routes
type Trip struct {
	APIKey      string
	Origin      string
	Destination string
	Time        string
	Mode        string

	Routes    []maps.Route
	Summaries []RouteSummary
}

// Fetch sends a request to the Google Maps API for directions, then populates
// the `Routes` key with the result.
func (t *Trip) Fetch() error {

	// check for required fields
	err := t.checkRequirements()
	if err != nil {
		return err
	}

	// create new maps client
	client, err := maps.NewClient(maps.WithAPIKey(t.APIKey))
	if err != nil {
		return err
	}

	// create request
	req := &maps.DirectionsRequest{
		Origin:        t.Origin,
		Destination:   t.Destination,
		Mode:          maps.Mode(t.Mode),
		DepartureTime: t.Time,
		Alternatives:  true, // provide alt routes
	}

	// get routes
	routes, _, err := client.Directions(context.Background(), req)
	if err != nil {
		return err
	}

	// assign routes
	t.Routes = routes

	// generate route summaries
	t.Summaries, err = generateSummaries(routes, req.Mode)
	if err != nil {
		return err
	}

	return nil
}

// Render generates an HTML file for the Trip and
// renders to the given `io.Writer`
func (t *Trip) Render(wr io.Writer, i int) error {

	// make sure routes have been fetched
	err := t.checkForRoutes()
	if err != nil {
		return err
	}

	// make sure route index exists
	if i < 0 || i > len(t.Routes) {
		return errors.New("Route index out of bounds")
	}

	// get route data for template
	data, err := t.getTemplateData(t.Routes[i])
	if err != nil {
		return err
	}

	// compile template
	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	// render template
	err = tpl.Execute(wr, data)
	if err != nil {
		return err
	}

	return nil
}

// Convert route data for use in template
func (t *Trip) getTemplateData(route maps.Route) (maps.Route, error) {
	return route, nil
}

// Get summaries for given routes
func generateSummaries(r []maps.Route, m maps.Mode) ([]RouteSummary, error) {

	summaries := make([]RouteSummary, len(r))

	for i, route := range r {
		if m == maps.TravelModeTransit {
			s, err := transitRouteSummary(route)
			if err != nil {
				return nil, err
			}
			summaries[i] = s
		} else {
			s := RouteSummary{route.Legs[0].Duration, route.Summary}
			summaries[i] = s
		}
	}

	return summaries, nil
}

// Make sure routes have been loaded
func (t *Trip) checkForRoutes() error {

	if t.Routes == nil || len(t.Routes) < 1 {
		return errors.New("No Routes available. Be sure to execute the Fetch() method")
	}

	return nil
}

// Make sure Trip requirements have been met
func (t *Trip) checkRequirements() error {

	if t.APIKey == "" {
		return errors.New("APIKey has not been configured")
	}

	if t.Origin == "" {
		return errors.New("Origin has not been configured")
	}

	if t.Destination == "" {
		return errors.New("Destination has not been configured")
	}

	if t.Mode == "" {
		return errors.New("Travel mode has not been specified")
	}

	return nil
}

// Convert a polyline object to a base64 PNG string
func (t *Trip) polyLineToB64(polyline maps.Polyline) (string, error) {

	// create static image url
	imageURL := fmt.Sprintf(imgURLStr, imgW, imgH, polyline.Points, t.APIKey)

	// get image
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

	// encode image
	b64 := base64.StdEncoding.EncodeToString(body)

	return b64, nil
}

// Create a route summary for transit route
func transitRouteSummary(route maps.Route) (RouteSummary, error) {

	duration := route.Legs[0].Duration
	lines := make([]string, 0)

	// iterate through transit steps and add to description
	for _, step := range route.Legs[0].Steps {

		// if this step has transit details
		if step.TransitDetails != nil {

			line := step.TransitDetails.Line
			var name string

			// use the short name if there is one, otherwise
			// use long name
			if line.ShortName != "" {
				name = line.ShortName
			} else {
				name = line.Name
			}
			lines = append(lines, name)
		}
	}

	// join line names into one string
	description := strings.Join(lines, transitLineDelimiter)
	summary := RouteSummary{duration, description}

	return summary, nil
}

// // GoogleMapsData data for google maps directions
// type GoogleMapsData struct {
// 	Origin      string
// 	Destination string
// 	Mode        string
// 	Time        string
// }

// // Settings configuration struct
// type Settings struct {
// 	APIKey string
// }

// var config *Settings

// // Configure configure the module
// func Configure(c *Settings) error {

// 	config = c

// 	if config.APIKey == "" {
// 		return errors.New("API Key has not been configured")
// 	}

// 	return nil
// }

// func prepRouteForTemplate(route maps.Route) (maps.Route, error) {
// 	return route, nil
// }

// // GetRoutes get directions for a set of directions data
// func GetRoutes(data GoogleMapsData) ([]maps.Route, error) {
// 	if config == nil {
// 		return nil, errors.New("Module has not been configured")
// 	}

// 	c, err := maps.NewClient(maps.WithAPIKey(config.APIKey))
// 	if err != nil {
// 		return nil, err
// 	}
// 	r := &maps.DirectionsRequest{
// 		Origin:        data.Origin,
// 		Destination:   data.Destination,
// 		Mode:          maps.Mode(data.Mode),
// 		DepartureTime: data.Time,
// 		Alternatives:  true,
// 	}

// 	routes, _, err := c.Directions(context.Background(), r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return routes, nil
// }

// // GetTransitRouteSummary get a summary of transit options
// func GetTransitRouteSummary(route maps.Route) (RouteSummary, error) {

// 	duration := route.Legs[0].Duration
// 	lines := make([]string, 0)

// 	for _, step := range route.Legs[0].Steps {
// 		if step.TransitDetails != nil {
// 			line := step.TransitDetails.Line
// 			var name string

// 			if line.ShortName != "" {
// 				name = line.ShortName
// 			} else {
// 				name = line.Name
// 			}
// 			lines = append(lines, name)
// 		}
// 	}

// 	summary := RouteSummary{duration, strings.Join(lines, " > ")}

// 	return summary, nil
// }

// // GetRouteSummaries get summaries of each route
// func GetRouteSummaries(mode maps.Mode, routes []maps.Route) ([]RouteSummary, error) {
// 	summaries := make([]RouteSummary, 0)

// 	for _, route := range routes {
// 		if mode == maps.TravelModeTransit {
// 			s, _ := GetTransitRouteSummary(route)
// 			summaries = append(summaries, s)
// 		} else {
// 			s := RouteSummary{route.Legs[0].Duration, route.Summary}
// 			summaries = append(summaries, s)
// 		}
// 	}

// 	return summaries, nil
// }

// // RenderRoute render a given route
// func RenderRoute(wr io.Writer, route maps.Route) error {

// 	// timestamp := strconv.FormatInt(time.Now().Unix(), 10)
// 	// filename := fmt.Sprintf("%s%s.html", os.TempDir(), timestamp)

// 	t, err := template.ParseFiles("./directions.html")
// 	if err != nil {
// 		return err
// 	}

// 	// f, err := os.Create(filename)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// defer f.Close()

// 	prepped, err := prepRouteForTemplate(route)

// 	err = t.Execute(wr, prepped)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // MapImageBase64 take a map polyline and encode as Base64 image
// func MapImageBase64(polyline maps.Polyline) (string, error) {
// 	if config == nil {
// 		return "", errors.New("Module has not been configured")
// 	}

// 	imageURL := fmt.Sprintf(imgURLStr, imgW, imgH, polyline.Points, config.APIKey)

// 	// get image url
// 	resp, err := http.Get(imageURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// read response body into bytes
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	str := base64.StdEncoding.EncodeToString(body)

// 	return str, nil
// }
