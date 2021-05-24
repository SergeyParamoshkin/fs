package main

import (
	"io"
	"io/fs"
	"os"
	"sort"
	"time"
)

type MockFile struct {
	FS      MockFS
	isDir   bool
	modTime time.Time
	mode    fs.FileMode
	name    string
	size    int64
	sys     interface{}
}

type MockFS []*MockFile

func (mfs MockFS) Open(name string) (fs.File, error) {
	for _, f := range mfs {
		if f.Name() == name {
			return f, nil
		}
	}

	if len(mfs) > 0 {
		return mfs[0].FS.Open(name)
	}

	return nil, &fs.PathError{
		Op:   "read",
		Path: name,
		Err:  os.ErrNotExist,
	}
}

func (m *MockFile) Name() string {
	return m.name
}

func (m *MockFile) IsDir() bool {
	return m.isDir
}

func (mf *MockFile) Info() (fs.FileInfo, error) {
	return mf.Stat()
}

func (mf *MockFile) Stat() (fs.FileInfo, error) {
	return mf, nil
}

func (m *MockFile) Size() int64 {
	return m.size
}

func (m *MockFile) Mode() os.FileMode {
	return m.mode
}

func (m *MockFile) ModTime() time.Time {
	return m.modTime
}

func (m *MockFile) Sys() interface{} {
	return m.sys
}

func (m *MockFile) Type() fs.FileMode {
	return m.Mode().Type()
}

func (mf *MockFile) Read(p []byte) (int, error) {
	panic("not implemented")
}

func (mf *MockFile) Close() error {
	return nil
}

func (m *MockFile) ReadDir(n int) ([]fs.DirEntry, error) {
	if !m.IsDir() {
		return nil, os.ErrNotExist
	}

	if m.FS == nil {
		return nil, nil
	}
	return m.FS.ReadDir(n)
}

func (mfs MockFS) ReadDir(n int) ([]fs.DirEntry, error) {
	list := make([]fs.DirEntry, 0, len(mfs))

	for _, v := range mfs {
		list = append(list, v)
	}

	sort.Slice(list, func(a, b int) bool {
		return list[a].Name() > list[b].Name()
	})

	if n < 0 {
		return list, nil
	}

	if n > len(list) {
		return list, io.EOF
	}
	return list[:n], io.EOF
}

func NewFile(name string) *MockFile {
	return &MockFile{
		name: name,
	}
}

func NewDir(name string, files ...*MockFile) *MockFile {
	return &MockFile{
		FS:    files,
		isDir: true,
		name:  name,
	}
}
