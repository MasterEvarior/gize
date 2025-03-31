package view

import (
	_ "embed"
	"html/template"
	"net/http"

	"github.com/MasterEvarior/gize/cmd/git"
	"github.com/MasterEvarior/gize/cmd/helper"
)

//go:embed templates/base.html
var baseTemplate string

//go:embed templates/overview.html
var overviewTemplate string

func Overview(w http.ResponseWriter, r *http.Request) {
	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, _ := git.GetAllRepositories(rootDir)

	tmpl := template.Must(template.New("overview").Parse(overviewTemplate))
	tmpl.Execute(w, repositories)
}
