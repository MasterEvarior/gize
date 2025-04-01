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

	//tmpl := template.Must(template.New("overview").Parse(overviewTemplate))
	tmpl := renderTemplate()["overview"]
	tmpl.Execute(w, repositories)
}

func renderTemplate() map[string]*template.Template {
	base, _ := template.New("base").Parse(baseTemplate)

	tmpl := make(map[string]*template.Template)
	tmpl["overview"] = template.Must(base.Parse(overviewTemplate))
	return tmpl
}
