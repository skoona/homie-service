package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

func LocateFile(dFile string) string {
	var foundFile string
	dataFile := filepath.FromSlash(dFile)

	if fileExists(dataFile) {
		foundFile = dataFile
	} else {
		if strings.HasPrefix(dataFile, ".."+string(os.PathSeparator)+"") {
			foundFile = strings.ReplaceAll(dataFile, ".."+string(os.PathSeparator)+".", "")
			if !fileExists(foundFile) {
				fmt.Printf("error FILE NOT FOUND: %s or %s", foundFile, dataFile)
				return ""
			}
		}
	}
	return foundFile
}


// RemoveIndexFromSlice()
// -- No good way to handle various types of slices
//    since reflection does not support custom types
func RemoveIndexFromSlice(slice interface{}, index int) interface{} {
	var idx  int

	switch a := slice.(type) {
	case []int:
		if index > len(a) {
			idx = len(a) - 1
		} else if index < 0 {
			idx = 0
		} else {
			idx = index
		}
		return append(a[:idx], a[idx+1:]...)

	}

	return  slice
}

