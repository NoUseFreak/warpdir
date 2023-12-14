package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_Delete(t *testing.T) {
	deleteCfg := getTempCfgPath(t, "delete.yaml")
	defer os.Remove(deleteCfg)

	sout, serr, err := execute(t, deleteCfg, "add", "opt2", "/opt2")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(serr, "opt2") {
		t.Errorf("Expected stderr to contain 'opt2'")
	}
	if sout != "" {
		t.Errorf("Expected stdout to be empty")
	}

    sout, serr, err = execute(t, deleteCfg, "delete", "opt2")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if strings.Contains(serr, "opt2") {
		t.Errorf("Expected stderr to contain 'opt2'")
	}
	if sout != "" {
		t.Errorf("Expected stdout to be empty")
	}
}
