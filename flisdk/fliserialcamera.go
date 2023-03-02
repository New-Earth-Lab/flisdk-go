package flisdk

/*
#cgo LDFLAGS: -L/opt/FirstLightImaging/FliSdk/lib/release -lFliSdk
#cgo CFLAGS: -I/opt/FirstLightImaging/FliSdk/include

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

#include "FliSdk_C_V2.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type FliSerialCamera struct {
	FliSdk
}

func (f *FliSerialCamera) GetFps() (float64, error) {
	var cFps C.double
	if !C.FliSerialCamera_getFps_V2(f.context, &cFps) {
		return 0, fmt.Errorf("fliserialcamera: Failed getting FPS")
	}
	return float64(cFps), nil
}

func (f *FliSerialCamera) SetFps(fps float64) error {
	if !C.FliSerialCamera_setFps_V2(f.context, C.double(fps)) {
		return fmt.Errorf("fliserialcamera: Failed setting FPS")
	}
	return nil
}

func (f *FliSerialCamera) GetFpsMax() (float64, error) {
	var cFps C.double
	if !C.FliSerialCamera_getFpsMax_V2(f.context, &cFps) {
		return 0, fmt.Errorf("fliserialcamera: Failed getting FPS")
	}
	return float64(cFps), nil
}

func (f *FliSerialCamera) SendCommand(command string) (string, error) {
	cCommand := C.CString(string(command))
	defer C.free(unsafe.Pointer(cCommand))
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))
	if !C.FliSerialCamera_sendCommand_V2(f.context, cCommand, text, textSize) {
		return "", fmt.Errorf("flicredtwo: Failed setting conversion gain")
	}
	return C.GoString(text), nil
}

func (f *FliSerialCamera) EnableBias(enable bool) error {
	if !C.FliSerialCamera_enableBias_V2(f.context, C.bool(enable)) {
		return fmt.Errorf("fliserialcamera: Failed enable bias: %t", enable)
	}
	return nil
}

func (f *FliSerialCamera) EnableFlat(enable bool) error {
	if !C.FliSerialCamera_enableFlat_V2(f.context, C.bool(enable)) {
		return fmt.Errorf("fliserialcamera: Failed enable flat: %t", enable)
	}
	return nil
}
