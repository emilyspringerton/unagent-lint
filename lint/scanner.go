package lint

import (
	"io/fs"
	"os"
	"path/filepath"
)

var allowedExtensions = map[string]bool{
	".md":   true,
	".txt":  true,
	".yml":  true,
	".yaml": true,
	".json": true,
}

func ScanFiles(path string) []string {
	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	if !info.IsDir() {
		if allowedExtensions[filepath.Ext(path)] {
			return []string{path}
		}
		return nil
	}

	files := []string{}
	_ = filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if allowedExtensions[filepath.Ext(p)] {
			files = append(files, p)
		}
		return nil
	})
	return files
}

func LoadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}
