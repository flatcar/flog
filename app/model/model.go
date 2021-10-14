package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
}

type Changelog struct {
	gorm.Model
	ProjectID uint
	Text      string
}

type ProjectVersion struct {
	gorm.Model
	ProjectID uint
	Version   string
}
