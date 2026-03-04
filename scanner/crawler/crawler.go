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

			files, err := os.ReadDir(newPath)

			if err != nil {
				return err
			}

			for _, file := range files {
				ext := strings.ToLower(filepath.Ext(file.Name()))
				if supportedFormats[ext] {
					fmt.Println(file.Name())
				}
			}
		}
	}

	return nil
}
