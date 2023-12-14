package cmd

import (
	"strings"
	"testing"
    "os"
)


func Test_List(t *testing.T) {
    cfg := getTempCfgPath(t, "list.yaml") 
    defer os.Remove(cfg)
	sout, serr, err := execute(t, cfg, "list")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !strings.Contains(serr, "home") {
		t.Errorf("Expected stderr to contain 'home'")
	}
	if !strings.Contains(serr, "opt") {
		t.Errorf("Expected stderr to contain 'opt'")
	}

	if sout != "" {
		t.Errorf("Expected stdout to be empty")
	}
}
