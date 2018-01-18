package main

import (
	"syscall"
	"os"
	"os/exec"
	"strings"
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

		panicIf(err)
	}
}
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
func which(command string) string {
	var commandPath string
	commandPath, err := exec.LookPath(command)
	panicIf(err)

	return commandPath
}