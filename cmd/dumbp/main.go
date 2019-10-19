package main

import (
	"github.com/andyinabox/go-dumbphone/internal/commands"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

const version = "0.1.0"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := cli.NewApp()
	app.Name = "dumbphone"
	app.Usage = "A set of tools to make your dumbphone a little smarter"
	app.Version = version
	app.Commands = []cli.Command{
		// directions
		commands.DirectionsSubcommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
