package usb

type Config struct {
	Root          string `desc:"Path to your phone when plugged into USB"`
	PodcastsDir   string `default:"/Podcasts" desc:"Podcasts folder on your phone"`
	NotesDir      string `default:"/Notes"desc:"Notes folder on your phone"`
	ReaderDir     string `default:"/Reading"desc:"Reading folder on your phone"`
	DirectionsDir string `default:"/Directions"desc:"Directions folder on your phone"`
}

var ConfigDefaults = &Config{
	"",
	"/Podcasts",
	"/Notes",
	"/Reading",
	"/Directions"
}