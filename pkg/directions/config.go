package directions

type Config struct {
	GoogleAPIKey string `desc:"Google Maps API Key"`
	HomeAddress  string `desc:"Default starting address for directions"`
}

var ConfigDefaults = &Config{
	"",
	"",
}
