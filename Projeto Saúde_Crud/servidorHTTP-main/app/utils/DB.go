package utils

import (
	"database/sql"
	"os"
	"path/filepath"
)

var DB *sql.DB

// GetTemplatePath returns the correct path to a template file
func GetTemplatePath(relPath string) string {
	// Try from current working directory
	if _, err := os.Stat(relPath); err == nil {
		return relPath
	}

	// Try with "app/" prefix
	appPath := filepath.Join("app", relPath)
	if _, err := os.Stat(appPath); err == nil {
		return appPath
	}

	// Default to the relative path (will fail with better error message)
	return relPath
}
