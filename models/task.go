package models

type Task struct {
	Id int `json:"id"`
	Description string  `json:"description"`
	Status TaskStatus `json:"status"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type TaskStore struct {
    Tasks  []Task `json:"tasks"`
    LastID int    `json:"lastId"`
}