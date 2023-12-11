package cmd

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.warpdir.yaml)")
	if err := rootCmd.PersistentFlags().MarkHidden("config"); err != nil {
		logrus.Error(err)
	}
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home)
		viper.SetConfigName(".warpdir")
	}

	if err := viper.ReadInConfig(); err != nil {
		if err := viper.SafeWriteConfigAs(home + "/.warpdir.yaml"); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Warn("Created new config")
	}
}
