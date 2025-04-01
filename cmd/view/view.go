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

type templateData struct {
	Title       string
	Description string
	Data        interface{}
}

func Overview(w http.ResponseWriter, r *http.Request) {
	applicationTitle := helper.GetEnvVarWithDefault("GIZE_TITLE", "Gize")
	applicationDescription := helper.GetEnvVarWithDefault("GIZE_DESCRIPTION", "Your local Git repository browser")

	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, _ := git.GetAllRepositories(rootDir)

	tmpl := renderTemplate()["overview"]
	data := templateData{
		Title:       applicationTitle,
		Description: applicationDescription,
		Data:        repositories,
	}
	tmpl.Execute(w, data)
}

func renderTemplate() map[string]*template.Template {
	base, _ := template.New("base").Parse(baseTemplate)

	tmpl := make(map[string]*template.Template)
	tmpl["overview"] = template.Must(base.Parse(overviewTemplate))
	return tmpl
}
