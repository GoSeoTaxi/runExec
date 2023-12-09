package main

import (
	"os"

	"runExec/cmd"
)

func main() {
	cfg, err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	cmd.Work(cfg)
}
