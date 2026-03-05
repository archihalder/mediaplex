package crawler

import (
	"os"
	"path/filepath"
	"strings"
)

var supportedFormats = map[string]bool{
	".mp4": true,
	".mkv": true,
}

func Crawl(path string) (error, []string) {
	var results []string
	entries, err := os.ReadDir(path)

	if err != nil {
		return err, results
	}

	for _, entry := range entries {
		if entry.IsDir() {
			newPath := filepath.Join(path, entry.Name())

			_, err := os.ReadDir(newPath)

			if err != nil {
				return err, results
			}

			_, subResults := Crawl(newPath)
			results = append(results, subResults...)
		} else {
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if supportedFormats[ext] {
				results = append(results, entry.Name())
			}
		}
	}

	return nil, results
}
