package main

import (
	"fmt"
	"time"

	"github.com/gabrielteiga/kanban/app"
)

func main() {
	BancoDeDados := app.NewBanco()
	projectId := 1
	taskId := 1
	fmt.Println("Welcome to Kanban!")

	for {
		BancoDeDados.PrintAllProjects()
		ShowMenu()

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			title, description := getNewProjectData()
			project := app.NewProject(projectId, title, description)
			BancoDeDados.AddProject(*project)

			projectId++
			fmt.Println("\nProject created!")
		case 2:
			title, description, dueDate, idProject := getNewTaskData()
			task := app.NewTask(taskId, title, description, dueDate, idProject)
			project, _ := BancoDeDados.GetProjectById(idProject)
			project.AddTask(*task)

			taskId++
			fmt.Println("\nTask created!")
		case 3:
			var projectId int
			fmt.Println("Digit the project Id: ")
			fmt.Scanln(&projectId)
			project, _ := BancoDeDados.GetProjectById(projectId)

			project.PrintTasks()
			fmt.Println("List all tasks of a project")
		case 4:
			fmt.Println("Complete a task")
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Printf("\n\nPress enter to continue...")
		fmt.Scanln()
	}
}

func ShowMenu() {
	fmt.Println("1. Create a new project")
	fmt.Println("2. Create a new task")
	fmt.Println("3. List all tasks of a project")
	fmt.Println("4. Mark a task as completed")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func getNewProjectData() (string, string) {
	var title string
	var description string

	fmt.Print("Title: ")
	fmt.Scanln(&title) //BUG: titulos com espaço contam como outras entradas.

	fmt.Print("Description: ")
	fmt.Scanln(&description)

	return title, description
}

func getNewTaskData() (string, string, time.Time, int) {
	var title, description, dueDateString string
	var idProject int

	fmt.Print("What's the project ID of this task? ")
	fmt.Scanln(&idProject)

	fmt.Print("Title: ")
	fmt.Scanln(&title)

	fmt.Print("Description: ")
	fmt.Scanln(&description)

	fmt.Print("Due date (YYYY-MM-DD): ")
	fmt.Scanln(&dueDateString)
	dueDate, _ := parseStringToDate(dueDateString)

	return title, description, dueDate, idProject
}

func parseStringToDate(dateString string) (time.Time, error) {
	return (time.Parse("2006-01-02", dateString))
}
