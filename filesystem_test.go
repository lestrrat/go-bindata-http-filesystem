package bindata

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fileinfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func (fi *fileinfo) Name() string {
	return fi.name
}

func (fi *fileinfo) Size() int64 {
	return fi.size
}

func (fi *fileinfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *fileinfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *fileinfo) IsDir() bool {
	return fi.isDir
}

func (fi *fileinfo) Sys() interface{} {
	return fi.sys
}

func TestFilesystem(t *testing.T) {
	fs := &FileSystem{}
	fs.AssetLoadFunc = func(name string) ([]byte, error) {
		return []byte("Hello, World!"), nil
	}
	fs.AssetInfoLoadFunc = func(name string) (os.FileInfo, error) {
		return &fileinfo{
			name:    "helloworld",
			size:    13,
			mode:    os.FileMode(0644),
			modTime: time.Now(),
			isDir:   false,
			sys:     nil,
		}, nil
	}

	_, err := fs.Open("helloworld")
	if !assert.NoError(t, err) {
		return
	}
}