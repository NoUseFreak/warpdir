package cmd

import (
	"os"
	"path/filepath"

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
	viper.SetConfigName("warpdir")
	viper.AddConfigPath(home + "/.config/warpdir")
	viper.AddConfigPath(home)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn(err)
        cfgFile = home + "/.config/warpdir/warpdir.yaml"
		if err := os.MkdirAll(filepath.Dir(cfgFile), 0755); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Warn("Created new config")
	}
}
