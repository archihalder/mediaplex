package crawler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var supportedFormats = map[string]bool{
	".mp4": true,
	".mkv": true,
}

func Crawl(path string) error {
	entries, err := os.ReadDir(path)

	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			newPath := filepath.Join(path, entry.Name())

			_, err := os.ReadDir(newPath)

			if err != nil {
				return err
			}

			Crawl(newPath)
		} else {
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if supportedFormats[ext] {
				fmt.Println(entry.Name())
			}
		}
	}

	return nil
}
