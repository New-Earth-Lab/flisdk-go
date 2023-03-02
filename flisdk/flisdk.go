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
	"runtime/cgo"
	"strings"
	"unsafe"
)

type FliSdk struct {
	context C.FliContext
}

type NewImageAvailableCallBack C.newImageAvailableCallBack

type CallbackHandler struct {
	callbackHandler C.callbackHandler
	handle          cgo.Handle
}

type CroppingData struct {
	Col1    uint16
	Col2    uint16
	Row1    uint16
	Row2    uint16
	Enabled bool
}

const (
	textSize = 512
)

type mode int

const (
	Mode_Full mode = iota
	Mode_GrabOnly
	Mode_ConfigOnly
)

func Init() (*FliSdk, error) {
	context := C.FliSdk_init_V2()
	if context == nil {
		return nil, fmt.Errorf("flisdk: Unable to init SDK")
	}

	return &FliSdk{
		context: context,
	}, nil
}

func (f *FliSdk) Exit() {
	C.FliSdk_exit_V2(f.context)
}

func (f *FliSdk) DetectGrabbers() ([]string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	C.FliSdk_detectGrabbers_V2(f.context, text, textSize)
	grabberStrings := strings.Split(C.GoString(text), ";")
	if len(grabberStrings) == 0 {
		return nil, fmt.Errorf("flisdk: No grabbers found")
	}

	return grabberStrings, nil
}

func (f *FliSdk) GetDetectedGrabbers() ([]string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	C.FliSdk_getDetectedGrabbers_V2(f.context, text, textSize)
	grabberStrings := strings.Split(C.GoString(text), ";")
	if len(grabberStrings) == 0 {
		return nil, fmt.Errorf("flisdk: No grabbers found")
	}

	return grabberStrings, nil
}

func (f *FliSdk) DetectCameras() ([]string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	C.FliSdk_detectCameras_V2(f.context, text, textSize)
	cameraStrings := strings.Split(C.GoString(text), ";")
	if len(cameraStrings) == 0 {
		return nil, fmt.Errorf("flisdk: No cameras found")
	}

	return cameraStrings, nil
}

func (f *FliSdk) GetDetectedCameras() ([]string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	C.FliSdk_getDetectedCameras_V2(f.context, text, textSize)
	cameraStrings := strings.Split(C.GoString(text), ";")
	if len(cameraStrings) == 0 {
		return nil, fmt.Errorf("flisdk: No cameras found")
	}

	return cameraStrings, nil
}

func (f *FliSdk) SetGrabber(grabberName string) error {
	cGrabberName := C.CString(grabberName)
	defer C.free(unsafe.Pointer(cGrabberName))

	if !C.FliSdk_setGrabber_V2(f.context, cGrabberName) {
		return fmt.Errorf("flisdk: Unable to set grabber: %s", grabberName)
	}

	return nil
}

func (f *FliSdk) SetCamera(cameraName string) error {
	cCameraName := C.CString(cameraName)
	defer C.free(unsafe.Pointer(cCameraName))

	if !C.FliSdk_setCamera_V2(f.context, cCameraName) {
		return fmt.Errorf("flisdk: Unable to set camera: %s", cameraName)
	}

	return nil
}

func (f *FliSdk) GetCurrentCameraName() (string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	if !C.FliSdk_getCurrentCameraName_V2(f.context, text, textSize) {
		return "", fmt.Errorf("flisdk: Could not get camera name")
	}

	return C.GoString(text), nil
}

func (f *FliSdk) SetMode(mode mode) error {
	var cMode C.Mode
	switch mode {
	case Mode_Full:
		cMode = C.Full
	case Mode_GrabOnly:
		cMode = C.GrabOnly
	case Mode_ConfigOnly:
		cMode = C.ConfigOnly
	default:
		return fmt.Errorf("flisdk: Invalid mode %d", mode)
	}

	C.FliSdk_setMode_V2(f.context, cMode)

	return nil
}

