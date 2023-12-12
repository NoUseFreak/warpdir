package main

import (
	"github.com/nousefreak/warpdir/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date

	cmd.Execute()
}
