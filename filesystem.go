package bindata

import (
	"bytes"
	"net/http"
	"os"
	"syscall"
)

type file struct {
	*bytes.Reader
	fi os.FileInfo
}

type AssetLoadFunc func(string) ([]byte, error)
type AssetInfoLoadFunc func(string) (os.FileInfo, error)

type FileSystem struct {
	AssetLoadFunc     AssetLoadFunc
	AssetInfoLoadFunc AssetInfoLoadFunc
}

func (fs *FileSystem) Open(name string) (http.File, error) {
	asset, err := fs.AssetLoadFunc(name)
	if err != nil {
		return nil, err
	}

	fi, err := fs.AssetInfoLoadFunc(name)
	if err != nil {
		return nil, err
	}

	f := &file{
		Reader: bytes.NewReader(asset),
		fi:     fi,
	}

	return f, nil
}

func (f *file) Close() error {
	return nil
}

func (f *file) Readdir(_ int) ([]os.FileInfo, error) {
	return nil, &os.SyscallError{Syscall: "readdirent", Err: syscall.EINVAL}
}

func (f *file) Stat() (os.FileInfo, error) {
	return f.fi, nil
}