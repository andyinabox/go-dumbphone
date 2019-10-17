package cmd

import (
	// "errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	// "github.com/andyinabox/go-dumbphone/pkg/directions"
)

func promptOrigin() (string, error) {
	prompt := promptui.Prompt{
		Label: "Origin",
		Validate: func(input string) error {
			return nil
		},
		Default: "300 Nicollet Mall, Minneapolis, MN",
	}

	return prompt.Run()
}

// DirectionsSubcommand Subcommand for directions
var DirectionsSubcommand = cli.Command{
	Name:  "directions",
	Usage: "Get directions from Google Maps",
	Action: func(c *cli.Context) error {

		origin, err := promptOrigin()
		if err != nil {
			return err
		}

		fmt.Println(origin)

		return nil
	},
}
