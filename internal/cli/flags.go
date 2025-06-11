package cli

import (
	"flag"
	"fmt"
	"runtime"
	"calq/internal/calculator"
)

type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}

func RunParse(info VersionInfo, runApp func()) int {
    showVersion := flag.Bool("version", false, "show the current version")
    showShort := flag.Bool("v", false, "show the current version (short)")
	evalMode := flag.String("eval", "", "run in eval mode with a single expression (e.g. '1+1')")

    flag.Parse()

    if *showVersion || *showShort {
        printVersion(info)
        return 0
    }

    if *evalMode != "" {
		result, err := calculator.Calculate("", *evalMode)
		if err != nil {
			fmt.Printf("Calculation error: %v\n", err)
			return 1
		}

		fmt.Printf("%.2f\n", result)
        return 0
    }

    runApp()
    return 0
}


func printVersion(info VersionInfo) {
	fmt.Printf("calq %s (%s, %s) [%s-%s]\n",
		info.Version,
		info.Date,
		info.Commit,
		runtime.GOARCH,
		runtime.GOOS,
	)
}