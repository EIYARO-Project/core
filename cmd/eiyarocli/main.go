package main

import (
	"runtime"

	cmd "core/cmd/eiyarocli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
