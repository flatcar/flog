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
