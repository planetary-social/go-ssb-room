// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

//go:build dev
// +build dev

package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.mindeco.de/goutils"
)

const Production = false

// absolute path of where this package is located
var pkgDir = goutils.MustLocatePackage("github.com/ssbc/go-ssb-room/v2/web")

var Templates = http.Dir(filepath.Join(pkgDir, "templates"))

var Assets http.FileSystem

func init() {
	assetsDir := os.DirFS(filepath.Join(pkgDir, "assets"))
	appDir := os.DirFS(filepath.Join(pkgDir, "app"))

	prefixedAssets, err := newAppFS(appDir, assetsDir)
	if err != nil {
		log.Fatal(err)
	}
	Assets = http.FS(prefixedAssets)
}
