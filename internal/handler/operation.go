package handler

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"syscall"
	"time"

	"cleaning/configs"
	"cleaning/pkg/file"
	"cleaning/pkg/logging"

	"github.com/pkg/errors"
)

type Handler interface {
	Runner() error
}

type handler struct {
	fileInfos *FileInfos
}

func GetAllFileName(filePath string) ([]string, error) {

	result := []string{}
	files, err := ioutil.ReadDir(filePath)

	if err != nil {
		logging.Error(err)

		return result, err
	}

	for _, file := range files {
		fullPath := filePath + "/" + file.Name()

		if file.IsDir() {
			next, err := GetAllFileName(fullPath)

			if err != nil {
				logging.Error(err)
				return result, err
			}

			result = append(result, next...)
		} else {
			result = append(result, fullPath)
		}
	}
	return result, nil
}

func (h *handler) removeFile() error {
	logging.Info("message", "Remove", h.fileInfos.Name)
	ok := false
	switch h.fileInfos.Suffix {
	case ".gz":
		ok = true
	default:
		old := h.fileInfos.Path + ".gz"
		_, ok = file.IsExists(old)

		if !ok {
			logging.Info("Info", h.fileInfos.Path, "not compress")
			// todo
		}
	}

	if ok {
		err := os.Remove(h.fileInfos.Path)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (h *handler) gzipFile() error {
	logging.Info("message", "Gzip", h.fileInfos.Name)

	finFile := fmt.Sprintf("%s.gz", h.fileInfos.Path)

	cFile, err := os.Create(finFile)
	if err != nil {
		return errors.WithStack(err)
	}

	defer cFile.Close()

	file, err := os.Open(h.fileInfos.Path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	read := bufio.NewReader(file)

	data, err := ioutil.ReadAll(read)
	if err != nil {
		return errors.WithStack(err)
	}

	zw := gzip.NewWriter(cFile)
	zw.Write(data)

	zw.Flush()
	zw.Close()
	if err := zw.Close(); err != nil {
		return nil
	}
	return nil
}

func (h *handler) Runner() error {

	suffix := h.fileInfos.Suffix
	mtime := h.fileInfos.ModificationTime

	ntime := time.Now().Unix()
	switch suffix {
	case ".gz":
		if ntime-mtime > int64(configs.Get().Logger.Zipdays*24*60*60) {
			if err := h.removeFile(); err != nil {
				return errors.WithStack(err)
			}
		}
	default:
		if ntime-mtime > int64(configs.Get().Logger.Normaldays*24*60*60) {
			if err := h.gzipFile(); err != nil {
				return errors.WithStack(err)
			}
			if err := h.removeFile(); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func NewFileHandler(filename, filepath string) (Handler, error) {
	filestat, err := os.Stat(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	stat := filestat.Sys().(*syscall.Stat_t)

	modificationTime := int64(stat.Mtimespec.Sec)
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
