package main

import (
	"github.com/weakish/goaround"
	"github.com/weakish/gosugar"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		goaround.ErrPrint("Please specify input file name.")
		os.Exit(64)
	} else {
		var inputFileName string = os.Args[1]
		var basename []string = strings.Split(inputFileName, ".")
		var outputFileName string = strings.Join(basename, "")
		args := []string{"ffmpeg", "-i", inputFileName, "-c", "copy", outputFileName}
		gosugar.Exec(args)
	}
}
