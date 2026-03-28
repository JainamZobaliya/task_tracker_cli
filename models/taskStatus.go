package models

import "fmt"

type TaskStatus string

const (
    ToDo TaskStatus = "todo"
    InProgress TaskStatus = "in-progress"
    Done TaskStatus = "done"
)

var allowedTaskStatus = map[string]TaskStatus{
    "todo":        ToDo,
    "in-progress": InProgress,
    "done":        Done,
}

func parseStatus(status string) (TaskStatus, error) {
    val, ok := allowedTaskStatus[status]
    if !ok {
        return "", fmt.Errorf("invalid status: %s", status)
    }
    return val, nil
}


