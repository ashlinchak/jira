package lib

import (
	"os"
	"path/filepath"
)

func FilePath(file string) string {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	return exPath + string(os.PathSeparator) + file
}
