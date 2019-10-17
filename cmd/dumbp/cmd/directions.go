package cmd

import (
	"fmt"
	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"os"
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

func promptRoutes(t *directions.Trip) (int, string, error) {

	var options = make([]string, len(t.Summaries))

	for i, s := range t.Summaries {
		options[i] = s.ToString()
	}

	prompt := promptui.Select{
		Label: "Route Options",
		Items: options,
	}

	return prompt.Run()
}

// DirectionsSubcommand Subcommand to get directions
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

		trip := directions.Trip{
			APIKey:      os.Getenv("GOOGLE_API_KEY"),
			Origin:      origin,
			Destination: destination,
			Mode:        mode,
			Time:        departureTime,
		}

		err = trip.Fetch()
		if err != nil {
			return err
		}

		index, summary, err := promptRoutes(&trip)
		if err != nil {
			return err
		}

		fmt.Printf("%v: %v\n", index, summary)

		return nil
	},
}
