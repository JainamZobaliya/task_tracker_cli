package handlers

import (
	"fmt"
	"strconv"

	"task_tracker_cli/models"
)

// PrintUsage writes command help to stdout.
func PrintUsage() {
	fmt.Println(`Task tracker — commands:
  add <description>
  list [todo|in-progress|done]
  update <id> <description>
  delete <id>
  mark-in-progress <id>
  mark-done <id>`)
}

// Add creates a task from positional args (single description string).
func Add(args []string) error {
	if len(args) == 0 {
		fmt.Println("Task Description is required")
		return nil
	}
	if len(args) == 1 && args[0] != "" {
		id, err := models.CreateTask(args[0])
		if err != nil {
			return err
		}
		fmt.Println("Task Created Successfully, id::", id)
		return nil
	}
	fmt.Println("Too many Arguments passed")
	return nil
}

// List prints all tasks or filters by optional status.
func List(args []string) error {
	if len(args) == 0 {
		result, err := models.ListTasks(nil)
		if err != nil {
			return err
		}
		fmt.Println("Task List :: ", result)
		return nil
	}
	if len(args) == 1 && args[0] != "" {
		status := args[0]
		result, err := models.ListTasks(&status)
		if err != nil {
			return err
		}
		fmt.Println("Task List :: ", result)
		return nil
	}
	fmt.Println("Too many Arguments passed")
	return nil
}

// Update changes a task description (id + new description).
func Update(args []string) error {
	if len(args) < 2 {
		fmt.Println("Task Id & Description is required")
		return nil
	}
	if len(args) == 2 && args[0] != "" {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}
		desc := args[1]
		ok, err := models.UpdateTaskDescription(taskID, desc)
		if err != nil {
			return err
		}
		if ok {
			fmt.Println("Task Updated Successfully")
			return nil
		}
		fmt.Println("Task Not Updated / not found..")
		return nil
	}
	fmt.Println("Too many Arguments passed")
	return nil
}

// Delete removes a task by id.
func Delete(args []string) error {
	if len(args) == 0 {
		fmt.Println("Task Id is required")
		return nil
	}
	if len(args) == 1 && args[0] != "" {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}
		ok, err := models.DeleteTask(taskID)
		if err != nil {
			return err
		}
		if ok {
			fmt.Printf("Task %v Deleted Successfully\n", taskID)
			return nil
		}
		fmt.Println("Mentioned Task not found")
		return nil
	}
	fmt.Println("Too many Arguments passed")
	return nil
}

// MarkInProgress sets task status to in-progress.
func MarkInProgress(args []string) error {
	if len(args) == 0 {
		fmt.Println("Task Id is required")
		return nil
	}
	if len(args) > 1 {
		fmt.Println("Too many Arguments passed")
		return nil
	}
	if args[0] == "" {
		fmt.Println("Task Id is required")
		return nil
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task id: %w", err)
	}
	ok, err := models.UpdateTaskStatus(taskID, models.InProgress)
	if err != nil {
		return err
	}
	if ok {
		fmt.Printf("Task %v Updated to in-progress\n", taskID)
		return nil
	}
	fmt.Println("Mentioned Task not found")
	return nil
}

// MarkDone sets task status to done.
func MarkDone(args []string) error {
	if len(args) == 0 {
		fmt.Println("Task Id is required")
		return nil
	}
	if len(args) > 1 {
		fmt.Println("Too many Arguments passed")
		return nil
	}
	if args[0] == "" {
		fmt.Println("Task Id is required")
		return nil
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task id: %w", err)
	}
	ok, err := models.UpdateTaskStatus(taskID, models.Done)
	if err != nil {
		return err
	}
	if ok {
		fmt.Printf("Task %v Updated to done\n", taskID)
		return nil
	}
	fmt.Println("Mentioned Task not found")
	return nil
}

// Run parses argv[1:] as a single subcommand dispatch (stdlib mode).
func Run(argv []string) error {
	if len(argv) < 1 {
		PrintUsage()
		return fmt.Errorf("no command given")
	}
	switch argv[0] {
	case "add":
		return Add(argv[1:])
	case "list":
		return List(argv[1:])
	case "update":
		return Update(argv[1:])
	case "delete":
		return Delete(argv[1:])
	case "mark-in-progress":
		return MarkInProgress(argv[1:])
	case "mark-done":
		return MarkDone(argv[1:])
	case "help", "-h", "--help":
		PrintUsage()
		return nil
	default:
		return fmt.Errorf("unknown command: %s", argv[0])
	}
}
