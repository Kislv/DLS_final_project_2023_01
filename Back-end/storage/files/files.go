package files

import (
	"errors"
	"path/filepath"

	"read-adviser-bot/lib/e"
	"read-adviser-bot/utils/log"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

var (
	ErrBadFileExtension = errors.New("bad file extension")
)

//Saving file
func IsValidExtension (filePath string) (error) {
	validExtensions := map[string]struct{}{
		".jpg": {},
		".jpeg": {},
		".png": {},
	}
	ext := filepath.Ext(filePath)
	log.Info("ext:" + ext +"\n")

	if _, is := validExtensions[ext]; !is {
		return e.Wrap("Extension = " + ext, ErrBadFileExtension)
	}
	return nil
}
