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
	project_id int64
	text       string
}

type ProjectVersion struct {
	gorm.Model
	project_id int64
	version    string
}
