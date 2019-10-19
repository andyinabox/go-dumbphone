package commands

import (
	"net/url"

	"github.com/andyinabox/go-dumbphone/internal/utils"
	"github.com/andyinabox/go-dumbphone/pkg/reader"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// ReaderSubcommand Subcommand to get reading material
var ReaderSubcommand = cli.Command{
	Name:  "reader",
	Usage: "Convert web page to readable text",
	Action: func(c *cli.Context) error {

		const (
			templateFile = "./pkg/reader/reader.html"
		)

		var (
			promptURL = func() (string, error) {
				prompt := promptui.Prompt{
					Label: "URL",
					Validate: func(input string) error {
						_, err := url.Parse(input)
						if err != nil {
							return err
						}

						return nil
					},
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

		url, err := promptURL()
		if err != nil {
			return err
		}

		article, err := reader.New(url)
		if err != nil {
			return err
		}

		file, err := utils.CreateTempFile("")
		if err != nil {
			return err
		}
		defer file.Close()

		err = article.Render(file, templateFile)
		if err != nil {
			return err
		}

		index, _, err := promptTransfer()
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
