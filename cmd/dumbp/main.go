package main

import (
	"log"
	"os"

	"github.com/andyinabox/go-dumbphone/internal/commands"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

const version = "0.1.2"

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
		commands.ReaderSubcommand,
		commands.MarkdownSubcommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
