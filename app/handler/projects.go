package handler

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/flatcar-linux/flog/pkg/db"

	"github.com/gorilla/mux"
)

const projects = `<ul>{{ range $proj := .}}
{{ range .Versions }}
<li>{{ $proj.Name }}: <a href="/{{ $proj.Name | lower }}@{{ . | lower }}">{{ . }}</a></li>
{{ end }}
{{ end }}</ul>`

var (
	funcMap = template.FuncMap{
		"lower": strings.ToLower,
	}
)

type Projects struct {
	Name     string
	Versions []string
}

func GetAllProjects(db db.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p := db.GetProjects()
	tmpl, _ := template.New("projects").Funcs(funcMap).Parse(projects)

	rendered := make([]Projects, len(p))

	// TODO: this is super slow, do it into the DB.
	for i, project := range p {
		rendered[i].Name = project.Name

		versions := db.GetVersionsForProjects(project.ID)

		for _, version := range versions {
			rendered[i].Versions = append(rendered[i].Versions, version.Version)
		}

	}

	tmpl.Execute(w, rendered)
}

func GetProject(db db.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	respondJSON(w, http.StatusOK, id)
}
