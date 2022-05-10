package file

import (
	"cleaning/pkg/logging"
	"io/ioutil"
	"os"
)

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)

	return f, err == nil || os.IsExist(err)
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
