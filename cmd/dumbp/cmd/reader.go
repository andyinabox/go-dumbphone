package cmd

import (
	"net/url"

	"github.com/andyinabox/go-dumbphone/bin/data"
	"github.com/andyinabox/go-dumbphone/pkg/reader"
	"github.com/andyinabox/go-dumbphone/pkg/transfer"
	"github.com/andyinabox/go-dumbphone/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// ReaderSubcommand Subcommand to get reading material
var ReaderSubcommand = cli.Command{
	Name:    "reader",
	Usage:   "Convert web page to readable text",
	Aliases: []string{"r"},
	Action: func(c *cli.Context) error {

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

		tpl, err := data.Asset("bin/data/reader.html")
		if err != nil {
			return err
		}
		err = article.Render(file, tpl)
		if err != nil {
			return err
		}

		index, _, err := promptTransfer()
		if err != nil {
			return err
		}

		switch index {
		case 0:
			return transfer.Send(file, transfer.BLUETOOTH_SEND, "")
		case 1:
			return transfer.Send(file, transfer.USB_SEND, viper.GetString("modules.directions.dir"))
		case 2:
			return transfer.Send(file, transfer.BROWSER_SEND, "")
		}

		return nil
	},
}
