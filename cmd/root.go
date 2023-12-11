package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdName string = "wd"

var rootCmd = &cobra.Command{
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
		return warpCmd.RunE(cmd, args)
	},
}

func Execute() {
	rootCmd.DisableSuggestions = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println(" ")
}
