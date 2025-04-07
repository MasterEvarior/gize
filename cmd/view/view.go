package view

import (
	_ "embed"
	"html/template"
	"net/http"

	"github.com/MasterEvarior/gize/cmd/git"
	"github.com/MasterEvarior/gize/cmd/helper"
)

//go:embed templates/overview.html
var overviewTemplate string

type templateData struct {
	Title        string
	Description  string
	Footer       template.HTML
	Repositories []git.GitRepository
}

func Overview(w http.ResponseWriter, r *http.Request) {
	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, _ := git.GetAllRepositories(rootDir)
	data := getTemplateData(repositories)

	tmpl := template.Must(template.New("overview").Parse(overviewTemplate))
	tmpl.Execute(w, data)
}

func getTemplateData(additionalData []git.GitRepository) templateData {
	applicationTitle := helper.GetEnvVarWithDefault("GIZE_TITLE", "Gize")
	applicationDescription := helper.GetEnvVarWithDefault("GIZE_DESCRIPTION", "Your local Git repository browser")
	applicationFooter := helper.GetEnvVarWithDefault("GIZE_FOOTER", "Made with ❤️ and published on <a href='https://github.com/MasterEvarior/gize'>GitHub</a>")

	return templateData{
		Title:        applicationTitle,
		Description:  applicationDescription,
		Footer:       template.HTML(applicationFooter),
		Repositories: additionalData,
	}
}
