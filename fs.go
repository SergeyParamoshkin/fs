package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GoFilesFS(root string, sys fs.FS) ([]string, error) {
	var data []string

	err := fs.WalkDir(sys, ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		for _, sp := range SkipPaths {
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
