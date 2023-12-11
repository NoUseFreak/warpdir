package cmd

import (
	"fmt"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var warpCmd = &cobra.Command{
	Use:   "warp",
	Short: "Warp to a directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) > 0 {
			name = args[0]
		}

		index := viper.GetStringMapString("index")

		if name == "-" {
			history := viper.GetStringSlice("history")
			if len(history) > 1 {
				name = history[len(history)-2]
				return warp(
					name,
					index[name],
				)
			}
		}

		if target, ok := index[name]; ok {
			return warp(name, target)
		}

		data := newFzfItemSlice(index)
		idx, err := fuzzyfinder.Find(
			data,
			func(i int) string {
				return fmt.Sprintf("%-20s %s", data[i].name, data[i].path)
			},
			fuzzyfinder.WithQuery(name),
		)
		if err == nil {
			return warp(data[idx].name, data[idx].path)
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(warpCmd)
}

func warp(name string, path string) error {
	recordHistory(name)

	logrus.Infof("Warping to %s", path)
	fmt.Printf("cd %s", path)
	return nil
}

func recordHistory(name string) {
	history := viper.GetStringSlice("history")
	if len(history) > 0 && history[len(history)-1] != name {
		history = append(history, name)
		if len(history) > 2 {
			history = history[len(history)-2:]
		}
		viper.Set("history", history)
		if err := viper.WriteConfig(); err != nil {
			logrus.Warn("Failed to record history")
		}
	}
}

type fzfItem struct {
	name string
	path string
}

func newFzfItemSlice(index map[string]string) []fzfItem {
	items := []fzfItem{}
	for name, path := range index {
		items = append(items, fzfItem{name, path})
	}
	return items
}
