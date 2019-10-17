package main

import (
	"github.com/andyinabox/go-dumbphone/cmd/dumbp/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		// directions
		cmd.DirectionsSubcommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
