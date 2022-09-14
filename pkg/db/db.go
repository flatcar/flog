package db

import "github.com/flatcar/flog/app/model"

// DB is the interface to implement to access the projects
// and other entites.
type DB interface {
	// GetProjects returns a list of projects.
	GetProjects() []model.Project
	// GetProject returns a project by its ID.
	GetProject(uint) *model.Project
	// GetVersionsForProjects returns the available version for a given project.
	GetVersionsForProjects(uint) []model.ProjectVersion
	// GetConnector returns the connector for a project.
	// TODO: result should be an enum
	GetConnector(string) string
}
