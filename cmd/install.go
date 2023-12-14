package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func getInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install the wd command",
		Run: func(cmd *cobra.Command, args []string) {
			exec, _ := os.Executable()
			script := fmt.Sprintf("%s () { eval $(%s $@) }", cmdName, exec)
			installScript(script, os.Getenv("SHELL"))
		},
	}
}

func installScript(script string, shell string) {
	var shellFile string
	switch shell {
	case "/bin/bash":
		shellFile = os.Getenv("HOME") + "/.bashrc"
	case "/bin/zsh":
		shellFile = os.Getenv("HOME") + "/.zshrc"
	default:
		logrus.Errorf("Shell '%s' not supported", shell)
		logrus.Info("Add the following to your shell config file:")
		logrus.Info(script)
	}

	found, err := upsertScript(shellFile, script)
	if err != nil {
		logrus.Fatal(err)
	}

	if found {
		logrus.Infof("Updated script in %s", shellFile)
	} else {
		logrus.Infof("Added script to %s", shellFile)
	}
}

func upsertScript(shellFile, scriptLine string) (found bool, err error) {
	file, err := os.ReadFile(shellFile)
	if err != nil {
		logrus.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if strings.Contains(line, cmdName) {
			lines[i] = scriptLine
			found = true
		}
	}
	if !found {
		lines = append(lines, scriptLine)
	}

	return found, os.WriteFile(shellFile, []byte(strings.Join(lines, "\n")), 0644)
}
