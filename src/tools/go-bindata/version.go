package main

import (
	"fmt"
	"runtime"
)

const (
	AppName         = "bindata"
	AppVersionMajor = 3
	AppVersionMinor = 1
)

// revision part of the program version.
// This will be set automatically at build time like so:
//
//     go build -ldflags "-X main.AppVersionRev `date -u +%s`"
var AppVersionRev string

func Version() string {
	if len(AppVersionRev) == 0 {
		AppVersionRev = "0"
	}

	return fmt.Sprintf("%s %d.%d.%s (Go runtime %s).\nCopyright (c) 2020-2021, Md Delwar.",
		AppName, AppVersionMajor, AppVersionMinor, AppVersionRev, runtime.Version())
}
