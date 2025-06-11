package cli

import (
	"flag"
	"fmt"
	"runtime"
)

type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}

func RunParse(info VersionInfo, runApp func()) int {
    showVersion := flag.Bool("version", false, "show the current version")
    showShort := flag.Bool("v", false, "show the current version (short)")
    flag.Parse()

    if *showVersion || *showShort {
        printVersion(info)
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