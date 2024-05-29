package evsecc

import (
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/usecases/internal"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// handle SPINE events
func (e *CemEVSECC) HandleEvent(payload spineapi.EventPayload) {
	// only about events from an EVSE entity or device changes for this remote device

	if !e.IsCompatibleEntity(payload.Entity) {
		return
	}

	if internal.IsEntityConnected(payload) {
		e.evseConnected(payload)
		return
	} else if internal.IsEntityDisconnected(payload) {
		e.evseDisconnected(payload)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	switch payload.Data.(type) {
	case *model.DeviceClassificationManufacturerDataType:
		e.evseManufacturerDataUpdate(payload)
	case *model.DeviceDiagnosisStateDataType:
		e.evseStateUpdate(payload)
	}
}

// an EVSE was connected
func (e *CemEVSECC) evseConnected(payload spineapi.EventPayload) {
	if evseDeviceClassification, err := client.NewDeviceClassification(e.LocalEntity, payload.Entity); err == nil {
		_, _ = evseDeviceClassification.RequestManufacturerDetails()
	}

	if evseDeviceDiagnosis, err := client.NewDeviceDiagnosis(e.LocalEntity, payload.Entity); err == nil {
		_, _ = evseDeviceDiagnosis.RequestState()
	}

	e.EventCB(payload.Ski, payload.Device, payload.Entity, EvseConnected)
}

// an EVSE was disconnected
func (e *CemEVSECC) evseDisconnected(payload spineapi.EventPayload) {
	e.EventCB(payload.Ski, payload.Device, payload.Entity, EvseDisconnected)
}

// the manufacturer Data of an EVSE was updated
func (e *CemEVSECC) evseManufacturerDataUpdate(payload spineapi.EventPayload) {
	if evDeviceClassification, err := client.NewDeviceClassification(e.LocalEntity, payload.Entity); err == nil {
		if _, err := evDeviceClassification.GetManufacturerDetails(); err == nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateManufacturerData)
		}
	}
}

// the operating State of an EVSE was updated
func (e *CemEVSECC) evseStateUpdate(payload spineapi.EventPayload) {
	if evDeviceDiagnosis, err := client.NewDeviceDiagnosis(e.LocalEntity, payload.Entity); err == nil {
		if _, err := evDeviceDiagnosis.GetState(); err == nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOperatingState)
		}
	}
}