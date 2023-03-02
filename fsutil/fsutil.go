package fsutil

import (
	"os"
)

func IsDir(name string) (bool, error) {
	stat, err := os.Lstat(name)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

func IsSymLink(name string) (bool, error) {
	stat, err := os.Lstat(name)
	if err != nil {
		return false, err
	}

	if stat.Mode()&os.ModeSymlink != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func IsRegularFile(name string) (bool, error) {
	stat, err := os.Lstat(name)
	if err != nil {
		return false, err
	}

	return stat.Mode().IsRegular(), nil
}
