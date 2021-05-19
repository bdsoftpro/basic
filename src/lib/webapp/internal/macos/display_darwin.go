package macos

import (
	// #import "displays.h"
	"C"
)

import (
	"unsafe"

	"lib/webapp"
)

func (d *driver) Displays() []*webapp.Display {
	var count C.ulong
	ptr := unsafe.Pointer(C.displays(&count))
	displays := (*[1 << 30]C.Display)(ptr)
	result := make([]*webapp.Display, count)
	for i := range result {
		dsp := &webapp.Display{}
		dsp.Bounds.X = float64(displays[i].bounds.origin.x)
		dsp.Bounds.Y = float64(displays[i].bounds.origin.y)
		dsp.Bounds.Width = float64(displays[i].bounds.size.width)
		dsp.Bounds.Height = float64(displays[i].bounds.size.height)
		dsp.UsableBounds.X = float64(displays[i].usableBounds.origin.x)
		dsp.UsableBounds.Y = float64(displays[i].usableBounds.origin.y)
		dsp.UsableBounds.Width = float64(displays[i].usableBounds.size.width)
		dsp.UsableBounds.Height = float64(displays[i].usableBounds.size.height)
		dsp.ScalingFactor = float64(displays[i].scalingFactor)
		dsp.IsMain = displays[i].isMain != 0
		result[i] = dsp
	}
	C.free(ptr)
	return result
}
