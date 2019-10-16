package config

import "github.com/sakirsensoy/genv"

type directionsConfig struct {
	GmapsKey    string
	HomeAddress string
}

var DirectionsConfig = &directionsConfig{
	GmapsKey:    genv.Key("DUMBP_GMAPS_API_KEY").String(),
	HomeAddress: genv.Key("DUMBP_HOME_ADDRESS").String(),
}
