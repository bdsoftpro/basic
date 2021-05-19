package main

import (
	"fmt"
	"os"

	"tools/cef-installer/cmd"
	"lib/toolbox/atexit"
	"lib/toolbox/cmdline"
)

const desiredCEFVersion = "75.1.4+g4210896+chromium-75.0.3770.100"

func main() {
	cmdline.CopyrightYears = "2018-2020"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.cef"
	cl := cmdline.New(true)
	cl.Description = "Utilities for managing setup of the Chromium Embedded Framework."
	cl.AddCommand(cmd.NewInstall(desiredCEFVersion))
	cl.AddCommand(cmd.NewDist())
	if err := cl.RunCommand(cl.Parse(os.Args[1:])); err != nil {
		fmt.Fprintln(os.Stderr, err)
		atexit.Exit(1)
	}
}
