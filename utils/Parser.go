package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ProjectSettings struct {
	Project string
	Group   string
	Entity  string
}

var funcMap = template.FuncMap{
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
}

func ParseFile(file string, tp ProjectSettings) {
	paths := []string{
		file,
	}

	if strings.Contains(paths[0], ".png") {
		return
	}

	var t *template.Template

	t = template.Must(template.New(filepath.Base(paths[0])).Funcs(funcMap).ParseFiles(paths...))
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

func revealPaths(folder string, tp ProjectSettings) string {
	tp.Group = strings.ReplaceAll(tp.Group, ".", string(os.PathSeparator))
	var t *template.Template
	t, err := template.New(folder).Funcs(funcMap).Parse(folder)
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

func ParsePaths(root string, tp ProjectSettings) error {
	//var files []string
	var oldPaths []string
	err := filepath.Walk(root, func(currentPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			newPath := revealPaths(currentPath, tp)
			err = os.MkdirAll(newPath, 0700)
			if err != nil {
				panic(err)
			}
			if newPath != currentPath {
				oldPaths = append(oldPaths, currentPath)
			}

		} else {
			newPath := revealPaths(currentPath, tp)
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
