package utils

import (
	uuid "github.com/satori/go.uuid"
	"path/filepath"
)

func GenerateUniqueFileId(filename string) string {
	ext := filepath.Ext(filename)
	return uuid.NewV4().String() + ext
}
