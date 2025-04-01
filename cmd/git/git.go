package git

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type GitRepository struct {
	Name         string
	SizeInBytes  int64
	LastModified time.Time
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
			absolutePath := filepath.Join(dir, file.Name())
			if err != nil {
				log.Fatalf("Cannot open git repository with path %s because of an error: %v", absolutePath, err)
			}
			repositories = append(repositories, GitRepository{
				Name:         file.Name(),
				SizeInBytes:  file.Size(), //TODO: fix this
				LastModified: file.ModTime(),
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
