// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"lib/toolbox/errs"
	"lib/toolbox/log/jot"
)

import (
	// #include "AudioHandler_gen.h"
	"C"
)

// AudioHandlerProxy defines methods required for using AudioHandler.
type AudioHandlerProxy interface {
	OnAudioStreamStarted(self *AudioHandler, browser *Browser, audioStreamID, channels int32, channelLayout ChannelLayout, sampleRate, framesPerBuffer int32)
	OnAudioStreamPacket(self *AudioHandler, browser *Browser, audioStreamID int32, data **float32, frames int32, pts int64)
	OnAudioStreamStopped(self *AudioHandler, browser *Browser, audioStreamID int32)
}

// AudioHandler (cef_audio_handler_t from include/capi/cef_audio_handler_capi.h)
// Implement this structure to handle audio events All functions will be called
// on the UI thread
type AudioHandler C.cef_audio_handler_t

// NewAudioHandler creates a new AudioHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewAudioHandler(proxy AudioHandlerProxy) *AudioHandler {
	result := (*AudioHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_audio_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_audio_handler_proxy(result.toNative())
	}
	return result
}

func (d *AudioHandler) toNative() *C.cef_audio_handler_t {
	return (*C.cef_audio_handler_t)(d)
}

func lookupAudioHandlerProxy(obj *BaseRefCounted) AudioHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(AudioHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type AudioHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *AudioHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnAudioStreamStarted (on_audio_stream_started)
// Called when the stream identified by |audio_stream_id| has started.
// |audio_stream_id| will uniquely identify the stream across all future
// cef_audio_handler_t callbacks. OnAudioSteamStopped will always be called
// after OnAudioStreamStarted; both functions may be called multiple times for
// the same stream. |channels| is the number of channels, |channel_layout| is
// the layout of the channels and |sample_rate| is the stream sample rate.
// |frames_per_buffer| is the maximum number of frames that will occur in the
// PCM packet passed to OnAudioStreamPacket.
func (d *AudioHandler) OnAudioStreamStarted(browser *Browser, audioStreamID, channels int32, channelLayout ChannelLayout, sampleRate, framesPerBuffer int32) {
	lookupAudioHandlerProxy(d.Base()).OnAudioStreamStarted(d, browser, audioStreamID, channels, channelLayout, sampleRate, framesPerBuffer)
}

//nolint:gocritic
//export gocef_audio_handler_on_audio_stream_started
func gocef_audio_handler_on_audio_stream_started(self *C.cef_audio_handler_t, browser *C.cef_browser_t, audioStreamID C.int, channels C.int, channelLayout C.cef_channel_layout_t, sampleRate C.int, framesPerBuffer C.int) {
	me__ := (*AudioHandler)(self)
	proxy__ := lookupAudioHandlerProxy(me__.Base())
	proxy__.OnAudioStreamStarted(me__, (*Browser)(browser), int32(audioStreamID), int32(channels), ChannelLayout(channelLayout), int32(sampleRate), int32(framesPerBuffer))
}

// OnAudioStreamPacket (on_audio_stream_packet)
// Called when a PCM packet is received for the stream identified by
// |audio_stream_id|. |data| is an array representing the raw PCM data as a
// floating point type, i.e. 4-byte value(s). |frames| is the number of frames
// in the PCM packet. |pts| is the presentation timestamp (in milliseconds
// since the Unix Epoch) and represents the time at which the decompressed
// packet should be presented to the user. Based on |frames| and the
// |channel_layout| value passed to OnAudioStreamStarted you can calculate the
// size of the |data| array in bytes.
func (d *AudioHandler) OnAudioStreamPacket(browser *Browser, audioStreamID int32, data **float32, frames int32, pts int64) {
	lookupAudioHandlerProxy(d.Base()).OnAudioStreamPacket(d, browser, audioStreamID, data, frames, pts)
}

//nolint:gocritic
//export gocef_audio_handler_on_audio_stream_packet
func gocef_audio_handler_on_audio_stream_packet(self *C.cef_audio_handler_t, browser *C.cef_browser_t, audioStreamID C.int, data **C.float, frames C.int, pts C.int64) {
	me__ := (*AudioHandler)(self)
	proxy__ := lookupAudioHandlerProxy(me__.Base())
	proxy__.OnAudioStreamPacket(me__, (*Browser)(browser), int32(audioStreamID), (**float32)(unsafe.Pointer(data)), int32(frames), int64(pts))
}

// OnAudioStreamStopped (on_audio_stream_stopped)
// Called when the stream identified by |audio_stream_id| has stopped.
// OnAudioSteamStopped will always be called after OnAudioStreamStarted; both
// functions may be called multiple times for the same stream.
func (d *AudioHandler) OnAudioStreamStopped(browser *Browser, audioStreamID int32) {
	lookupAudioHandlerProxy(d.Base()).OnAudioStreamStopped(d, browser, audioStreamID)
}

//nolint:gocritic
//export gocef_audio_handler_on_audio_stream_stopped
func gocef_audio_handler_on_audio_stream_stopped(self *C.cef_audio_handler_t, browser *C.cef_browser_t, audioStreamID C.int) {
	me__ := (*AudioHandler)(self)
	proxy__ := lookupAudioHandlerProxy(me__.Base())
	proxy__.OnAudioStreamStopped(me__, (*Browser)(browser), int32(audioStreamID))
}
