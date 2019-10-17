package main

import (
	"github.com/andyinabox/go-dumbphone/cmd/dumbp/cmd"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app.Commands = []cli.Command{
		// directions
		cmd.DirectionsSubcommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
