package stdmenu

import (
	"fmt"
	"runtime"

	"lib/toolbox"
	"lib/toolbox/cmdline"
	"lib/toolbox/i18n"
	"lib/webapp"
)

// NewHelpMenu creates a standard 'Help' menu.
func NewHelpMenu(aboutHandler func(), includeDevTools bool) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDHelpMenu, i18n.Text("Help"))
	if runtime.GOOS != toolbox.MacOS {
		menu.InsertItem(-1, webapp.MenuIDAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), nil, 0, nil, aboutHandler)
	}
	if includeDevTools {
		if menu.Count() > 0 {
			menu.InsertSeparator(-1)
		}
		menu.InsertItem(-1, webapp.MenuIDDevToolsItem, i18n.Text("Toggle Development Tools"), nil, 0, validateToggleDevTools, toggleDevTools)
	}
	return menu
}

func validateToggleDevTools() bool {
	wnd := webapp.KeyWindow()
	return wnd != nil && wnd.Browser != nil
}

func toggleDevTools() {
	if wnd := webapp.KeyWindow(); wnd != nil && wnd.Browser != nil {
		wnd.ToggleDevTools()
	}
}
