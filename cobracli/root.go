package cobracli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task_tracker_cli",
	Short: "Track tasks from the command line (Cobra UI)",
	Long:  `Add, list, update, delete, and mark tasks. Data is stored in task_master.json in the current directory.`,
}

// Execute runs the Cobra command tree.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
