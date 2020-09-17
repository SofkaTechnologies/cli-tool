package utils

import (
	"github.com/hashicorp/go-getter"
	"github.com/plus3it/gorecurcopy"
	"os"
	"path/filepath"
)

func walkFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func CreateDirectoryStructure(destination string, language string, fileName string, tp ProjectSettings) {
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		panic(err)
	}
	source := "tmp"
	err = os.MkdirAll(source, 0755)
	if err != nil {
		panic(err)
	}
	err = getter.Get(source, "https://cli-tool.s3.amazonaws.com/"+language+"/"+fileName+".zip")
	if err != nil {
		panic(err)
	}

	err = ParsePaths(source, tp)
	if err != nil {
		panic(err)
	}

	paths, err := walkFiles(source)
	if err != nil {
		panic(err)
	}

	for _, file := range paths {
		ParseFile(file, tp)
	}

	err = gorecurcopy.CopyDirectory(source, destination)
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll(source)
	if err != nil {
		panic(err)
	}
}
