package main

import (
	"runtime"

	cmd "ey/cmd/eiyarocli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
