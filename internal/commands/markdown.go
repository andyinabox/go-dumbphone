package commands

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/andyinabox/go-dumbphone/internal/utils"
	"github.com/andyinabox/go-dumbphone/pkg/markdown"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// MarkdownSubcommand Subcommand to parse markdown file
var MarkdownSubcommand = cli.Command{
	Name:    "markdown",
	Usage:   "Parse markdown file into html",
	Aliases: []string{"m"},
	Action: func(c *cli.Context) error {

		var (
			promptFilename = func(f *os.File) (string, error) {

				// there are a million ways to do this
				// this is the way I chose
				fn := f.Name()
				base := filepath.Base(fn)
				ext := filepath.Ext(fn)
				name := strings.Replace(base, ext, "", 1)

				prompt := promptui.Prompt{
					Label: "Enter a file name (extension will be added)",
					Validate: func(input string) error {
						// validate file name?
						return nil
					},
					Default: name,
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

		fpath := c.Args().Get(0)

		if fpath == "" {
			return errors.New("Must provide the file path as argument")
		}

		mdfile, err := os.Open(fpath)
		if err != nil {
			return err
		}

		md, err := ioutil.ReadAll(mdfile)
		if err != nil {
			return err
		}

		html, err := markdown.Parse(md)
		if err != nil {
			return err
		}

		fn, err := promptFilename(mdfile)

		file, err := utils.CreateTempFile(fn)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(html)
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
			return errors.New("USB transfer not yet implemented")
			// err := utils.USBSend(file)
			// if err != nil {
			// 	return err
			// }
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
