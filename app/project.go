package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	p.Tasks[task.Id] = task
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

func (p *Project) ShowMenu() {
	for {
		var option int

		fmt.Printf("\n----------- %s -----------\n", p.Title)
		p.printTasks()
		fmt.Println("1. Mark a task as completed")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, _ = strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			var idTask int

			fmt.Print("Digit the task ID: ")
			scanner.Scan()
			idTask, _ = strconv.Atoi(scanner.Text())

			task := p.GetTask(idTask)
			p.CompleteTask(task)
		case 2:
			return
		default:
			fmt.Println("Invalid option")
		}
	}

}

func (p *Project) printTasks() {
	for _, task := range p.Tasks {
		fmt.Printf("\n----------- Task %d -----------\n", task.Id)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Println("Completed? ", task.IsCompleted)
		fmt.Printf("Deadline (YYYY-MM-DD): %d-%d-%d\n", task.DueDate.Year(), task.DueDate.Month(), task.DueDate.Day())
		fmt.Printf("-------------------------------\n")
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
