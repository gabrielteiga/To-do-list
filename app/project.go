package app

import (
	"fmt"
	"time"
)

type Project struct {
	Id           int
	Title        string
	Description  string
	CreationDate time.Time
	Tasks        map[int]Task
}

func NewProject(id int, title string, description string) *Project {
	return &Project{
		Id:           id,
		Title:        title,
		Description:  description,
		CreationDate: time.Now(),
		Tasks:        make(map[int]Task),
	}
}

func (p *Project) AddTask(task Task) {
	p.Tasks[task.Id] = task // BUG - Não funciona
}

func (p *Project) RemoveTask(task Task) {
	delete(p.Tasks, task.Id)
}

func (p *Project) GetTask(id int) Task {
	return p.Tasks[id]
}

func (p *Project) GetTasks() []Task {
	tasks := make([]Task, 0, len(p.Tasks))
	for _, task := range p.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (p *Project) PrintTasks() {
	for _, task := range p.Tasks {
		fmt.Printf("\n----------- Task %d -----------\n", task.Id)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Println("Completed? ", task.IsCompleted)
		fmt.Printf("Deadline (YYYY-MM-DD): %d-%d-%d\n", task.DueDate.Year(), task.DueDate.Month(), task.DueDate.Day())
		fmt.Printf("-------------------------------\n\n")
	}
}

func (p *Project) CompleteTask(task Task) {
	taskToComplete, ok := p.Tasks[task.Id]
	if ok {
		taskToComplete.Complete()
		p.Tasks[task.Id] = taskToComplete
	} else {
		fmt.Printf("Task %d not found", task.Id)
	}
}
