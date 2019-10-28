package commands

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/andyinabox/go-dumbphone/internal/utils"
	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// DirectionsSubcommand Subcommand to get directions
var DirectionsSubcommand = cli.Command{
	Name:    "directions",
	Usage:   "Get directions from Google Maps",
	Aliases: []string{"d"},
	Action: func(c *cli.Context) error {

		const (
			templateFile = "./pkg/directions/directions.html"
		)

		var (
			promptOrigin = func() (string, error) {
				prompt := promptui.Prompt{
					Label: "Origin",
					Validate: func(input string) error {
						return nil
					},
					Default: os.Getenv("DUMBP_HOME_ADDRESS"),
				}

				return prompt.Run()
			}

			promptDestination = func() (string, error) {
				prompt := promptui.Prompt{
					Label: "Destination",
					Validate: func(input string) error {
						return nil
					},
					Default: "",
				}

				return prompt.Run()
			}

			promptDepartureTime = func() (string, error) {
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

			promptMode = func() (int, string, error) {
				prompt := promptui.Select{
					Label: "Mode",
					Items: []string{"driving", "bicycling", "transit", "walking"},
				}

				return prompt.Run()
			}

			promptDetailedDirections = func() (int, string, error) {
				prompt := promptui.Select{
					Label: "Include detailed directions?",
					Items: []string{"Yes", "No"},
				}

				return prompt.Run()
			}

			promptRoutes = func(t *directions.Trip) (int, string, error) {

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

			promptFileName = func() (string, error) {
				timestamp := strconv.FormatInt(time.Now().Unix(), 10)
				prompt := promptui.Prompt{
					Label: "Filename",
					Validate: func(input string) error {
						// validate filename?
						return nil
					},
					Default: timestamp,
				}

				return prompt.Run()
			}

			promptTransfer = func() (int, string, error) {
				prompt := promptui.Select{
					Label: "How would you like to deliver the directions?",
					Items: []string{
						"Bluetooth",
						"USB",
						"Open in Browser",
					},
				}

				return prompt.Run()
			}
		)

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

		var detailedDirections bool
		index, _, err := promptDetailedDirections()
		if err != nil {
			return err
		}

		switch index {
		case 0:
			detailedDirections = true
			break
		case 1:
			detailedDirections = false
			break
		}

		trip := directions.Trip{
			APIKey:             os.Getenv("GOOGLE_API_KEY"),
			Origin:             origin,
			Destination:        destination,
			Mode:               mode,
			Time:               departureTime,
			DetailedDirections: detailedDirections,
		}

		err = trip.Fetch()
		if err != nil {
			return err
		}

		index, _, err = promptRoutes(&trip)
		if err != nil {
			return err
		}

		filename, err := promptFileName()
		if err != nil {
			return err
		}

		filename = fmt.Sprintf("%s-%s.html", filename, trip.Mode)
		file, err := utils.CreateTempFile(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = trip.Render(file, index, templateFile)
		if err != nil {
			return err
		}

		index, _, err = promptTransfer()
		if err != nil {
			return err
		}

		switch index {
		case 0:
			err := utils.BluetoothSend(file)
			if err != nil {
				return err
			}
			break
		case 1:
			err := utils.USBSend(file)
			if err != nil {
				return err
			}
			break
		case 2:
			err := utils.BrowserSend(file)
			if err != nil {
				return err
			}
			break
		}

		return nil
	},
}
