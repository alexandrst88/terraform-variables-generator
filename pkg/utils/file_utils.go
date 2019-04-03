package utils

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// GetAllFiles will get all files in current directory
func GetAllFiles(ext string) ([]string, error) {
	dir, err := os.Getwd()
	CheckError(err)

	var files []string
	log.Infof("Finding files in %q directory", dir)
	files, err = filepath.Glob(ext)
	CheckError(err)

	if len(files) == 0 {
		log.Infof("No files with %q extensions found in %q", ext, dir)
	}
	return files, nil
}

// FileExists checks if file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
