package git

import (
	"fmt"
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

func GetRepository(dir string, name string) (*GitRepository, error) {
	all, err := GetAllRepositories(dir)

	if err != nil {
		return nil, err
	}

	for _, repo := range all {
		if repo.Name == name {
			return &repo, nil
		}
	}

	return nil, fmt.Errorf("repository %q not found", name)
}

func GetAllRepositories(dir string) ([]GitRepository, error) {
	f, err := os.Open(dir)
	if err != nil {
		log.Printf("Cannot open directory %s because of an error: %v", dir, err)
		return nil, err
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Printf("Cannot read directory %s because of an error: %v", f.Name(), err)
		return nil, err
	}

	var repositories []GitRepository
	for _, file := range fileInfo {
		if isAGitRepository(dir, file) {
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
