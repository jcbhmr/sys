//go:build wasip1 && wasm32

package wasip1

import (
	"structs"
)

type Variant[Tag ~uint8 | ~uint16 | ~uint32, Shape, Align any] struct {
	_    structs.HostLayout
	Tag  Tag
	_    [0]Align
	Data Shape // [unsafe.Sizeof(*(*Shape)(unsafe.Pointer(nil)))]byte
}
