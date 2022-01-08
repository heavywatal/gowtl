package pathlib

import (
	"log"
	"os"

	"path/filepath"
)

func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func WithSuffix(path string, ext string) string {
	orig_ext := filepath.Ext(path)
	return path[0:len(path)-len(orig_ext)] + ext
}
