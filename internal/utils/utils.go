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

