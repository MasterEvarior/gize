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

//go:embed templates/detail.html
var detailTemplate string

type templateData struct {
	Title       string
	Description string
	Footer      template.HTML
	Data        interface{}
}

func Overview(w http.ResponseWriter, r *http.Request) {
	base, _ := template.New("base").Parse(baseTemplate)

	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, _ := git.GetAllRepositories(rootDir)
	data := getTemplateData(repositories)

	template.Must(base.Parse(overviewTemplate)).Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	base, _ := template.New("base").Parse(baseTemplate)

	repoPath := r.PathValue("repository")
	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repository, _ := git.GetRepository(rootDir, repoPath)

	data := getTemplateData(repository)
	template.Must(base.Parse(detailTemplate)).Execute(w, data)
}

func getTemplateData(additionalData interface{}) templateData {
	applicationTitle := helper.GetEnvVarWithDefault("GIZE_TITLE", "Gize")
	applicationDescription := helper.GetEnvVarWithDefault("GIZE_DESCRIPTION", "Your local Git repository browser")
	applicationFooter := helper.GetEnvVarWithDefault("GIZE_FOOTER", "Made with ❤️ by <a href='https://github.com/MasterEvarior/gize'>MasterEvarior</a>")

	return templateData{
		Title:       applicationTitle,
		Description: applicationDescription,
		Footer:      template.HTML(applicationFooter),
		Data:        additionalData,
	}
}
