package spahandler

import (
	"io/fs"
	"net/http"
)

type spaHandler struct {
	files       map[string]struct{}
	fileHandler http.Handler
}

func New(subbedFS fs.FS) *spaHandler {
	h := &spaHandler{}
	h.files = make(map[string]struct{})
	h.fileHandler = http.FileServer(http.FS(subbedFS))

	fs.WalkDir(subbedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		h.files["/"+path] = struct{}{}
		return nil
	})

	return h
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.files[r.URL.Path]; !ok {
        r.URL.Path = "/"
	} 
	h.fileHandler.ServeHTTP(w, r)
}
