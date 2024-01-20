package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gabrielteiga/kanban/app"
)

func main() {
	Database := app.NewBanco()
	projectId := 1
	taskId := 1
	fmt.Println("Welcome to Kanban!")

	for {
		Database.PrintAllProjects()
		ShowMenu()

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			title, description := getNewProjectData()
			project := app.NewProject(projectId, title, description)
			Database.AddProject(*project)

			projectId++
			fmt.Println("\nProject created!")
		case 2:
			title, description, dueDate, idProject := getNewTaskData()
			task := app.NewTask(taskId, title, description, dueDate, idProject)
			project, _ := Database.GetProjectById(idProject)
			project.AddTask(*task)

			taskId++
			fmt.Println("\nTask created!")
		case 3:
			var projectId int

			fmt.Print("Digit the project ID: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			projectId, _ = strconv.Atoi(scanner.Text())
			project, _ := Database.GetProjectById(projectId)

			project.PrintTasks()
		case 4:
			fmt.Println("Complete a task")
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Printf("\nPress enter to continue...")
		fmt.Scanln()
		fmt.Print("\n\n\n\n")
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Description: ")
	scanner.Scan()
	description = scanner.Text()

	return title, description
}

func getNewTaskData() (string, string, time.Time, int) {
	var title, description, dueDateString string
	var idProject int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digit the project ID: ")
	scanner.Scan()
	idProject, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Description: ")
	scanner.Scan()
	description = scanner.Text()

	fmt.Print("Due date (YYYY-MM-DD): ")
	scanner.Scan()
	dueDateString = scanner.Text()
	dueDate, _ := parseStringToDate(dueDateString)

	return title, description, dueDate, idProject
}

func parseStringToDate(dateString string) (time.Time, error) {
	return (time.Parse("2006-01-02", dateString))
}
