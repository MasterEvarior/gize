package git

import (
	"log"
	"os"
	"path/filepath"
)

type GitRepository struct {
	Name string
}

func GetAllRepositories(dir string) ([]GitRepository, error) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Cannot open directory %s because of an error: %v", dir, err)
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatalf("Cannot read directory %s because of an error: %v", f.Name(), err)
	}

	var repositories []GitRepository
	for _, file := range fileInfo {
		if isAGitRepository(dir, file) {
			repositories = append(repositories, GitRepository{
				Name: file.Name(),
			})
		}
	}

	return repositories, nil
}

func isAGitRepository(root string, fileInfo os.FileInfo) bool {
	if fileInfo.IsDir() == false {
		return false
	}

	absolutePath := filepath.Join(root, fileInfo.Name(), ".git")
	_, err := os.Stat(absolutePath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
