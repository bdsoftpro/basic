package windows

import (
	"syscall"
	"unsafe"

	"lib/cef"
	"lib/toolbox/atexit"
	"lib/toolbox/errs"
	"lib/toolbox/log/jot"
	"lib/webapp"
	"lib/win32"
)

type driver struct {
	instance             win32.HINSTANCE
	windows              map[win32.HWND]*webapp.Window
	menubars             map[win32.HMENU]*webapp.MenuBar
	menus                map[win32.HMENU]*webapp.Menu
	menuitems            map[int]*menuItem
	windowClass          win32.ATOM
	awaitingQuitDecision bool
}

var drv = &driver{
	windows:   make(map[win32.HWND]*webapp.Window),
	menubars:  make(map[win32.HMENU]*webapp.MenuBar),
	menus:     make(map[win32.HMENU]*webapp.Menu),
	menuitems: make(map[int]*menuItem),
}

// Driver returns the Windows implementation of the driver.
func Driver() webapp.Driver {
	return drv
}

func (d *driver) PrepareForStart() error {
	d.instance = win32.GetModuleHandleS("")
	wcx := win32.WNDCLASSEX{
		Size:     uint32(unsafe.Sizeof(win32.WNDCLASSEX{})),
		Style:    win32.CS_HREDRAW | win32.CS_VREDRAW,
		WndProc:  syscall.NewCallback(d.wndProc),
		Instance: d.instance,
		Cursor:   win32.LoadSystemCursor(win32.IDC_ARROW),
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	var err error
	if wcx.ClassName, err = syscall.UTF16PtrFromString(windowClassName); err != nil {
		return errs.NewWithCause("unable to convert className to utf-16", err)
	}
	d.windowClass = win32.RegisterClassEx(&wcx)
	return nil
}

func (d *driver) PrepareForEventLoop() {
	webapp.WillFinishStartupCallback()
	webapp.DidFinishStartupCallback()
}

func (d *driver) RunEventLoop() {
	cef.RunMessageLoop()
	cef.Shutdown()
	atexit.Exit(0)
}

func (d *driver) AttemptQuit() {
	switch webapp.CheckQuitCallback() {
	case webapp.Cancel:
		return
	case webapp.Now:
		d.quit()
	case webapp.Later:
		d.awaitingQuitDecision = true
	}
}

func (d *driver) MayQuitNow(quit bool) {
	if d.awaitingQuitDecision {
		d.awaitingQuitDecision = false
		if quit {
			d.quit()
		}
	} else {
		jot.Error("Call to MayQuitNow without AttemptQuit")
	}
}

func (d *driver) quit() {
	webapp.QuittingCallback()
	win32.PostQuitMessage(0)
	cef.QuitMessageLoop()
}
