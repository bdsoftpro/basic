package main

import (
	"net/http"

	"app/server"
	"lib/toolbox/atexit"
	"lib/toolbox/cmdline"
	"lib/toolbox/log/jot"
	"lib/toolbox/log/jotrotate"
	"lib/webapp"
	"lib/webapp/driver"
)
func main() {
	go func() {
		http.ListenAndServe(":8372", http.HandlerFunc(routers.Serve))
	}()
	cmdline.AppName = "Shop Management"
	cmdline.AppCmdName = "shopmanagement"
	cmdline.AppVersion = "1.0.0"
	cmdline.CopyrightYears = "2020-2021"
	cmdline.CopyrightHolder = "Md Delwar Hossain"
	cmdline.AppIdentifier = "com.przon.shopmanagement"

	args, err := webapp.Initialize(driver.ForPlatform())
	jot.FatalIfErr(err)

	cl := cmdline.New(true)
	jotrotate.ParseAndSetup(cl)
	webapp.WillFinishStartupCallback = finishStartup

	// Start only returns on error
	jot.FatalIfErr(webapp.Start(args, nil, nil))
	atexit.Exit(0)
}
func finishStartup() {
	wnd, err := webapp.NewWindow(webapp.StdWindowMask, webapp.MainDisplay().UsableBounds, "Shop Management", "http://127.0.0.1:8372")
	jot.FatalIfErr(err)
	wnd.ToFront()
}