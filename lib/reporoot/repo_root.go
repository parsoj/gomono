package reporoot

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

var ErrNotInGitRepo = errors.New("Not in a Git Repo")

func FindParentRepoForCwd(cwd_path string) (string, error) {

	//get absolute path
	abs, err := filepath.Abs(cwd_path)
	if err != nil {
		return "", err
	}

	return findRepo(abs)
}

// recursive helper func
func findRepo(root string) (string, error) {

	//Check if root is a real directory
	if !directoryExists(root) {
		return "", os.ErrNotExist
	}

	//Check if root contains a git repo
	if pathIsGitRepo(root) {
		return root, nil
	}

	//Check if root is the filesystem root
	if root == "/" {
		return "", ErrNotInGitRepo
	}

	//Recurse on parent
	return findRepo(path.Base(root))
}

func pathIsGitRepo(path string) bool {
	if _, err := git.PlainOpen(path); err == nil {
		return true
	} else {
		return false
	}

}

func directoryExists(path string) bool {
	folderInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return folderInfo.IsDir()
}
