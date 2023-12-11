package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add name [path]",
	Short: "Add a warp point",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == args[0] {
				logrus.Errorf("'%s' is a reserved keyword", args[0])
				os.Exit(1)
			}
		}

		curdir, err := os.Getwd()
		if len(args) >= 2 {
			curdir = args[1]
		} else if err != nil {
			logrus.Fatal("Failed to find current directory")
		}

		index := viper.GetStringMap("index")
		index[args[0]] = curdir
		viper.Set("index", index)

		if err := viper.WriteConfig(); err != nil {
			logrus.Fatal(err)
		}
		logrus.Infof("Adding '%s' with path '%s'", args[0], curdir)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
