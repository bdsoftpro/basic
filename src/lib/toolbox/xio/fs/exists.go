// Copyright ©2016-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package fs

import "os"

// FileExists returns true if the path points to a regular file.
func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		mode := fi.Mode()
		return !mode.IsDir() && mode.IsRegular()
	}
	return false
}
