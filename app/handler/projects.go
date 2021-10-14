package handler

import (
	"net/http"

	"github.com/flatcar-linux/flog/app/model"
	"gorm.io/gorm"
)

func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}
