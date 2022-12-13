// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

// insert-user is a utility to create a new member and fallback password for them
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/ssh/terminal"

	refs "github.com/ssbc/go-ssb-refs"
	"github.com/ssbc/go-ssb-room/v2/internal/repo"
	"github.com/ssbc/go-ssb-room/v2/roomdb"
	"github.com/ssbc/go-ssb-room/v2/roomdb/sqlite"
)

func main() {
	u, err := user.Current()
	check(err)

	var (
		role     roomdb.Role = roomdb.RoleAdmin
		repoPath string
		password string
	)

	flag.StringVar(&repoPath, "repo", filepath.Join(u.HomeDir, ".ssb-go-room"), "[optional] where the locally stored files of the room are located")
	flag.Func("role", "[optional] which role the new member should have (values: mod[erator], admin, or member. default is admin)", func(val string) error {
		switch strings.ToLower(val) {
		case "admin":
			role = roomdb.RoleAdmin
		case "mod":
			fallthrough
		case "moderator":
			role = roomdb.RoleModerator
		case "member":
			role = roomdb.RoleMember
		default:
			return fmt.Errorf("unknown member role: %q", val)
		}

		return nil
	})
	flag.StringVar(&password, "password", "", "[optional] specify a password using -password, or be prompted to set the password instead")

	flag.Parse()

	if _, err := os.Stat(repoPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error: %s does not exist (-repo)?\n", repoPath)
			os.Exit(1)
		}
	}

	// we require one more argument which is not a flag.
	if len(flag.Args()) != 1 {
		cliMissingArguments("please provide a public key")
	}

	pubKey, err := refs.ParseFeedRef(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid ssb public-key reference:", err)
		os.Exit(1)
	}

	r := repo.New(repoPath)
	db, err := sqlite.Open(r, nil)
	check(err)
	defer db.Close()

	bytePassword := password

	if bytePassword == "" {
		fmt.Fprintln(os.Stderr, "Choose a password to be able to log into the web frontend: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		check(err)

		fmt.Fprintln(os.Stderr, "Repeat Password: ")
		bytePasswordRepeat, err := terminal.ReadPassword(int(syscall.Stdin))
		check(err)

		if !bytes.Equal(bytePassword, bytePasswordRepeat) {
			fmt.Fprintln(os.Stderr, "Passwords didn't match")
			os.Exit(1)
		}
	}

	ctx := context.Background()
	mid, err := db.Members.Add(ctx, pubKey, role)
	check(err)

	err = db.AuthFallback.SetPassword(ctx, mid, string(bytePassword))
	check(err)

	fmt.Fprintf(os.Stderr, "Created member (%s) with ID %d\n", role, mid)
}

func cliMissingArguments(message string) {
	executable := strings.TrimPrefix(os.Args[0], "./")
	fmt.Fprintf(os.Stderr, "%s: %s\nusage:%s <optional flags> <@base64-encoded-public-key=.ed25519>\n", executable, message, executable)
	flag.Usage()
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
