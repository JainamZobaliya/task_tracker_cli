package cobracli

import (
	"log"
	"task_tracker_cli/internal/handlers"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `Add a new task to the task tracker.

Provide the task description as an argument.

Example:
  task-cli add "Buy groceries"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.Add(args); err != nil {
			log.Fatal(err)
		}
	},
}

var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List all tasks in the system.

Optionally provide a status to filter:
- todo
- in-progress
- done

Examples:
  task-cli list
  task-cli list done
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.List(args); err != nil {
			log.Fatal(err)
		}
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task description",
	Long: `Update an existing task's description.

Provide task ID and new description.

Example:
  task-cli update 1 "Buy groceries and cook dinner"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.Update(args); err != nil {
			log.Fatal(err)
		}
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task using its ID.

Example:
  task-cli delete 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.Delete(args); err != nil {
			log.Fatal(err)
		}
	},
}

var markTaskInProgress = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark task as in-progress",
	Long: `Mark a task status as in-progress.

Example:
  task-cli mark-in-progress 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.MarkInProgress(args); err != nil {
			log.Fatal(err)
		}
	},
}

var markTaskCompleted = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark task as done",
	Long: `Mark a task status as done.

Example:
  task-cli mark-done 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.MarkDone(args); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(listTaskCmd)
	rootCmd.AddCommand(updateTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
	rootCmd.AddCommand(markTaskInProgress)
	rootCmd.AddCommand(markTaskCompleted)
}
