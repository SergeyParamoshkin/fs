package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

/*
Первый подход к тестированию кода файловой системы
это создание необходимых структур файлов / папок,
необходимых для этого теста во время выполнения.
*/
func BenchmarkGoFilesJIT(b *testing.B) {
	for i := 0; i < b.N; i++ {

		dir, err := os.MkdirTemp("", "demodir")
		if err != nil {
			b.Fatal(err)
		}

		names := []string{"foo.go", "web/routes.go"}

		for _, s := range SkipPaths {
			// ex: ./.git/git.go
			// ex: ./node_modules/node_modules.go
			names = append(names, filepath.Join(s, s+".go"))
		}

		for _, f := range names {
			if err := os.MkdirAll(filepath.Join(dir, filepath.Dir(f)), 0755); err != nil {
				b.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, f), nil, 0666); err != nil {
				b.Fatal(err)
			}
		}

		list, err := GoFiles(dir)

		if err != nil {
			b.Fatal(err)
		}

		lexp := 2
		lact := len(list)
		if lact != lexp {
			b.Fatalf("expected list to have %d files, but got %d", lexp, lact)
		}

		sort.Strings(list)

		exp := []string{"foo.go", "web/routes.go"}
		for i, a := range list {
			e := exp[i]
			if !strings.HasSuffix(a, e) {
				b.Fatalf("expected %q to match expected %q", list, exp)
			}
		}

	}
}

func BenchmarkGoFilesFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		files := MockFS{
			// ./foo.go
			NewFile("foo.go"),
			// ./web/routes.go
			NewDir("web", NewFile("routes.go")),
		}

		for _, s := range SkipPaths {
			// ex: ./.git/git.go
			// ex: ./node_modules/node_modules.go
			files = append(files, NewDir(s, NewFile(s+".go")))
		}

		mfs := MockFS{
			// ./
			NewDir(".", files...),
		}

		list, err := GoFilesFS("/", mfs)

		if err != nil {
			b.Fatal(err)
		}

		lexp := 2
		lact := len(list)
		if lact != lexp {
			b.Fatalf("expected list to have %d files, but got %d", lexp, lact)
		}

		sort.Strings(list)

		exp := []string{"foo.go", "web/routes.go"}
		for i, a := range list {
			e := exp[i]
			if e != a {
				b.Fatalf("expected %q to match expected %q", list, exp)
			}
		}

	}
}

func BenchmarkGoFilesExistingFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {

		list, err := GoFiles("./testdata/case")

		if err != nil {
			b.Fatal(err)
		}

		lexp := 2
		lact := len(list)
		if lact != lexp {
			b.Fatalf("expected list to have %d files, but got %d", lexp, lact)
		}

		sort.Strings(list)

		exp := []string{"foo.go", "testdata/routes.go"}
		for i, a := range list {
			e := exp[i]
			if !strings.HasSuffix(a, e) {
				b.Fatalf("expected %q to match expected %q", list, exp)
			}
		}

	}
}
