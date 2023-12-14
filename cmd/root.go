package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cmdName string = "wd"
)

func GetCLI() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [name]", cmdName),
		Short: "Warp Directory",
		Long:  "Warp Directory allows you to quickly warp between directories.",
		Args:  cobra.MaximumNArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			index := viper.GetStringMap("index")
			var suggestions []string
			for name, path := range index {
				if name == toComplete {
					return []string{}, cobra.ShellCompDirectiveNoFileComp
				}
				suggestions = append(suggestions, fmt.Sprintf("%s\tpath: %s", name, path))
			}
			return suggestions, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return getWarpCmd().RunE(cmd, args)
		},
	}

	rootCmd.DisableSuggestions = true
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "config file (default is $HOME/.warpdir.yaml)")
	if err := rootCmd.PersistentFlags().MarkHidden("config"); err != nil {
		logrus.Error(err)
	}


	rootCmd.AddCommand(getAddCmd())
	rootCmd.AddCommand(getDeleteCmd())
	rootCmd.AddCommand(getInstallCmd())
	rootCmd.AddCommand(getListCmd())
	rootCmd.AddCommand(getManCmd())
	rootCmd.AddCommand(getVersionCmd())
	rootCmd.AddCommand(getWarpCmd())

	rootCmd.SetOut(os.Stdout)
	rootCmd.InitDefaultCompletionCmd()
	rootCmd.SetOut(os.Stderr)

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := GetCLI().Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println(" ")
}
