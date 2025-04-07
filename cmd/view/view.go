package view

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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

func Download(w http.ResponseWriter, r *http.Request) {
	repositoryName := r.PathValue("repository")
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	data, err := os.ReadFile(repositoryName)
	if err != nil {
		log.Fatal(err)
	}
	f, err := writer.Create(repositoryName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", repositoryName))
	w.Write(buf.Bytes())
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
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
