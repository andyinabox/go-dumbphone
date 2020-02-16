package directions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/andyinabox/go-dumbphone/data"
	"github.com/joho/godotenv"
)

const (
	configDir  = ".dumbp/"
	configName = "config"
)

var apiKey string

func init() {
	err := godotenv.Load("../../test/.env")
	if err != nil {
		panic("Unable to load testing .env file. Make sure you have test/.env set")
	}
}

func defaultTrip() Trip {
	trip := Trip{
		APIKey:             os.Getenv("GOOGLE_API_KEY"),
		Origin:             "300 Nicollet Mall, Minneapolis, MN",
		Destination:        "90 W 4th St, St Paul, MN",
		Mode:               "driving",
		Time:               "now",
		DetailedDirections: true,
	}
	return trip
}

func previewB64(t *testing.T, b64 string) error {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	filename := fmt.Sprintf("%s/%s.html", os.TempDir(), timestamp)
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

	err = openBrowser("file://" + filename)

	return nil
}

func openBrowser(url string) error {
	cmd := exec.Command("open", url)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func TestConfig(t *testing.T) {

	trip := defaultTrip()
	trip.APIKey = ""

	err := trip.Fetch()
	if err == nil {
		t.Errorf("Expected requirements error, got none")
	} else {
		t.Log(err)
	}

	trip.APIKey = apiKey
	trip.Origin = ""
	err = trip.Fetch()
	if err == nil {
		t.Errorf("Expected requirements error, got none")
	} else {
		t.Log(err)
	}

	trip.Origin = "300 Nicollet Mall, Minneapolis, MN"
	trip.Destination = ""
	err = trip.Fetch()
	if err == nil {
		t.Errorf("Expected requirements error, got none")
	} else {
		t.Log(err)
	}

	trip.Destination = "90 W 4th St, St Paul, MN"
	trip.Mode = ""
	err = trip.Fetch()
	if err == nil {
		t.Errorf("Expected requirements error, got none")
	} else {
		t.Log(err)
	}

}

func TestFetch(t *testing.T) {

	trip := defaultTrip()
	err := trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	if len(trip.Routes) < 1 {
		t.Errorf("Expected 1 or more routes, found %d", len(trip.Routes))
	} else {
		t.Logf("Found %d routes", len(trip.Routes))
	}

}

func TestRouteSummaries(t *testing.T) {

	trip := defaultTrip()
	err := trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	if len(trip.Summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(trip.Summaries))
		for _, s := range trip.Summaries {
			t.Logf(s.ToString())
		}
	}

}

func TestTransitRouteSummaries(t *testing.T) {

	trip := defaultTrip()
	trip.Mode = "transit"
	err := trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	if len(trip.Summaries) < 1 {
		t.Errorf("No summaries found")
	} else {
		t.Logf("Found %d summaries", len(trip.Summaries))
		for _, s := range trip.Summaries {
			t.Logf(s.ToString())
		}
	}
}

func TestBase64(t *testing.T) {

	trip := defaultTrip()
	err := trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	b64, err := trip.PolyLineToB64(trip.Routes[0].OverviewPolyline)

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

func TestRenderDriving(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	filename := fmt.Sprintf("%s/%s.html", os.TempDir(), timestamp)

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		t.Errorf("File Error: %v", err)
	}

	trip := defaultTrip()
	err = trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	tpl, err := data.Asset("bin/data/directions.html")
	if err != nil {
		t.Fatal(err)
	}
	err = trip.Render(f, 0, tpl)
	if err != nil {
		t.Errorf("Rendering error: %v", err)
	}

	if testing.Verbose() {
		err = openBrowser("file://" + filename)
		if err != nil {
			t.Logf("Error opening preview: %v\n", err)
		}
	}
}

func TestRenderTransit(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	filename := fmt.Sprintf("%s/%s.html", os.TempDir(), timestamp)

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		t.Errorf("File Error: %v", err)
	}

	trip := defaultTrip()
	trip.Mode = "transit"
	err = trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	tpl, err := data.Asset("bin/data/directions.html")
	if err != nil {
		t.Fatal(err)
	}
	err = trip.Render(f, 0, tpl)
	if err != nil {
		t.Errorf("Rendering error: %v", err)
	}

	if testing.Verbose() {
		err = openBrowser("file://" + filename)
		if err != nil {
			t.Logf("Error opening preview: %v\n", err)
		}
	}
}

func TestRenderBicycling(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	filename := fmt.Sprintf("%s/%s.html", os.TempDir(), timestamp)

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		t.Errorf("File Error: %v", err)
	}

	trip := defaultTrip()
	trip.Mode = "bicycling"
	trip.DetailedDirections = false
	err = trip.Fetch()

	if err != nil {
		t.Errorf("Fetch Error: %v", err)
	}

	tpl, err := data.Asset("bin/data/directions.html")
	if err != nil {
		t.Fatal(err)
	}
	err = trip.Render(f, 0, tpl)
	if err != nil {
		t.Errorf("Rendering error: %v", err)
	}

	if testing.Verbose() {
		err = openBrowser("file://" + filename)
		if err != nil {
			t.Logf("Error opening preview: %v\n", err)
		}
	}
}
