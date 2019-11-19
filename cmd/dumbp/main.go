package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andyinabox/go-dumbphone/internal/commands"
	"github.com/andyinabox/go-dumbphone/pkg/config"
	"github.com/urfave/cli"
)

const version = "0.1.2"
const configFile = "dumbp/config.yaml"

var configPath string

func main() {

	userDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := fmt.Sprintf("%s/%s", configFile)

	config, exists, err := config.Load()
	// err := godotenv.Load()
	// if err != nil {
	// panic("Error loading .env file")
	// }

	app := cli.NewApp()
	app.Name = "dumbphone"
	app.Usage = "A set of tools to make your dumbphone a little smarter"
	app.Version = version
	app.Commands = []cli.Command{
		// directions
		commands.DirectionsSubcommand,
		commands.ReaderSubcommand,
		commands.MarkdownSubcommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
