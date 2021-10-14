package handler

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/flatcar-linux/flog/pkg/db"

	"github.com/gorilla/mux"
)

const projects = `<ul>{{ range . }}
<li>{{ .Name }}</li>
{{ end }}</ul>`

var (
	funcMap = template.FuncMap{
		"lower": strings.ToLower,
	}
)

func GetAllProjects(db db.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p := db.GetProjects()
	tmpl, _ := template.New("projects").Funcs(funcMap).Parse(projects)
	tmpl.Execute(w, p)
}

func GetProject(db db.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	respondJSON(w, http.StatusOK, id)
}
