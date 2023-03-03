package findroot

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrReachedEnd = errors.New("findroot: reached end")

func FindRoot(patterns ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for i := 0; i < 100; i++ {
		for _, pattern := range patterns {
			_, err := os.Stat(filepath.Join(wd, pattern))
			if err != nil && !os.IsNotExist(err) {
				return "", err
			} else {
				return wd, nil
			}
		}

		nwd := filepath.Dir(wd)
		if nwd == wd {
			return "", ErrReachedEnd
		}
		wd = nwd
	}

	return wd, nil
}
