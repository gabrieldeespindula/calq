package main

import (
	"os"
	"calq/internal/repl"
)

var Version = "dev"

func main() {
	repl.Run(os.Stdin, os.Stdout)
}
