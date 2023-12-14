package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getTempCfgPath(t *testing.T, testCfg string) string {
    // t.Helper()

    wd, err := os.Getwd()
    if err != nil {
        t.Errorf("Error: %s", err)
    }
    srcPath := filepath.Join(filepath.Dir(wd), "testdata", testCfg)

    file, err := os.CreateTemp("", "*-"+testCfg)
    if err != nil {
        t.Errorf("Error: %s", err)
        return ""
    }
    src, err := os.ReadFile(srcPath)
    if err != nil {
        t.Errorf("Error: %s", err)
        return ""
    }
    if err = os.WriteFile(file.Name(), src, 0644); err != nil {
        t.Errorf("Error: %s", err)
        return ""
    }

    return file.Name()
}

func execute(t *testing.T, testCfg string, args ...string) (string, string, error) {
	// t.Helper()

	buferr := bytes.NewBufferString("")
	bufout := bytes.NewBufferString("")
    logrus.SetOutput(buferr)

	cmd := GetCLI()
	cmd.SetOut(bufout)
	cmd.SetArgs(append([]string{"--config", testCfg}, args...))
		viper.SetConfigFile(testCfg)
        viper.ReadInConfig()

	err := cmd.Execute()

	return strings.TrimSpace(bufout.String()), strings.TrimSpace(buferr.String()), err
}
