package macos

import "lib/cef"

func (d *driver) OnPreKeyEvent(event *cef.KeyEvent, isKeyboardShortcut *int32) int32 {
	return 0
}

func (d *driver) OnKeyEvent(event *cef.KeyEvent) int32 {
	return 0
}
