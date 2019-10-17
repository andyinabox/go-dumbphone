package directions

import (
	"testing"
	"time"
)

func TestConfig(t *testing.T) {

	data := GoogleMapsData{
		"2105 22nd Ave S, Minneapolis, MN",
		"1074 Vail Dr, Mendota Heights, MN",
		"driving",
		string(time.Now().Unix()),
	}

	_, err := GetRoutes(data)

	if err == nil {
		t.Errorf("Expected config error, got none")
	}

	config := Settings{}

	err = Configure(&config)

	if err == nil {
		t.Errorf("Should have recieved no API Key error")
	}

	config = Settings{
		"AIzaSyCbLP2s621kGDdESEGvVW0bhO1qkSu7WjQ",
		"2105 22nd ave s, minneapolis, mn, 55404",
	}

	err = Configure(&config)

	if err != nil {
		t.Errorf("Configuration error: %s", err)
	}

}
