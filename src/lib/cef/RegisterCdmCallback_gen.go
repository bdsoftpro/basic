// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"lib/toolbox/errs"
	"lib/toolbox/log/jot"
)

import (
	// #include "RegisterCdmCallback_gen.h"
	"C"
)

// RegisterCdmCallbackProxy defines methods required for using RegisterCdmCallback.
type RegisterCdmCallbackProxy interface {
	OnCdmRegistrationComplete(self *RegisterCdmCallback, result CdmRegistrationError, errorMessage string)
}

// RegisterCdmCallback (cef_register_cdm_callback_t from include/capi/cef_web_plugin_capi.h)
// Implement this structure to receive notification when CDM registration is
// complete. The functions of this structure will be called on the browser
// process UI thread.
type RegisterCdmCallback C.cef_register_cdm_callback_t

// NewRegisterCdmCallback creates a new RegisterCdmCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRegisterCdmCallback(proxy RegisterCdmCallbackProxy) *RegisterCdmCallback {
	result := (*RegisterCdmCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_register_cdm_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_register_cdm_callback_proxy(result.toNative())
	}
	return result
}

func (d *RegisterCdmCallback) toNative() *C.cef_register_cdm_callback_t {
	return (*C.cef_register_cdm_callback_t)(d)
}

func lookupRegisterCdmCallbackProxy(obj *BaseRefCounted) RegisterCdmCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RegisterCdmCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RegisterCdmCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RegisterCdmCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnCdmRegistrationComplete (on_cdm_registration_complete)
// Method that will be called when CDM registration is complete. |result| will
// be CEF_CDM_REGISTRATION_ERROR_NONE if registration completed successfully.
// Otherwise, |result| and |error_message| will contain additional information
// about why registration failed.
func (d *RegisterCdmCallback) OnCdmRegistrationComplete(result CdmRegistrationError, errorMessage string) {
	lookupRegisterCdmCallbackProxy(d.Base()).OnCdmRegistrationComplete(d, result, errorMessage)
}

//nolint:gocritic
//export gocef_register_cdm_callback_on_cdm_registration_complete
func gocef_register_cdm_callback_on_cdm_registration_complete(self *C.cef_register_cdm_callback_t, result C.cef_cdm_registration_error_t, errorMessage *C.cef_string_t) {
	me__ := (*RegisterCdmCallback)(self)
	proxy__ := lookupRegisterCdmCallbackProxy(me__.Base())
	errorMessage_ := cefstrToString(errorMessage)
	proxy__.OnCdmRegistrationComplete(me__, CdmRegistrationError(result), errorMessage_)
}
