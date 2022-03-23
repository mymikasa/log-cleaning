package main

import (
	"cleaning/configs"
	"cleaning/internal/handler"
	"cleaning/pkg/logging"
	"path"
)

func main() {
	filepath := configs.Get().Logger.LogPath

	filenames, err := handler.GetAllFileName(filepath)

	if err != nil {
		logging.Panic(err)
	}

	for _, fpath := range filenames {
		name := path.Base(fpath)
		h, err := handler.NewFileHandler(name, fpath)
		if err != nil {
			logging.Panic(err)
		}
		h.Runner()
	}
	// fmt.Println(configs.Get().Logger.LogPath)
}
