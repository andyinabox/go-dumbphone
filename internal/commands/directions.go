package commands

import (
	// "errors"
	"fmt"
	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"googlemaps.github.io/maps"
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

func promptDestination() (string, error) {
	prompt := promptui.Prompt{
		Label: "Destination",
		Validate: func(input string) error {
			return nil
		},
		Default: "90 W 4th St, St Paul, MN",
	}

	return prompt.Run()
}

func promptDepartureTime() (string, error) {
	prompt := promptui.Prompt{
		Label: "Departure Time",
		Validate: func(input string) error {
			// validate time expression?
			return nil
		},
		Default: "now",
	}

	return prompt.Run()
}

func promptMode() (int, string, error) {
	prompt := promptui.Select{
		Label: "Mode",
		Items: []string{"driving", "bicycling", "transit", "walking"},
	}

	return prompt.Run()
}

func promptRoutes(mode string, routes []maps.Route) (int, string, error) {

	var options = make([]string, 0)

	summaries, err := directions.GetRouteSummaries(maps.Mode(mode), routes)
	if err != nil {
		return -1, "", err
	}

	for _, s := range summaries {
		options = append(options, s.ToString())
	}

	prompt := promptui.Select{
		Label: "Mode",
		Items: options,
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

		destination, err := promptDestination()
		if err != nil {
			return err
		}

		_, mode, err := promptMode()
		if err != nil {
			return err
		}

		departureTime, err := promptDepartureTime()
		if err != nil {
			return err
		}

		data := directions.GoogleMapsData{
			Origin:      origin,
			Destination: destination,
			Mode:        mode,
			Time:        departureTime,
		}

		directions.Configure(&directions.Settings{
			APIKey: "AIzaSyCbLP2s621kGDdESEGvVW0bhO1qkSu7WjQ",
		})

		routes, err := directions.GetRoutes(data)
		if err != nil {
			return err
		}

		index, summary, err := promptRoutes(mode, routes)
		if err != nil {
			return err
		}

		fmt.Printf("%v: %v", index, summary)

		return nil
	},
}
