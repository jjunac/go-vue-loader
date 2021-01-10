package internal

import (
	"path/filepath"
	"strings")


const (
	toLower = 'a' - 'A'
)

func PascalCaseToKebabCase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if 'Z' >= r && r >= 'A' {
			if i != 0 {
				builder.WriteRune('-')
			}
			builder.WriteRune(r + toLower)
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func FilenameWithoutExtension(path string) string {
	filename := filepath.Base(path)
	return filename[:len(filename) - len(filepath.Ext(filename))]
}
