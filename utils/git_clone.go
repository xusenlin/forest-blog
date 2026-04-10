package utils

import (
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

func Clone(repoURL, dest string) error {
	if strings.HasSuffix(dest, "/") {
		dest = strings.TrimSuffix(dest, "/")
	}

	dir := dest[:strings.LastIndex(dest, "/")]

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return err
}

func Pull(repoDir string) error {
	r, err := git.PlainOpen(repoDir)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{
		Progress: os.Stdout,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

func CloneOrPull(repoURL, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		return Clone(repoURL, dest)
	}
	return Pull(dest)
}
