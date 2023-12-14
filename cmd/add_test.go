package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_Add(t *testing.T) {
    cfg := getTempCfgPath(t, "add.yaml") 
    defer os.Remove(cfg)
	sout, serr, err := execute(t, cfg, "add", "newitem", "/new/path")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(serr, "newitem") {
		t.Errorf("Expected stderr to contain 'home'")
	}
	if sout != "" {
		t.Errorf("Expected stdout to be empty")
	}

	sout, serr, err = execute(t, cfg, "list")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !strings.Contains(serr, "newitem") {
		t.Errorf("Expected stderr to contain 'newitem'")
	}

	if sout != "" {
		t.Errorf("Expected stdout to be empty")
	}
}
