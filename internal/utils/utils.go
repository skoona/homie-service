package utils

import (
	"fmt"
	"os"
	"strings"
)



// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func LocateFile(dataFile string) string {
	var foundFile string

	if fileExists(dataFile) {
		foundFile = dataFile
	} else {
		if strings.HasPrefix(dataFile, "../") {
			foundFile = strings.ReplaceAll(dataFile, "../.", "")
			if !fileExists(foundFile) {
				fmt.Printf("error FILE NOT FOUND: %s or %s", foundFile, dataFile)
				return ""
			}
		}
	}
	return foundFile
}

