package mock

import (
	"os"
	"time"
)

type FakeFileInfo struct {
	Dir bool
}

func (_ FakeFileInfo) Name() string {
	return ""
}

func (_ FakeFileInfo) Size() int64 {
	return 0
}

func (_ FakeFileInfo) Mode() os.FileMode {
	return 0
}

func (_ FakeFileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

func (ffi FakeFileInfo) IsDir() bool {
	return ffi.Dir
}

func (_ FakeFileInfo) Sys() any {
	return nil
}
