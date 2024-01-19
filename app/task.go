package app

import "time"

type Task struct {
	Id           int
	Title        string
	Description  string
	IsCompleted  bool
	CreationDate time.Time
	DueDate      time.Time
	IdProject    int
}

func NewTask(id int, title string, description string, dueDate time.Time, idProject int) *Task {
	return &Task{
		Id:           id,
		Title:        title,
		Description:  description,
		IsCompleted:  false,
		CreationDate: time.Now(),
		DueDate:      dueDate,
		IdProject:    idProject,
	}
}

func (t *Task) Complete() {
	t.IsCompleted = true
}

func (t *Task) ExtendDueDate(dueDate time.Time) {
	t.DueDate = dueDate
}
