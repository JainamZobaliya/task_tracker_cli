package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const TASK_MASTER_FILE = "task_master.json"

func readStore() (TaskStore, error) {
	file, err := os.OpenFile(TASK_MASTER_FILE, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return TaskStore{}, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return TaskStore{}, err
	}

	if stat.Size() == 0 {
		return TaskStore{
			Tasks:  []Task{},
			LastID: -1,
		}, nil
	}

	data, err := os.ReadFile(TASK_MASTER_FILE)
	if err != nil {
		return TaskStore{}, err
	}

	var store TaskStore
	err = json.Unmarshal(data, &store)
	if err != nil {
		return TaskStore{}, err
	}

	return store, nil
}

func writeStore(store TaskStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(TASK_MASTER_FILE, data, 0666)
}

func ListTasks(status *string) ([]Task, error) {
	store, err := readStore()
	if err != nil {
		return nil, err
	}

	// No filter → return all
	if status == nil {
		return store.Tasks, nil
	}

	correctStatus, err := parseStatus(*status)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, val := range store.Tasks {
		if val.Status == correctStatus {
			tasks = append(tasks, val)
		}
	}

	return tasks, nil
}

func CreateTask(taskDescription string) (int, error) {
	store, err := readStore()
	if err != nil {
		return -1, err
	}

	// Generate next ID
	nextID := store.LastID + 1
	fmt.Println("nextID: ", nextID)

	newTask := Task{
		Id:          nextID,
		Description: taskDescription,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
		Status:      ToDo,
	}

	// Append task
	store.Tasks = append(store.Tasks, newTask)
	fmt.Println("store.Tasks: ", store.Tasks)

	// Update last ID
	store.LastID = nextID

	// Write back
	err = writeStore(store)
	if err != nil {
		return -1, err
	}
	return nextID, nil
}

func DeleteTask(taskId int) (bool, error) {
	store, err := readStore()
	if err != nil {
		return false, err
	}

	for ind := range store.Tasks{
		if(store.Tasks[ind].Id == taskId) {
			store.Tasks = append(store.Tasks[:ind], store.Tasks[ind+1:]...)
			err = writeStore(store)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func UpdateTaskDescription(taskId int, taskDescription string) (bool, error) {
	store, err := readStore()
	if err != nil {
		return false, err
	}
	tasks := store.Tasks
	for index := range tasks{
		if(store.Tasks[index].Id == taskId) {
			store.Tasks[index].Description = taskDescription
			store.Tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)
			// Write back
			err = writeStore(store)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func UpdateTaskStatus(taskId int, taskStatus TaskStatus) (bool, error) {
	store, err := readStore()
	if err != nil {
		return false, err
	}
	tasks := store.Tasks
	for index := range tasks{
		if(store.Tasks[index].Id == taskId) {
			store.Tasks[index].Status = taskStatus
			store.Tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)
			// Write back
			err = writeStore(store)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}
