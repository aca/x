package main

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

type spaHandler struct {
	files map[string][]byte
}

// var staticFiles embed.FS
// subFS, _ := fs.Sub(staticFiles, "dist")
// r.PathPrefix("/").Handler(New(subFS))
func New(subbedFS fs.FS) *spaHandler {
	h := &spaHandler{}
	h.files = map[string][]byte{}

	fs.WalkDir(subbedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, _ := subbedFS.Open(path)
		h.files["/"+path], _ = io.ReadAll(f)
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

	if b, ok := h.files[path]; ok {
		_, err := w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		_, err := w.Write(h.files["/index.html"])
		if err != nil {
			log.Println(err)
		}
		return
	}
}
