package view

import (
	"html/template"
	"net/http"

	"github.com/MasterEvarior/gize/cmd/git"
	"github.com/MasterEvarior/gize/cmd/helper"
)

var overviewTemplate = `
	<h1>Overview<h1>
	<ul>
		{{ range . }}
		<li>{{ .Name }}</li>
		{{ end }}
	</ul>
`

func Overview(w http.ResponseWriter, r *http.Request) {
	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, _ := git.GetAllRepositories(rootDir)

	tmpl := template.Must(template.New("overview").Parse(overviewTemplate))
	tmpl.Execute(w, repositories)
}
