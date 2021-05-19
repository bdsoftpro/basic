package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"lib/toolbox"
	"lib/toolbox/atexit"
)

var (
	installPrefix = "cef"
	cefPlatform   string
)

func checkPlatform() {
	switch runtime.GOOS {
	case toolbox.MacOS:
		cefPlatform = "macosx64"
	case toolbox.LinuxOS:
		cefPlatform = "linux64"
	case toolbox.WindowsOS:
		if os.Getenv("MSYSTEM") != "MINGW64" {
			fmt.Println("Windows is only supported through the use of MINGW64")
			atexit.Exit(1)
		}
		cefPlatform = "windows64"
		pwd, _ := os.Getwd()
		installPrefix = path.Join(strings.ReplaceAll(pwd, "\\", "/"), installPrefix)
	default:
		fmt.Println("Unsupported OS: ", runtime.GOOS)
		atexit.Exit(1)
	}
}
