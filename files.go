package main

import (
	"os"
	"path/filepath"
	"strings"
)

func GoFiles(root string) ([]string, error) {
	var data []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		base := filepath.Base(path)
		for _, sp := range SkipPaths {
			// if the name of the folder has a prefix listed in SkipPaths
			// then we should skip the directory.
			// e.g. node_modules, testdata, _foo, .git
			if strings.HasPrefix(base, sp) {
				return filepath.SkipDir
			}
		}

		// skip non-go files
		if filepath.Ext(path) != ".go" {
			return nil
		}

		data = append(data, path)

		return nil
	})

	return data, err
}
