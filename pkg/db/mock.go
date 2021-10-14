package db

import "github.com/flatcar-linux/flog/app/model"

// MockDB used for testing purposes.
type MockDB struct{}

func (m *MockDB) GetProjects() []model.Project {
	return []model.Project{
		model.Project{Name: "flatcar-linux/mantle"},
	}
}

func (m *MockDB) GetProject(id uint) *model.Project {
	return &model.Project{Name: "flatcar-linux/mantle"}
}

// GetVersionsForProjects returns the available version for a given project.
func (m *MockDB) GetVersionsForProjects(id uint) []model.ProjectVersion {
	return []model.ProjectVersion{
		model.ProjectVersion{
			ProjectID: id,
			Version:   "v0.16.0",
		},
		model.ProjectVersion{
			ProjectID: id,
			Version:   "v0.17.0",
		},
	}
}
