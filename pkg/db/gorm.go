package db

import (
	"fmt"

	"github.com/flatcar/flog/app/model"
	"github.com/flatcar/flog/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GormDB is one database implementation
type GormDB struct{ db *gorm.DB }

// NewGorm creates and configures a gorm compliant database.
func NewGorm(cfg *config.Config) (*GormDB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("opening DB: %w", err)
	}

	if err := db.AutoMigrate(&model.Project{}, &model.Changelog{}, &model.ProjectVersion{}); err != nil {
		return nil, fmt.Errorf("doing migration: %w", err)
	}

	return &GormDB{db: db}, nil
}

func (g *GormDB) GetProjects() []model.Project {
	projects := []model.Project{}
	// TODO: handle transaction
	g.db.Find(&projects)
	return projects
}

func (g *GormDB) GetProject(id uint) model.Project {
	project := model.Project{}
	// TODO: handle transaction
	g.db.First(&project, id)
	return project
}

func (g *GormDB) GetVersionsForProjects(id uint) []model.ProjectVersion {
	// TODO: implement me :)
	return nil
}

// GetConnector returns the connector for a project.
func (g *GormDB) GetConnector(project string) string {
	// TODO: implement me :)
	return ""
}
