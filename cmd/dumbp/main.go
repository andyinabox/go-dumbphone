package main

import (
	"github.com/andyinabox/go-dumbphone/internal/commands"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		// directions
		commands.DirectionsSubcommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
