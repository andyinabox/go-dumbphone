package directions

import (
	"googlemaps.github.io/maps"
	"testing"
)

var testConfig Settings

func init() {
	testConfig = Settings{
		"AIzaSyCbLP2s621kGDdESEGvVW0bhO1qkSu7WjQ",
		"300 Nicollet Mall, Minneapolis, MN",
	}
}

func defaultMapData() GoogleMapsData {
	data := GoogleMapsData{
		Origin:      "300 Nicollet Mall, Minneapolis, MN",
		Destination: "90 W 4th St, St Paul, MN",
		Mode:        "driving",
		Time:        "now",
	}
	return data
}

// func openBrowser(url string) {
// 	exec.Command("open", url).Start()
// }

func TestConfig(t *testing.T) {

	_, err := GetRoutes(defaultMapData())

	if err == nil {
		t.Errorf("Expected config error, got none")
	} else {
		t.Logf("Expected Error: %s", err)
	}

	c := Settings{}

	err = Configure(&c)

	if err == nil {
		t.Errorf("Should have recieved no API Key error")
	} else {
		t.Logf("Expected Error: %s", err)
	}

	err = Configure(&testConfig)

	if err != nil {
		t.Errorf("Configuration error: %s", err)
	}

}

func TestRoutes(t *testing.T) {

	routes, err := GetRoutes(defaultMapData())

	if err != nil {
		t.Errorf("Routes error: %s", err)
	}

	if len(routes) < 1 {
		t.Errorf("Expected 1 or more routes, found %d", len(routes))
	} else {
		t.Logf("Found %d routes", len(routes))
	}

}

func TestRouteSummaries(t *testing.T) {

	data := defaultMapData()
	routes, err := GetRoutes(data)

	if err != nil {
		t.Errorf("Routes error: %s", err)
	}

	summaries := GetRouteSummaries(maps.Mode(data.Mode), routes)

	if len(summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(summaries))
		for _, s := range summaries {
			t.Logf(s.toString())
		}
	}

}

func TestTransitRouteSummaries(t *testing.T) {

	data := defaultMapData()
	data.Mode = "transit"

	routes, err := GetRoutes(data)

	if err != nil {
		t.Errorf("Routes error: %s", err)
	}

	summaries := GetRouteSummaries(maps.Mode(data.Mode), routes)

	if len(summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(summaries))
		for _, s := range summaries {
			t.Logf(s.toString())
		}
	}
}

func TestBase64(t *testing.T) {

	routes, err := GetRoutes(defaultMapData())

	if err != nil {
		t.Errorf("Routes error: %s", err)
	}

	b64, err := MapImageBase64(routes[0].OverviewPolyline)

	if err != nil {
		t.Errorf("Error encoding image: %s", err)
	} else {
		t.Log(b64)
	}

}

func TestRender(t *testing.T) {

}
