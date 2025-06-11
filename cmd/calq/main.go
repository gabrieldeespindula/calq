package main

import (
	"os"

	"calq/internal/cli"
	"calq/internal/repl"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	info := cli.VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	}

	code := cli.RunParse(info, func() {
        repl.Run(os.Stdin, os.Stdout)
    })
	
    os.Exit(code)
}
