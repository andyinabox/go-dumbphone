package commands

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// DirectionsSubcommand Subcommand to get directions
var ConfigSubcommand = cli.Command{
	Name:    "config",
	Usage:   "View and edit config",
	Aliases: []string{"c"},
	Action: func(c *cli.Context) error {

		s := viper.AllSettings()
		bs, err := yaml.Marshal(s)
		if err != nil {
			return err
		}

		fmt.Println("")
		fmt.Println(viper.ConfigFileUsed())
		fmt.Println("")
		fmt.Println(string(bs))

		return nil
	},
}
