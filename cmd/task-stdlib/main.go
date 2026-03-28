package main

import (
	"fmt"
	"os"

	"task_tracker_cli/internal/handlers"
)

func main() {
	if len(os.Args) < 2 {
		handlers.PrintUsage()
		os.Exit(1)
	}
	if err := handlers.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
