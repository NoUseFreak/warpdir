package cmd

import (
	"strings"
	"testing"
    "os"
)

func Test_Warp(t *testing.T) {
    cfg := getTempCfgPath(t, "list.yaml") 
    defer os.Remove(cfg)
	sout, serr, err := execute(t, cfg, "warp", "opt")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !strings.Contains(sout, "cd /opt") {
		t.Errorf("Expected stdout to contain 'cd /opt'")
	}
	if !strings.Contains(serr, "opt") {
		t.Errorf("Expected stderr to contain 'opt'")
	}
}
