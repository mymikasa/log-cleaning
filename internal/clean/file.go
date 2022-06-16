package clean

import (
	"bufio"
	"cleaning/pkg/file"
	"cleaning/pkg/logging"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Handler interface {
	Clean(zipdays, normaldays int) error
}

type handler struct {
	fileInfos *FileInfos
}

func (h *handler) removeFile() error {
	ok := false
	switch h.fileInfos.Suffix {
	case ".gz":
		ok = true
	default:
		old := h.fileInfos.Path + ".gz"
		_, ok = file.IsExists(old)

		if !ok {
			logging.Info("Info", h.fileInfos.Path, "not compress")
		}
	}

	if ok {
		err := os.Remove(h.fileInfos.Path)
		if err != nil {
			return errors.WithStack(err)
		}
		logging.Info("cleaning", "remove", h.fileInfos.Name)
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

	logging.Info("cleaning", "gzip", h.fileInfos.Name)
	return nil
}

func (h *handler) Clean(zipdays, normaldays int) error {

	suffix := h.fileInfos.Suffix
	mtime := h.fileInfos.ModificationTime

	ntime := time.Now().Unix()
	switch suffix {
	case ".gz":
		if ntime-mtime > int64(zipdays*24*60*60) {
			if err := h.removeFile(); err != nil {
				return errors.WithStack(err)
			}
		}
	case ".log", ".txt", ".out":
		if ntime-mtime > int64(normaldays*24*60*60) {
			if err := h.gzipFile(); err != nil {
				return errors.WithStack(err)
			}
			if err := h.removeFile(); err != nil {
				return errors.WithStack(err)
			}
		}
	default:
		fmt.Println("Unexpect suffix")
	}
	return nil
}
