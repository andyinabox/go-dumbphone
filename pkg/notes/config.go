package notes

type Config struct {
	NotesDir string `desc:"Notes folder on your computer"`
}

var ConfigDefaults = &Config{
	"",
}
