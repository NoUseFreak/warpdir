package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add name [path]",
		Short: "Add a warp point",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if isReservedKeyword(cmd, args[0]) {
				return fmt.Errorf("'%s' is a reserved keyword", args[0])
			}

			curdir, err := os.Getwd()
			if len(args) >= 2 {
				curdir = args[1]
			} else if err != nil {
				return fmt.Errorf("Failed to find current directory")
			}

			index := viper.GetStringMap("index")
			index[args[0]] = curdir
			viper.Set("index", index)

			if err := viper.WriteConfig(); err != nil {
				return err
			}
			logrus.Infof("Adding '%s' with path '%s'", args[0], curdir)

			return nil
		},
	}
}

func isReservedKeyword(cmd *cobra.Command, name string) bool {
	for _, cmd := range GetCLI().Commands() {
		if cmd.Name() == name {
			return true
		}
	}
	return false
}
