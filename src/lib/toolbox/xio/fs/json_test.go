// Copyright ©2016-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package fs_test

import (
	"io/ioutil"
	"os"
	"testing"

	"lib/toolbox/xio/fs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadSaveJSON(t *testing.T) {
	type data struct {
		Name  string
		Count int
	}
	value := &data{
		Name:  "Rich",
		Count: 22,
	}
	f, err := ioutil.TempFile("", "json_test")
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.NoError(t, fs.SaveJSON(f.Name(), value, false))
	var value2 data
	require.NoError(t, fs.LoadJSON(f.Name(), &value2))
	require.NoError(t, os.Remove(f.Name()))
	assert.Equal(t, value, &value2)
}
