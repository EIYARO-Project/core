package main

import (
	"runtime"

	cmd "eiyaro/cmd/eiyarocli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
