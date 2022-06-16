package clean

import (
	"os"
	"path"
	"syscall"

	"github.com/pkg/errors"
)

func New(option ...Option) *FileInfos {
	o := &FileInfos{}
	for _, opt := range option {
		opt(o)
	}
	return o
}

func NewFileHandler(filename, filepath string) (Handler, error) {
	filestat, err := os.Stat(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	modTime := filestat.ModTime()
	stat := filestat.Sys().(*syscall.Stat_t)

	modificationTime := modTime.Unix()
	createtime := int64(stat.Ctimespec.Sec)
	visitTime := int64(stat.Atimespec.Sec)
	suffix := path.Ext(filename)

	f := New(
		WithCreatetime(createtime),
		WithModificationTime(modificationTime),
		WithVisitTime(visitTime),
		WithSize(int64(stat.Size)/1024),
		WithFileName(filename),
		WithPath(filepath),
		WithSuffix(suffix))

	h := &handler{
		fileInfos: f,
	}

	return h, nil
}
