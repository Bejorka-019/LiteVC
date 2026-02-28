package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

const RepoDir = ".litevc"

func Init() {
	if _, err := os.Stat(RepoDir); err == nil {
		fmt.Println("LiteVC repository already exists.")
		return
	}

	err := os.Mkdir(RepoDir, 0755)

	if err != nil {
		fmt.Println("Something error happen: ", err)
		return
	}

	foldersPath := []string{
		filepath.Join(RepoDir, "objects"),
		filepath.Join(RepoDir, "commits"),
	}

	for _, path := range foldersPath {
		err := os.MkdirAll(path, 0755)

		if err != nil {
			fmt.Println("Something error happen: ", err)
			return
		}
	}

	indexFile := filepath.Join(RepoDir, "index")
	if err := os.WriteFile(indexFile, []byte{}, 0644); err != nil {
		fmt.Println("Something error happen: ", err)
		return
	}

	headFile := filepath.Join(RepoDir, "HEAD")
	if err := os.WriteFile(headFile, []byte(""), 0644); err != nil {
		fmt.Println("Something error happen: ", err)
		return
	}

	fmt.Println("LiteVC repository initialized successfully.")
}
