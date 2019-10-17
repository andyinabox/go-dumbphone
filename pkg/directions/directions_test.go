package directions

import (
	"fmt"
	"googlemaps.github.io/maps"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

var testConfig Settings

func init() {
	testConfig = Settings{
		"AIzaSyCbLP2s621kGDdESEGvVW0bhO1qkSu7WjQ",
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

func previewB64(t *testing.T, b64 string) error {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	filename := fmt.Sprintf("%s%s.html", os.TempDir(), timestamp)
	html := fmt.Sprintf("<img src=\"data:image/png;base64,%s\">", b64)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(html)
	if err != nil {
		return err
	}

	t.Logf("Created temp file: %s\n", filename)

	cmd := exec.Command(
		"open",
		fmt.Sprintf("file://%s", filename),
	)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
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

	summaries, err := GetRouteSummaries(maps.Mode(data.Mode), routes)

	if len(summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(summaries))
		for _, s := range summaries {
			t.Logf(s.ToString())
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

	summaries, err := GetRouteSummaries(maps.Mode(data.Mode), routes)

	if len(summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(summaries))
		for _, s := range summaries {
			t.Logf(s.ToString())
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
		if testing.Verbose() {
			err = previewB64(t, b64)
			if err != nil {
				t.Logf("Error opening preview: %v\n", err)
			}
		}
	}

}

func TestRender(t *testing.T) {

}
