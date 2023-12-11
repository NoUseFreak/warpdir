package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete name",
	Short: "Delete a warp point",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := viper.GetStringMap("index")
		delete(index, args[0])
		viper.Set("index", index)
		if err := viper.WriteConfig(); err != nil {
			logrus.Error(err)
		}

		logrus.Info("Warp point deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
