package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	// "github.com/andyinabox/go-dumbphone/pkg/directions"
)

func validateOrigin(input string) error {
	return nil
}

// DirectionsSubcommand Subcommand for directions
var DirectionsSubcommand = cli.Command{
	Name:  "directions",
	Usage: "Get directions from Google Maps",
	Action: func(c *cli.Context) error {
		prompt := promptui.Prompt{
			Label:    "Origin",
			Validate: validateOrigin,
			Default:  "300 Nicollet Mall, Minneapolis, MN",
		}

		result, err := prompt.Run()

		if err != nil {
			return errors.New("Prompt failed")
		}

		fmt.Printf("The origin address is: %s", result)

		return nil
	},
}
