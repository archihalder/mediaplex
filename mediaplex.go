package main

import (
	"flag"
	"fmt"
	"mediaplex/scanner/crawler"
	"os"
)

func main() {
	// example usage: go run mediaplex.go -path "C:\Videos"
	path := flag.String("path", "", "Path to scan for media files")
	flag.Parse()

	// no file path provided, exit with error
	if *path == "" {
		fmt.Fprintln(os.Stderr, "Please provide a file path to scan using the -path flag")
		os.Exit(1)
	}

	if err := crawler.Crawl(*path); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
