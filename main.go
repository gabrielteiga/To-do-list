package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gabrielteiga/to-do-list/app"
	"github.com/gabrielteiga/to-do-list/internal/db_todo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/lib/pq"
)

const (
	user     = "sqlcgabrielteiga"
	password = "123abc"
	dbname   = "sqlc"
	sslmode  = "disable"
	port     = 5432
)

func main() {
	ctx := context.Background()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%d", user, password, dbname, sslmode, port)
	dbcon, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbcon.Close(ctx)

	err = executeSchema(dbcon, "sql/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	dt := db_todo.New(dbcon)

	fmt.Printf("----- Welcome to Kanban! -----\n\n")
	for {
		projects, _ := dt.GetProjects(ctx)
		for _, project := range projects {
			fmt.Printf("PROJECT ID: %d - %s\n", project.ID, project.Title)
		}

		fmt.Println("")
		ShowMenu()

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			title, description := getNewProjectData()
			err := dt.CreateProject(ctx, db_todo.CreateProjectParams{
				Title: title,
				Describe: pgtype.Text{
					String: description,
					Valid:  true,
				},
				CreatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				}})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\nProject created!")

		case 2:
			var projectId int
			fmt.Print("Digit the project ID: ")
			fmt.Scanln(&projectId)
			project, err := dt.GetProject(ctx, int32(projectId))

			if err != nil {
				fmt.Println("Error getting project: ", err)
				break
			}

			app.NewProject(project).ShowMenu(ctx, dt)

		case 3:
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
	fmt.Println("2. Select a project")
	fmt.Println("3. Exit")
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

func executeSchema(conn *pgx.Conn, filename string) error {
	schema, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(schema))
	if err != nil {
		return err
	}

	return nil
}
