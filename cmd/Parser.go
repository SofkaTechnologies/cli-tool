package cmd

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type settings struct {
	Project string
	Group   string
}

func ParseFile(file string, projectName string, group string) {
	paths := []string{
		file,
	}
	tp := settings{projectName, group}
	var t *template.Template
	t = template.Must(template.New(path.Base(paths[0])).ParseFiles(paths...))
	f, ex := os.Create(paths[0])
	if ex != nil {
		panic(ex)
	}
	err := t.Execute(f, tp)
	if err != nil {
		panic(err)
	}
	_ = f.Close()
}

func revealPaths(folder string, projectName string, group string) string {
	tp := settings{projectName, strings.ReplaceAll(group, ".", string(os.PathSeparator))}
	var t *template.Template
	t, err := template.New(folder).Parse(folder)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, tp)

	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func ParsePaths(root string, projectName string, group string) error {
	//var files []string
	var oldPaths []string
	err := filepath.Walk(root, func(currentPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			newPath := revealPaths(currentPath, projectName, group)
			err = os.MkdirAll(newPath, 0700)
			if err != nil {
				panic(err)
			}
			if newPath != currentPath {
				oldPaths = append(oldPaths, currentPath)
			}

		} else {
			newPath := revealPaths(currentPath, projectName, group)
			err = os.Rename(currentPath, newPath)
			if err != nil {
				panic(err)
			}
			//files = append(files, newPath)
		}
		return nil
	})
	for _, path := range oldPaths {
		err = os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}
	return err
}
