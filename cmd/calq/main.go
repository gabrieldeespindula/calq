package main

import (
	"flag"
	"fmt"
	"os"
	"calq/internal/repl"
)

var Version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Print version information")
  showVersionShort := flag.Bool("v", false, "Print version information (shorthand)")

  flag.Parse()

	if *showVersion || *showVersionShort {
		fmt.Println(Version)
		os.Exit(0)
	}

	repl.Run(os.Stdin, os.Stdout)
}
