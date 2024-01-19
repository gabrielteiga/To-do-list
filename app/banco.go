// Simulating a database
package app

import (
	"fmt"
)

type Banco struct {
	Projects []Project
}

func NewBanco() *Banco {
	return &Banco{
		Projects: make([]Project, 0),
	}
}

func (b *Banco) AddProject(p Project) {
	b.Projects = append(b.Projects, p)
}

func (b *Banco) PrintAllProjects() {
	fmt.Printf("\n---------- All Projects ----------\n")
	for _, project := range b.Projects {
		fmt.Printf("%s - ID %d\n", project.Title, project.Id)
		fmt.Printf("Description: %s\n", project.Description)
		fmt.Printf("Created on (YYYY-MM-DD): %d-%d-%d\n\n", project.CreationDate.Year(), project.CreationDate.Month(), project.CreationDate.Day())
	}
	fmt.Printf("----------------------------------\n\n")
}

func (b *Banco) GetProjectById(id int) (Project, error) {
	for _, project := range b.Projects {
		if project.Id == id {
			return project, nil
		}
	}
	return *NewProject(0, "", ""), nil // Realizar um tratamento de erro.
}