func (f *FliSdk) SetImageDimension(width, height uint16) {
	C.FliSdk_setImageDimension_V2(f.context, C.uint16_t(width), C.uint16_t(height))
}

func (f *FliSdk) Update() error {
	if !C.FliSdk_update_V2(f.context) {
		return fmt.Errorf("flisdk: Unable to update SDK")
	}

	return nil
}

func (f *FliSdk) Start() error {
	if !C.FliSdk_start_V2(f.context) {
		return fmt.Errorf("flisdk: Unable to start camera")
	}

	return nil
}

func (f *FliSdk) Stop() error {
	if !C.FliSdk_stop_V2(f.context) {
		return fmt.Errorf("flisdk: Unable to stop camera")
	}

	return nil
}

func (f *FliSdk) IsStarted() bool {
	return bool(C.FliSdk_isStarted_V2(f.context))
}

func (f *FliSdk) GetCameraModel() (string, error) {
	text := (*C.char)(C.malloc(textSize))
	defer C.free(unsafe.Pointer(text))

	C.FliSdk_getCameraModelAsString_V2(f.context, text, textSize)

	if text == nil {
		return "", fmt.Errorf("flisdk: Could not get camera model")
	}

	return C.GoString(text), nil
}

func (f *FliSdk) EnableGrabN(numFrames uint32) error {
	if !C.FliSdk_enableGrabN_V2(f.context, C.uint32_t(numFrames)) {
		return fmt.Errorf("flisdk: Unable to enable grab N: %d", numFrames)
	}

	return nil
}

func (f *FliSdk) DisableGrabN() error {
	if !C.FliSdk_disableGrabN_V2(f.context) {
		return fmt.Errorf("flisdk: Unable to disable grab")
	}

	return nil
}

func (f *FliSdk) IsGrabNEnabled() bool {
	return bool(C.FliSdk_isGrabNEnabled_V2(f.context))
}

func (f *FliSdk) IsGrabNFinished() bool {
	return bool(C.FliSdk_isGrabNFinished_V2(f.context))
}

func (f *FliSdk) GetRawImage(index uint64) (unsafe.Pointer, error) {
	image := unsafe.Pointer(C.FliSdk_getRawImage_V2(f.context, C.int64_t(index)))

	if image == nil {
		return nil, fmt.Errorf("flisdk: Invalid index: %d", index)
	}

	return image, nil
}

func (f *FliSdk) GetRawImageBytes(index uint64) ([]byte, error) {
	width, height := f.GetCurrentImageDimension()
	len := uint(width) * uint(height) * f.GetBytesPerPixel()
	image := C.GoBytes(unsafe.Pointer(C.FliSdk_getRawImage_V2(f.context, C.int64_t(index))), C.int(len))

	if image == nil {
		return nil, fmt.Errorf("flisdk: Invalid index: %d", index)
	}

	return image, nil
}

func (f *FliSdk) InitLog(appName string) {
	cAppName := C.CString(appName)
	defer C.free(unsafe.Pointer(cAppName))

	C.FliSdk_initLog_V2(f.context, cAppName)
}

func (f *FliSdk) GetFps() uint32 {
	return uint32(C.FliSdk_getFps_V2(f.context))
}

func (f *FliSdk) GetBufferFilling() uint32 {
	return uint32(C.FliSdk_getBufferFilling_V2(f.context))
}

func (f *FliSdk) SetBufferSizeInImages(numImages uint64) {
	C.FliSdk_setBufferSizeInImages_V2(f.context, C.uint64_t(numImages))
}

func (f *FliSdk) SetBufferSize(sizeMb uint16) {
	C.FliSdk_setBufferSize_V2(f.context, C.uint16_t(sizeMb))
}

func (f *FliSdk) GetBufferSize() uint16 {
	return uint16(C.FliSdk_getBufferSize_V2(f.context))
}

func (f *FliSdk) ResetBuffer() {
	C.FliSdk_resetBuffer_V2(f.context)
}

