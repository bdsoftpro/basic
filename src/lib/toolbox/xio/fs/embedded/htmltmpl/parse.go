// Copyright ©2016-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Package htmltmpl provides convenience utilities for using html templates in
// an embedded filesystem.
package htmltmpl

import (
	"html/template"
	"os"
	"path"

	"lib/toolbox/errs"
	"lib/toolbox/xio"
	"lib/toolbox/xio/fs/embedded"
)

// Load the templates found at the path, omitting any that the filter function
// returns true for. The filter function may be nil, in which case all files
// are loaded. The filter function is not called for the initial path. The
// template passed in will be used to load new templates and will be returned.
// If the passed in template is nil, a new one will be created.
func Load(tmpl *template.Template, fs embedded.FileSystem, p string, filter func(p string, isDir bool) bool) (*template.Template, error) {
	dir, err := fs.Open(p)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer xio.CloseIgnoringErrors(dir)
	var fi os.FileInfo
	if fi, err = dir.Stat(); err != nil {
		return nil, errs.Wrap(err)
	}
	if fi.IsDir() {
		fis, dirErr := dir.Readdir(-1)
		if dirErr != nil {
			return nil, errs.Wrap(dirErr)
		}
		for _, fi = range fis {
			onePath := path.Join(p, fi.Name())
			isDir := fi.IsDir()
			if filter == nil || !filter(onePath, isDir) {
				if isDir {
					if tmpl, err = Load(tmpl, fs, onePath, filter); err != nil {
						return nil, err
					}
				} else {
					if tmpl, err = load(tmpl, fs, onePath); err != nil {
						return nil, err
					}
				}
			}
		}
	} else if tmpl, err = load(tmpl, fs, p); err != nil {
		return nil, err
	}
	return tmpl, nil
}

func load(tmpl *template.Template, fs embedded.FileSystem, p string) (*template.Template, error) {
	str, ok := fs.ContentAsString(p)
	if !ok {
		return nil, errs.New("Unable to read " + p)
	}
	var t *template.Template
	if tmpl == nil {
		tmpl = template.New(p)
		t = tmpl
	} else {
		t = tmpl.New(p)
	}
	if _, err := t.Parse(str); err != nil {
		return nil, errs.NewWithCause("Unable to parse "+p, err)
	}
	return tmpl, nil
}
