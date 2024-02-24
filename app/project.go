package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gabrielteiga/to-do-list/internal/db_todo"
	"github.com/jackc/pgx/v5/pgtype"
)

type Project struct {
	Id           int32
	Title        string
	Description  string
	CreationDate time.Time
}

func NewProject(project db_todo.Project) *Project {
	return &Project{
		Id:           int32(project.ID),
		Title:        project.Title,
		Description:  project.Describe.String,
		CreationDate: project.CreatedAt.Time,
	}
}

func (p *Project) ShowMenu(ctx context.Context, dt *db_todo.Queries) {
	for {
		var option int

		fmt.Printf("\n----------- %s -----------\n", p.Title)
		fmt.Printf("Description: %s\n\n", p.Description)
		p.printTasks(ctx, dt, p.Id)
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as completed")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Delete Project")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, _ = strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			err := p.addTask(ctx, dt)
			if err != nil {
				fmt.Println("Error creating task: ", err)
			}

		case 2:
			err := p.completeTask(ctx, dt)
			if err != nil {
				fmt.Println("Error marking task as completed: ", err)
			}

		case 3:
			err := p.deleteTask(ctx, dt)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			}

		case 4:
			fmt.Printf("Are you sure you want to delete the project %s? (y/n): ", p.Title)
			scanner.Scan()
			res := scanner.Text()
			if res != "n" {
				fmt.Println("Operation canceled")
				break
			}

			err := p.deleteProject(ctx, dt)
			if err != nil {
				fmt.Println("Error deleting project: ", err)
			}
			return

		case 5:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}

func (p *Project) deleteProject(ctx context.Context, dt *db_todo.Queries) error {
	err := dt.DeleteProject(ctx, p.Id)
	if err != nil {
		return err
	}

	fmt.Println("Project deleted successfully!")
	return nil
}

func (p *Project) addTask(ctx context.Context, dt *db_todo.Queries) error {
	var title string
	var description string
	var dueDate time.Time

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Description: ")
	scanner.Scan()
	description = scanner.Text()

	fmt.Print("Due date (YYYY-MM-DD): ")
	scanner.Scan()
	dueDate, err := time.Parse("2006-01-02", scanner.Text())
	if err != nil {
		fmt.Println("Invalid date, try again!")
		return err
	}

	task := db_todo.CreateTaskParams{
		Title:       title,
		Describe:    pgtype.Text{String: description, Valid: true},
		IsCompleted: false,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Deadline:    pgtype.Timestamp{Time: dueDate, Valid: true},
		ProjectID:   p.Id,
	}

	err = dt.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	fmt.Println("Task created successfully!")
	return nil
}

func (p *Project) deleteTask(ctx context.Context, dt *db_todo.Queries) error {
	var idTask int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digit the task ID: ")
	scanner.Scan()
	idTask, _ = strconv.Atoi(scanner.Text())

	err := dt.DeleteTask(ctx, db_todo.DeleteTaskParams{
		ID:        int32(idTask),
		ProjectID: p.Id,
	})
	if err != nil {
		return err
	}

	fmt.Println("Task deleted successfully!")
	return nil
}

func (p *Project) printTasks(ctx context.Context, dt *db_todo.Queries, projectId int32) error {
	projects, err := dt.GetTasksByProject(ctx, projectId)
	if err != nil {
		fmt.Println("Error getting tasks: ", err)
		return err
	}

	for _, task := range projects {
		var status string
		if task.IsCompleted {
			status = "Completed"
		} else {
			status = "Pending"
		}

		fmt.Printf("TASK ID %d - %s - %s\nDescription: %s\n\n", task.ID, task.Title, status, task.Describe.String)
	}

	return nil
}

func (p *Project) completeTask(ctx context.Context, dt *db_todo.Queries) error {
	var idTask int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digit the task ID: ")
	scanner.Scan()
	idTask, _ = strconv.Atoi(scanner.Text())

	err := dt.MarkTaskAsCompleted(ctx, db_todo.MarkTaskAsCompletedParams{
		ID:        int32(idTask),
		ProjectID: p.Id,
	})

	if err != nil {
		return err
	}

	return nil
}
