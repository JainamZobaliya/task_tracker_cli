package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task_tracker_cli/models"

	"github.com/spf13/cobra"
)

// -------------------- ADD TASK --------------------

// addTaskCmd adds a new task with given description
var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `Add a new task to the task tracker.

Provide the task description as an argument.

Example:
  task-cli add "Buy groceries"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addTask called")
		fmt.Println("Args called: ", args)

		if len(args) == 0 {
			fmt.Println("Task Description is required")
			return
		}
		if len(args) == 1 && args[0] != "" {
			var taskDescription string = args[0]
			result, err := models.CreateTask(taskDescription)
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}
			fmt.Println("Task Created Successfully, id::", result)
			return
		}
		fmt.Println("Too many Arguments passed")
	},
}

// -------------------- LIST TASKS --------------------

// listTaskCmd lists all tasks or filters by status
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
		fmt.Println("listTask called")
		fmt.Println("Args called: ", args)

		if len(args) == 0 {
			result, err := models.ListTasks(nil)
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}
			fmt.Println("Task List :: ", result)
			return
		}
		if len(args) == 1 && args[0] != "" {
			var taskStatus string = args[0]
			result, err := models.ListTasks(&taskStatus)
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}
			fmt.Println("Task List :: ", result)
			return
		}
		fmt.Println("Too many Arguments passed")
	},
}

// -------------------- UPDATE TASK --------------------

// updateTaskCmd updates the description of a task
var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task description",
	Long: `Update an existing task's description.

Provide task ID and new description.

Example:
  task-cli update 1 "Buy groceries and cook dinner"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTask called")
		fmt.Println("Args called: ", args)

		if len(args) < 2 {
			fmt.Println("Task Id & Description is required")
			return
		}
		if len(args) == 2 && args[0] != "" {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Error: ", err)
			}
			var taskDescription string = args[1]
			result, err := models.UpdateTaskDescription(taskId, taskDescription)
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}
			if result {
				fmt.Println("Task Updated Successfully")
				return
			}
			fmt.Println("Task Not Updated / not found..")
			return
		}
		fmt.Println("Too many Arguments passed")
	},
}

// -------------------- DELETE TASK --------------------

// deleteTaskCmd deletes a task by ID
var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task using its ID.

Example:
  task-cli delete 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteTask called")
		fmt.Println("Args called: ", args)

		if len(args) == 0 {
			fmt.Println("Task Id is required")
			return
		}
		if len(args) == 1 && args[0] != "" {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Error: ", err)
			}
			result, err := models.DeleteTask(taskId)
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}
			if result {
				fmt.Printf("Task %v Deleted Successfully\n", taskId)
				return
			}
			fmt.Println("Mentioned Task not found")
			return
		}
		fmt.Println("Too many Arguments passed")
	},
}

// -------------------- MARK IN PROGRESS --------------------

// markTaskInProgress marks a task as in-progress
var markTaskInProgress = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark task as in-progress",
	Long: `Mark a task status as in-progress.

Example:
  task-cli mark-in-progress 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mark-in-progress called")
		fmt.Println("Args called: ", args)

		if len(args) == 0 {
			fmt.Println("Task Id is required")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many Arguments passed")
			return
		}
		if args[0] == "" {
			fmt.Println("Task Id is required")
			return
		}
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		result, err := models.UpdateTaskStatus(taskId, models.InProgress)
		if err != nil {
			log.Fatal("Something went wrong: ", err)
		}
		if result {
			fmt.Printf("Task %v Updated to in-progress\n", taskId)
			return
		}
		fmt.Println("Mentioned Task not found")
	},
}

// -------------------- MARK DONE --------------------

// markTaskCompleted marks a task as done
var markTaskCompleted = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark task as done",
	Long: `Mark a task status as done.

Example:
  task-cli mark-done 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mark-done called")
		fmt.Println("Args called: ", args)

		if len(args) == 0 {
			fmt.Println("Task Id is required")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many Arguments passed")
			return
		}
		if args[0] == "" {
			fmt.Println("Task Id is required")
			return
		}
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		result, err := models.UpdateTaskStatus(taskId, models.Done)
		if err != nil {
			log.Fatal("Something went wrong: ", err)
		}
		if result {
			fmt.Printf("Task %v Updated to done\n", taskId)
			return
		}
		fmt.Println("Mentioned Task not found")
	},
}

// -------------------- INIT --------------------

func init() {
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(listTaskCmd)
	rootCmd.AddCommand(updateTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
	rootCmd.AddCommand(markTaskInProgress)
	rootCmd.AddCommand(markTaskCompleted)
}