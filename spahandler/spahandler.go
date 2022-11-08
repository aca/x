package spahandler

import (
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
)

type file struct {
	contents []byte
	ctype    string
}

type spaHandler struct {
	files map[string]*file
}

func New(subbedFS fs.FS) *spaHandler {
	h := &spaHandler{}
	h.files = make(map[string]*file)

	fs.WalkDir(subbedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, _ := subbedFS.Open(path)
		b, _ := io.ReadAll(f)
		ctype := mime.TypeByExtension(filepath.Ext(path))
		if ctype == "" {
			ctype = "application/octet-stream"
		}
		h.files["/"+path] = &file{
			contents: b,
			ctype:    ctype,
		}
		f.Close()
		return nil
	})

	return h
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var f *file
	var ok bool
	if f, ok = h.files[path]; !ok {
		f = h.files["/index.html"]
	}

	w.Header().Set("Content-Type", f.ctype)
	w.Write(f.contents)
}
