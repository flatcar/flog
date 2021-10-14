package handler

import (
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/flatcar-linux/flog/pkg/connector"
	"github.com/flatcar-linux/flog/pkg/db"

	"github.com/gorilla/mux"
)

const (
	projects = `<ul>{{ range $proj := .}}
{{ range .Versions }}
<li>{{ $proj.Name }}: <a href="/projects/{{ $proj.Name | lower | escape }}@{{ . | lower }}">{{ . }}</a></li>
{{ end }}
{{ end }}</ul>`
	change = `<h1>{{ .Name }} - {{ .Version }}</h1><hr/>
{{ .Content }}
<br />
<a href="{{ .Source }}">source</a>`
)

var (
	funcMap = template.FuncMap{
		"lower":  strings.ToLower,
		"escape": url.QueryEscape,
	}

	// connectors is the list of available connectors
	connectors = map[string]connector.Connector{
		"github": &connector.GH{},
	}
)

// helpers for templates

// Projects
type Projects struct {
	Name     string
	Versions []string
}

// Changelog
type Changelog struct {
	Name    string
	Version string
	Content template.HTML
	Source  string
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

	id, ok := vars["id"]
	if !ok {
		http.Error(w, "id is not here", 406)
	}

	// TODO: handle me
	id, _ = url.QueryUnescape(id)

	project := strings.Split(id, "@")
	if len(project) != 2 {
		// TODO: handle me
		http.Error(w, "id should be project@version", 406)
	}

	connectorName := db.GetConnector(project[0])
	conn, ok := connectors[connectorName]
	if !ok {
		// TODO: handle me
		http.Error(w, "connector is not yet implemented", 406)
	}

	// TODO: handle me
	content, URL, _ := conn.FetchChangelog(project[0], project[1])

	changelog := &Changelog{
		Name:    project[0],
		Version: project[1],
		Content: template.HTML(content),
		Source:  URL,
	}

	// TODO: handle me
	tmpl, _ := template.New("change").Funcs(funcMap).Parse(change)
	tmpl.Execute(w, changelog)
}