func (f *FliSdk) IsCroppingDataValid(croppingData CroppingData) bool {
	cCroppingValid := C.CroppingData_C{
		col1: C.uint16_t(croppingData.Col1),
		col2: C.uint16_t(croppingData.Col2),
		row1: C.uint16_t(croppingData.Row1),
		row2: C.uint16_t(croppingData.Row2),
	}

	return bool(C.FliSdk_isCroppingDataValid_V2(f.context, cCroppingValid))
}

func (f *FliSdk) GetCroppingState() (CroppingData, error) {
	cCroppingData := C.CroppingData_C{}
	var cIsEnabled C.bool

	if !C.FliSdk_getCroppingState_V2(f.context, &cIsEnabled, &cCroppingData) {
		return CroppingData{}, fmt.Errorf("flisdk: Failed getting cropping state")
	}

	return CroppingData{
		Col1:    uint16(cCroppingData.col1),
		Col2:    uint16(cCroppingData.col2),
		Row1:    uint16(cCroppingData.row1),
		Row2:    uint16(cCroppingData.row2),
		Enabled: bool(cIsEnabled),
	}, nil
}

func (f *FliSdk) SetCroppingState(croppingData CroppingData) error {
	cCroppingData := C.CroppingData_C{
		col1: C.uint16_t(croppingData.Col1),
		col2: C.uint16_t(croppingData.Col2),
		row1: C.uint16_t(croppingData.Row1),
		row2: C.uint16_t(croppingData.Row2),
	}

	if !C.FliSdk_setCroppingState_V2(f.context, C.bool(croppingData.Enabled), cCroppingData) {
		return fmt.Errorf("flisdk: Failed setting cropping state")
	}

	return nil
}

func (f *FliSdk) AddCallbackNewImage(callback NewImageAvailableCallBack,
	fpsTrigger uint16, beforeCopy bool, ctx unsafe.Pointer) CallbackHandler {
	handle := cgo.NewHandle(ctx)

	callbackHandler := C.FliSdk_addCallbackNewImage_V2(f.context, callback,
		C.uint16_t(fpsTrigger), C.bool(beforeCopy), unsafe.Pointer(&handle))

	return CallbackHandler{
		callbackHandler: callbackHandler,
		handle:          handle,
	}
}

func (f *FliSdk) RemoveCallbackNewImage(callbackHandler CallbackHandler) {
	C.FliSdk_removeCallbackNewImage_V2(f.context, callbackHandler.callbackHandler)
	callbackHandler.handle.Delete()
}

func (f *FliSdk) GetCurrentImageDimension() (uint16, uint16) {
	var width, height C.uint16_t
	C.FliSdk_getCurrentImageDimension_V2(f.context, &width, &height)

	return uint16(width), uint16(height)
}

func (f *FliSdk) SetNumberImagesPerBuffer(numImages uint8) {
	C.FliSdk_setNbImagesPerBuffer_V2(f.context, C.uint8_t(numImages))
}

func (f *FliSdk) IsUnsignedPixel() bool {
	return bool(C.FliSdk_isUnsignedPixel_V2(f.context))
}

func (f *FliSdk) IsMono8Pixel() bool {
	return bool(C.FliSdk_isMono8Pixel_V2(f.context))
}

func (f *FliSdk) GetBytesPerPixel() uint {
	if f.IsMono8Pixel() {
		return 1
	} else {
		return 2
	}
}

func (f *FliSdk) GetImageSizeInBytes() uint {
	width, height := f.GetCurrentImageDimension()
	return uint(width) * uint(height) * f.GetBytesPerPixel()
}

func (f *FliSdk) EnableUnsignedPixel(enable bool) {
	C.FliSdk_enableUnsignedPixel_V2(f.context, C.bool(enable))
}

func (f *FliSdk) EnableRingBuffer(enable bool) {
	C.FliSdk_enableRingBuffer_V2(f.context, C.bool(enable))
}

func (f *FliSdk) GetNumCountError() uint64 {
	return uint64(C.FliSdk_getNbCountError_V2(f.context))
}
