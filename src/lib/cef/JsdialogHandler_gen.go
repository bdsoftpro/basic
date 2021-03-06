// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"lib/toolbox/errs"
	"lib/toolbox/log/jot"
)

import (
	// #include "JsdialogHandler_gen.h"
	"C"
)

// JsdialogHandlerProxy defines methods required for using JsdialogHandler.
type JsdialogHandlerProxy interface {
	OnJsdialog(self *JsdialogHandler, browser *Browser, originURL string, dialogType JsdialogType, messageText, defaultPromptText string, callback *JsdialogCallback, suppressMessage *int32) int32
	OnBeforeUnloadDialog(self *JsdialogHandler, browser *Browser, messageText string, isReload int32, callback *JsdialogCallback) int32
	OnResetDialogState(self *JsdialogHandler, browser *Browser)
	OnDialogClosed(self *JsdialogHandler, browser *Browser)
}

// JsdialogHandler (cef_jsdialog_handler_t from include/capi/cef_jsdialog_handler_capi.h)
// Implement this structure to handle events related to JavaScript dialogs. The
// functions of this structure will be called on the UI thread.
type JsdialogHandler C.cef_jsdialog_handler_t

// NewJsdialogHandler creates a new JsdialogHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewJsdialogHandler(proxy JsdialogHandlerProxy) *JsdialogHandler {
	result := (*JsdialogHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_jsdialog_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_jsdialog_handler_proxy(result.toNative())
	}
	return result
}

func (d *JsdialogHandler) toNative() *C.cef_jsdialog_handler_t {
	return (*C.cef_jsdialog_handler_t)(d)
}

func lookupJsdialogHandlerProxy(obj *BaseRefCounted) JsdialogHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(JsdialogHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type JsdialogHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *JsdialogHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnJsdialog (on_jsdialog)
// Called to run a JavaScript dialog. If |origin_url| is non-NULL it can be
// passed to the CefFormatUrlForSecurityDisplay function to retrieve a secure
// and user-friendly display string. The |default_prompt_text| value will be
// specified for prompt dialogs only. Set |suppress_message| to true (1) and
// return false (0) to suppress the message (suppressing messages is
// preferable to immediately executing the callback as this is used to detect
// presumably malicious behavior like spamming alert messages in
// onbeforeunload). Set |suppress_message| to false (0) and return false (0)
// to use the default implementation (the default implementation will show one
// modal dialog at a time and suppress any additional dialog requests until
// the displayed dialog is dismissed). Return true (1) if the application will
// use a custom dialog or if the callback has been executed immediately.
// Custom dialogs may be either modal or modeless. If a custom dialog is used
// the application must execute |callback| once the custom dialog is
// dismissed.
func (d *JsdialogHandler) OnJsdialog(browser *Browser, originURL string, dialogType JsdialogType, messageText, defaultPromptText string, callback *JsdialogCallback, suppressMessage *int32) int32 {
	return lookupJsdialogHandlerProxy(d.Base()).OnJsdialog(d, browser, originURL, dialogType, messageText, defaultPromptText, callback, suppressMessage)
}

//nolint:gocritic
//export gocef_jsdialog_handler_on_jsdialog
func gocef_jsdialog_handler_on_jsdialog(self *C.cef_jsdialog_handler_t, browser *C.cef_browser_t, originURL *C.cef_string_t, dialogType C.cef_jsdialog_type_t, messageText *C.cef_string_t, defaultPromptText *C.cef_string_t, callback *C.cef_jsdialog_callback_t, suppressMessage *C.int) C.int {
	me__ := (*JsdialogHandler)(self)
	proxy__ := lookupJsdialogHandlerProxy(me__.Base())
	originURL_ := cefstrToString(originURL)
	messageText_ := cefstrToString(messageText)
	defaultPromptText_ := cefstrToString(defaultPromptText)
	return C.int(proxy__.OnJsdialog(me__, (*Browser)(browser), originURL_, JsdialogType(dialogType), messageText_, defaultPromptText_, (*JsdialogCallback)(callback), (*int32)(suppressMessage)))
}

// OnBeforeUnloadDialog (on_before_unload_dialog)
// Called to run a dialog asking the user if they want to leave a page. Return
// false (0) to use the default dialog implementation. Return true (1) if the
// application will use a custom dialog or if the callback has been executed
// immediately. Custom dialogs may be either modal or modeless. If a custom
// dialog is used the application must execute |callback| once the custom
// dialog is dismissed.
func (d *JsdialogHandler) OnBeforeUnloadDialog(browser *Browser, messageText string, isReload int32, callback *JsdialogCallback) int32 {
	return lookupJsdialogHandlerProxy(d.Base()).OnBeforeUnloadDialog(d, browser, messageText, isReload, callback)
}

//nolint:gocritic
//export gocef_jsdialog_handler_on_before_unload_dialog
func gocef_jsdialog_handler_on_before_unload_dialog(self *C.cef_jsdialog_handler_t, browser *C.cef_browser_t, messageText *C.cef_string_t, isReload C.int, callback *C.cef_jsdialog_callback_t) C.int {
	me__ := (*JsdialogHandler)(self)
	proxy__ := lookupJsdialogHandlerProxy(me__.Base())
	messageText_ := cefstrToString(messageText)
	return C.int(proxy__.OnBeforeUnloadDialog(me__, (*Browser)(browser), messageText_, int32(isReload), (*JsdialogCallback)(callback)))
}

// OnResetDialogState (on_reset_dialog_state)
// Called to cancel any pending dialogs and reset any saved dialog state. Will
// be called due to events like page navigation irregardless of whether any
// dialogs are currently pending.
func (d *JsdialogHandler) OnResetDialogState(browser *Browser) {
	lookupJsdialogHandlerProxy(d.Base()).OnResetDialogState(d, browser)
}

//nolint:gocritic
//export gocef_jsdialog_handler_on_reset_dialog_state
func gocef_jsdialog_handler_on_reset_dialog_state(self *C.cef_jsdialog_handler_t, browser *C.cef_browser_t) {
	me__ := (*JsdialogHandler)(self)
	proxy__ := lookupJsdialogHandlerProxy(me__.Base())
	proxy__.OnResetDialogState(me__, (*Browser)(browser))
}

// OnDialogClosed (on_dialog_closed)
// Called when the default implementation dialog is closed.
func (d *JsdialogHandler) OnDialogClosed(browser *Browser) {
	lookupJsdialogHandlerProxy(d.Base()).OnDialogClosed(d, browser)
}

//nolint:gocritic
//export gocef_jsdialog_handler_on_dialog_closed
func gocef_jsdialog_handler_on_dialog_closed(self *C.cef_jsdialog_handler_t, browser *C.cef_browser_t) {
	me__ := (*JsdialogHandler)(self)
	proxy__ := lookupJsdialogHandlerProxy(me__.Base())
	proxy__.OnDialogClosed(me__, (*Browser)(browser))
}
