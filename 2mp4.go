package main

import (
	"github.com/weakish/goaround"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		os.Stderr.WriteString("Please specify input file name.\n")
		os.Exit(64)
	} else {
		var inputFileName string = os.Args[1]
		var basename []string = strings.Split(inputFileName, ".")
		var outputFileName string = strings.Join(basename, "")
		args := []string{"-i", inputFileName, "-c", "copy", outputFileName}
		err := syscall.Exec(which("ffmpeg"), args, os.Environ())

		goaround.PanicIf(err)
	}
}
func which(command string) string {
	var commandPath string
	commandPath, err := exec.LookPath(command)
	goaround.PanicIf(err)

	return commandPath
}