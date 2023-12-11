package cmd

import (
	"os"
)

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.InitDefaultCompletionCmd()
	rootCmd.SetOut(os.Stderr)
}
