package flag_parsers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// findFirstExistingParent walks up the path to find the closest existing directory.
func findFirstExistingParent(path string) (string, error) {
	dir := filepath.Dir(path)
	for {

		// If there's no error while getting the directory info, it exists, so return it.
		if _, err := os.Stat(dir); err == nil {
			return dir, nil

			// If the dir exists but there was an error getting its info, return an error
		} else if !os.IsNotExist(err) {
			return "", err
		}

		// The directory wasn't found, so we look for the directory of the directory.
		parent := filepath.Dir(dir)

		// If the parent is the same as the directory, then we've reached the root of the file system, so we can end the loop.
		if parent == dir {
			break // reached root
		}

		// The new directory will be the parent of the current directory.
		dir = parent
	}
	return "", errors.New("no existing parent directory found")
}

func ParseOutput(outputFilePath string) (string, error) {

	absPath, err := filepath.Abs(outputFilePath)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(absPath)

	// Check if file exists.
	if err == nil {

		// File exists — check if writable
		file, err := os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND, info.Mode())

		if err != nil {
			return "", errors.New("output file exists but is not writable")
		}

		// Close the file.
		file.Close()
		return absPath, nil
	}

	// Check if the file exists but there was some error based on the error returned from getting the info of the file.
	if !os.IsNotExist(err) {
		return "", err
	}

	// File doesn't exist — find nearest existing parent directory
	existingDir, err := findFirstExistingParent(absPath)

	if err != nil {
		return "", err
	}

	// Check if we can write in that directory
	testFile := filepath.Join(existingDir, ".perm_check_tmp")

	file, err := os.Create(testFile)

	if err != nil {
		return "", fmt.Errorf("no permission to create files in parent directory: %v", existingDir)
	}

	file.Close()

	os.Remove(testFile)

	return outputFilePath, nil

}
