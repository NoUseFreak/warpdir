package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all warp points",
		Run: func(cmd *cobra.Command, args []string) {
			index := viper.GetStringMap("index")

			for name, path := range index {
				logrus.Infof("%-20s (%s)\n", name, path)
			}
		},
	}
}
