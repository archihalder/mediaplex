package main

import (
	"flag"
	"fmt"
	"mediaplex/scanner"
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

	err, results := scanner.Scan(*path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning path: %v\n", err)
		os.Exit(1)
	} else {
		for _, movie := range results {
			fmt.Println(movie)
		}
	}

}
