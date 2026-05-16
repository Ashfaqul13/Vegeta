package cmd

import (
	"errors"
	"os"
	"path/filepath"
)

// ✅ read file
func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ✅ validate extension
func validateExt(path string, allowed ...string) error {
	ext := filepath.Ext(path)

	for _, a := range allowed {
		if ext == a {
			return nil
		}
	}

	return errors.New("unsupported file format")
}

// ✅ auto generate output file name
func generateOutput(input string, newExt string) string {
	base := input[:len(input)-len(filepath.Ext(input))]
	return base + newExt
}
