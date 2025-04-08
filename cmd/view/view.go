package view

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/MasterEvarior/gize/cmd/git"
	"github.com/MasterEvarior/gize/cmd/helper"
)

//go:embed templates/overview.html
var overviewTemplate string

type templateData struct {
	Title          string
	Description    string
	Footer         template.HTML
	EnableDownload bool
	Repositories   []git.GitRepository
}

func Overview(w http.ResponseWriter, r *http.Request) {
	rootDir := helper.GetEnvVar("GIZE_ROOT")
	repositories, err := git.GetAllRepositories(rootDir)
	if err != nil {
		errMessage := "Could not gather necessary information about Git repositories"
		log.Printf("%s: %v", errMessage, err)
		http.Error(w, errMessage, 500)
		return
	}

	data := getTemplateData(repositories)

	tmpl := template.Must(template.New("overview").Parse(overviewTemplate))
	tmpl.Execute(w, data)
}

func Download(w http.ResponseWriter, r *http.Request) {
	repositoryName := r.PathValue("repository")
	rootDir := helper.GetEnvVar("GIZE_ROOT")

	buf, err := zipDirectory(path.Join(rootDir, repositoryName))
	if err != nil {
		errMessage := "Could not ZIP repository"
		log.Printf("%s: %v", errMessage, err)
		http.Error(w, errMessage, 500)
		return
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

func zipDirectory(source string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := zip.NewWriter(&buf)
	defer writer.Close()

	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func getTemplateData(additionalData []git.GitRepository) templateData {
	applicationTitle := helper.GetEnvVarWithDefault("GIZE_TITLE", "Gize")
	applicationDescription := helper.GetEnvVarWithDefault("GIZE_DESCRIPTION", "Your local Git repository browser")
	applicationFooter := helper.GetEnvVarWithDefault("GIZE_FOOTER", "Made with ❤️ and published on <a href='https://github.com/MasterEvarior/gize'>GitHub</a>")
	enableDownload := helper.IsEnabled("GIZE_ENABLE_DOWNLOAD")

	return templateData{
		Title:          applicationTitle,
		Description:    applicationDescription,
		Footer:         template.HTML(applicationFooter),
		EnableDownload: enableDownload,
		Repositories:   additionalData,
	}
}
