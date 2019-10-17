package directions

import "github.com/sakirsensoy/genv"

type directionsConfig struct {
	APIKey      string
	HomeAddress string
}

// DirectionsConfig config object for directions
var DirectionsConfig = &directionsConfig{
	APIKey:      genv.Key("DUMBP_GMAPS_API_KEY").String(),
	HomeAddress: genv.Key("DUMBP_HOME_ADDRESS").String(),
}
