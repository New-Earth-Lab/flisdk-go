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

type FliCredTwo struct {
	FliSdk
}

type CredTwoAllTemp struct {
	Motherboard float64
	FrontEnd    float64
	Sensor      float64
	PowerBoard  float64
	Peltier     float64
	HeatSink    float64
}

type CredTwoAllPeltierPower struct {
	ExternalPeltierCurrent float64
	ExternalPeltierVoltage float64
	ExternalPeltierPower   float64
	InternalPeltierCurrent float64
	InternalPeltierVoltage float64
	InternalPeltierPower   float64
}

type CredTwoSensorPower struct {
	Current float64
	Voltage float64
	Power   float64
}

type CredTwoGain string

const (
	CredTwoGain_Low    CredTwoGain = "low"
	CredTwoGain_Medium CredTwoGain = "medium"
	CredTwoGain_High   CredTwoGain = "high"
)

func (f *FliCredTwo) GetAllTemp() (CredTwoAllTemp, error) {
	var cMb, cFe, cPw, cSensor, cPeltier, cHeatsink C.double
	if !C.FliCredTwo_getAllTemp_V2(f.context, &cMb, &cFe, &cPw, &cSensor, &cPeltier, &cHeatsink) {
		return CredTwoAllTemp{}, fmt.Errorf("flisdk: Failed reading all temperatures")
	}

	return CredTwoAllTemp{
		Motherboard: float64(cMb),
		FrontEnd:    float64(cFe),
		Sensor:      float64(cSensor),
		PowerBoard:  float64(cPw),
		Peltier:     float64(cPeltier),
		HeatSink:    float64(cHeatsink),
	}, nil
}

func (f *FliCredTwo) GetTint() (float64, error) {
	var val C.double
	if !C.FliCredTwo_getTint_V2(f.context, &val) {
		return 0, fmt.Errorf("flicredtwo: Failed getting Tint")
	}
	return float64(val), nil
}

func (f *FliCredTwo) GetTintMax() (float64, error) {
	var val C.double
	if !C.FliCredTwo_getTintMax_V2(f.context, &val) {
		return 0, fmt.Errorf("flicredtwo: Failed getting Tint")
	}
	return float64(val), nil
}

func (f *FliCredTwo) SetTint(val float64) error {
	if !C.FliCredTwo_setTint_V2(f.context, C.double(val)) {
		return fmt.Errorf("flicredtwo: Failed setting Tint")
	}
	return nil
}

func (f *FliCredTwo) SetConversionGain(val CredTwoGain) error {
	cGain := C.CString(string(val))
	defer C.free(unsafe.Pointer(cGain))
	if !C.FliCredTwo_setConversionGain_V2(f.context, cGain) {
		return fmt.Errorf("flicredtwo: Failed setting conversion gain")
	}
	return nil
}

func (f *FliCredTwo) SetSensorTemp(val float64) error {
	if !C.FliCredTwo_setSensorTemp_V2(f.context, C.double(val)) {
		return fmt.Errorf("flicredtwo: Failed setting temperature")
	}
	return nil
}

func (f *FliCredTwo) StartHttpServer() error {
	if !C.FliCredTwo_startHttpServer_V2(f.context) {
		return fmt.Errorf("flicredtwo: Unable to start HTTP server")
	}
	return nil
}

func (f *FliCredTwo) StopHttpServer() error {
	if !C.FliCredTwo_stopHttpServer_V2(f.context) {
		return fmt.Errorf("flicredtwo: Unable to stop HTTP server")
	}
	return nil
}

func (f *FliCredTwo) StartEthernetGrabber() error {
	if !C.FliCredTwo_startEthernetGrabber_V2(f.context) {
		return fmt.Errorf("flicredtwo: Unable to start Ethernet server")
	}
	return nil
}

func (f *FliCredTwo) StopEthernetGrabber() error {
	if !C.FliCredTwo_stopEthernetGrabber_V2(f.context) {
		return fmt.Errorf("flicredtwo: Unable to stop Ethernet server")
	}
	return nil
}
