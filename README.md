# MediaPlex

An interface for your local media such as movies, TV shows, etc.

## Getting Started

### Prerequisites

- Go 1.25.0+

### Usage

```bash
go run mediaplex.go -path "C:\Videos"
```

Scans a directory for supported media formats (`.mp4`, `.mkv`).

> **Note:** Currently scans one directory level. Recursive search coming soon.

## Supported Formats

- MP4 (`.mp4`)
- Matroska (`.mkv`)

## License

[MIT License](./LICENSE)

