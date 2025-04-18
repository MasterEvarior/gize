package git

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

type GitRepository struct {
	Name         string
	Size         string
	LastModified time.Time
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
			absolutePath := path.Join(dir, file.Name())
			totalSize, lastModified, _ := calculateRepositoryInfo(absolutePath)

			repositories = append(repositories, GitRepository{
				Name:         file.Name(),
				Size:         formatSize(totalSize),
				LastModified: lastModified,
			})
		}
	}
	return repositories, nil
}

func isAGitRepository(root string, fileInfo os.FileInfo) bool {
	if !fileInfo.IsDir() {
		return false
	}

	absolutePath := filepath.Join(root, fileInfo.Name(), ".git")
	_, err := os.Stat(absolutePath)
	return !os.IsNotExist(err)
}

func calculateRepositoryInfo(path string) (int64, time.Time, error) {
	var size int64
	var mostCurrentTimestamp time.Time
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		lastModified := info.ModTime()
		if lastModified.After(mostCurrentTimestamp) {
			mostCurrentTimestamp = lastModified
		}
		return err
	})
	return size, mostCurrentTimestamp, err
}

func formatSize(bytes int64) string {
	const bytesInKB = 1024
	const bytesInMB = 1024 * 1024
	if bytes < bytesInKB {
		return fmt.Sprintf("%d Bytes", bytes)
	}

	if bytes < bytesInMB {
		return fmt.Sprintf("%.1f KB", float64(bytes)/float64(bytesInKB))
	}

	return fmt.Sprintf("%.1f MB", float64(bytes)/float64(bytesInMB))
}
