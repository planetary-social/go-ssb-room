package web

import (
	"errors"
	"io/fs"
	"os"
)

type appFS struct {
	app    fs.FS
	assets fs.FS
}

func newAppFS(app fs.FS, assets fs.FS) (fs.FS, error) {
	return &appFS{
		app:    app,
		assets: assets,
	}, nil
}

func (a appFS) Open(name string) (fs.File, error) {
	v, err := a.app.Open(name)
	if errors.Is(err, os.ErrNotExist) {
		return a.assets.Open(name)
	}
	return v, err
}
