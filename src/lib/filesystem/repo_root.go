package filesystem

import (
	"fmt"
	"os"
	"path"
)

func FindRepoRootForCwd(path string) (string, error) {

	//----------------------------------------------------------------------
	// Input cleaning
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Error Opening file")
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("Error getting file stats")
	}

	if !fileInfo.IsDir() {
		return "", fmt.Errorf("Input filepath doesn't point to a valid directory")
	}

	//----------------------------------------------------------------------
	// convert to abs path

	//----------------------------------------------------------------------
	// check if Cwd is root (or homedir?)
	if path == "/" { // TODO is this enough?
		return "", fmt.Errorf("CWD does not exist in any git repo")
	}

	//----------------------------------------------------------------------
	// Check if Cwd is a git repo

	//----------------------------------------------------------------------
	// Recurse on parent

}
