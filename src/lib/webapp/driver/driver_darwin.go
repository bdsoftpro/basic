package driver

import (
	"lib/webapp"
	"lib/webapp/internal/macos"
)

// ForPlatform returns the driver for your platform.
func ForPlatform() webapp.Driver {
	return macos.Driver()
}
