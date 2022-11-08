package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/aca/x/spahandler"
	"github.com/go-chi/chi/v5"
)

//go:embed dist
var staticFiles embed.FS

func main() {
	subFS, _ := fs.Sub(staticFiles, "dist")
    mux := chi.NewRouter()
    mux.Handle("/*", spahandler.New(subFS))
    http.ListenAndServe(":8080", mux)
}
