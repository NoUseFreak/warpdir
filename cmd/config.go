package cmd

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)



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
        logrus.Warn(err)
		if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Warn("Created new config")
	}
}
