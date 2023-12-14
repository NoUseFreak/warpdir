package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getWarpCmd() *cobra.Command {
	return &cobra.Command{
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
						cmd.OutOrStdout(),
						name,
						index[name],
					)
				}
			}

			if target, ok := index[name]; ok {
				return warp(cmd.OutOrStdout(), name, target)
			}

            if !isatty.IsTerminal(os.Stdin.Fd()) {
                return cmd.Root().Usage()
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
				return warp(cmd.OutOrStdout(), data[idx].name, data[idx].path)
			}

			return err
		},
	}
}

func warp(o io.Writer, name string, path string) error {
	recordHistory(name)

	logrus.Infof("Warping to %s", path)
	fmt.Fprintf(o, "cd %s", path)
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
	sort.Slice(items, func(i, j int) bool {
		return items[i].name > items[j].name
	})
	return items
}
