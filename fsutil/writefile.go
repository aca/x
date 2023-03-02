package fsutil

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// Based on "tailscale.com/atomicfile"
//
// WriteFile writes data to filename+some suffix, then renames it
// into filename. The perm argument is ignored on Windows.
func WriteFile(filename string, data []byte) (err error) {
	f, err := os.CreateTemp(filepath.Dir(filename), filepath.Base(filename)+".tmp")
	if err != nil {
		return err
	}
	tmpName := f.Name()
	defer func() {
		if err != nil {
			f.Close()
			os.Remove(tmpName)
		}
	}()
	if _, err := f.Write(data); err != nil {
		return err
	}
	if runtime.GOOS != "windows" {
		if err := f.Chmod(0o777); err != nil {
			return err
		}
	}
	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, filename)
}

func WriteFileFromReader(filename string, reader io.Reader) (err error) {
	f, err := os.CreateTemp(filepath.Dir(filename), filepath.Base(filename)+".tmp")
	if err != nil {
		return err
	}
	tmpName := f.Name()
	defer func() {
		if err != nil {
			f.Close()
			os.Remove(tmpName)
		}
	}()
	if _, err := io.Copy(f, reader); err != nil {
		return err
	}
	if runtime.GOOS != "windows" {
		if err := f.Chmod(0o777); err != nil {
			return err
		}
	}
	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, filename)
}
