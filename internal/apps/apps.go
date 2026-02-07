package apps

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	// Cache detected apps to avoid scanning repeatedly if we used this in a daemon
	// For CLI, it's less critical but good practice.
	cache      map[string]string
	cacheOnce  sync.Once
	searchDirs = []string{
		"/Applications",
		"/System/Applications",
		"/Users/" + os.Getenv("USER") + "/Applications",
	}
)

// Find looks for an app with the given name (case-insensitive)
func Find(name string) (string, bool) {
	cacheOnce.Do(scan)
	path, ok := cache[strings.ToLower(name)]
	return path, ok
}

// List returns all detected apps
func List() map[string]string {
	cacheOnce.Do(scan)
	return cache
}

func scan() {
	cache = make(map[string]string)
	
	for _, dir := range searchDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".app") {
				// Key is lowercase, spaces removed, extension removed
				// e.g. "Google Chrome.app" -> "googlechrome" and "google chrome"
				// We want flexible matching.
				
				baseName := strings.TrimSuffix(entry.Name(), ".app")
				fullPath := filepath.Join(dir, entry.Name())
				
				// Add exact name (lowercase)
				cache[strings.ToLower(baseName)] = fullPath
				
				// Add name without spaces (lowercase)
				cache[strings.ReplaceAll(strings.ToLower(baseName), " ", "")] = fullPath
			}
		}
	}
}
