package file

import "os"

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)

	return f, err == nil || os.IsExist(err)
}
