package directions

import (
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"testing"
)

func TestConfig(t *testing.T) {
	if DirectionsConfig == nil {
		t.Errorf("No DirectionsConfig found")
	}
}
