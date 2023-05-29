package csvdb

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jszwec/csvutil"
)

// DB is a database backed by a JSON file.
type DB[E ~[]T, T any] struct {
	// Data is the contents of the database.
	Data *E

	path string
}

// Open opens the database at path, creating it with a zero value if
// necessary.
func Open[T any](path string) (*DB[[]T, T], error) {
	bs, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return &DB[[]T, T]{
			Data: new([]T),
			path: path,
		}, nil
	} else if err != nil {
		return nil, err
	}

	var val []T
	if err := csvutil.Unmarshal(bs, &val); err != nil {
		return nil, err
	}

	return &DB[[]T, T]{
		Data: &val,
		path: path,
	}, nil
}

// Save writes db.Data back to disk.
func (db *DB[E, T]) Save() error {
	bs, err := csvutil.Marshal(db.Data)
	if err != nil {
		return err
	}

	return WriteFile(db.path, bs, 0o600)
}

// WriteFile writes data to filename+some suffix, then renames it
// into filename.
func WriteFile(filename string, data []byte, perm os.FileMode) (err error) {
	f, err := ioutil.TempFile(filepath.Dir(filename), filepath.Base(filename)+".tmp")
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
		if err := f.Chmod(perm); err != nil {
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
