// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

//go:build !dev
// +build !dev

package web

import (
	"embed"
	"fmt"
	"github.com/friendsofgo/errors"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// Production can be used to determain different aspects at compile time (like hot template reloading)
const Production = true

var (
	Assets    http.FileSystem
	Templates http.FileSystem
)

// correct the paths by stripping their prefixes
func init() {
	var err error

	prefixedAssets, err := newAppFS(embedAssets, embedApp)
	if err != nil {
		log.Fatal(err)
	}
	Assets = http.FS(prefixedAssets)

	prefixedTemplates, err := fs.Sub(embedTemplates, "templates")
	if err != nil {
		log.Fatal(err)
	}
	Templates = http.FS(prefixedTemplates)
}

//go:embed templates/*
var embedTemplates embed.FS

//go:embed assets/*
var embedAssets embed.FS

//go:embed app/*
var embedApp embed.FS

type appFS struct {
	app    fs.FS
	assets fs.FS
}

func newAppFS(assets embed.FS, app embed.FS) (fs.FS, error) {
	prefixedApp, err := fs.Sub(app, "app")
	if err != nil {
		return nil, err
	}

	prefixedAssets, err := fs.Sub(assets, "assets")
	if err != nil {
		return nil, err
	}

	return &appFS{
		app:    prefixedApp,
		assets: prefixedAssets,
	}, nil
}

func (a appFS) Open(name string) (fs.File, error) {
	fmt.Println("open", name)
	v, err := a.app.Open(name)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("returning from assets", name)
		return a.assets.Open(name)
	}
	fmt.Println("returning from app", name)
	return v, err
}
