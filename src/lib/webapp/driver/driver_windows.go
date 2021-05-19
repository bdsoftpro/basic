package driver

import (
	"lib/webapp"
	"lib/webapp/internal/windows"
)

// ForPlatform returns the driver for your platform.
func ForPlatform() webapp.Driver {
	return windows.Driver()
}
