package scanner

import (
	"mediaplex/scanner/cleaner"
	"mediaplex/scanner/crawler"
)

func Scan(path string) (error, []string) {
	err, results := crawler.Crawl(path)

	if err != nil {
		return err, []string{}
	}

	return nil, cleaner.Clean(results)
}
