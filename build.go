package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	showBuild = flag.Bool("build", false, "Show the build")
	build     = "unknown, please recompile with -ldflags=\"-X main.build `date -u +%Y%m%d.%H%M%S`\""
)

func init() {
	if *showBuild {
		fmt.Printf("Build: %s\n", build)
		os.Exit(0)
	}
}
