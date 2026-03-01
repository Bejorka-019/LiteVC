package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		//if the file is a .litevc file, skip it
		if strings.Contains(path, ".litevc") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		//if the file is a directory and starts with a dot, skip it
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") && d.Name() != "." {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
