package mrt

import (
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/usecases/internal"
	"github.com/enbility/ship-go/logging"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// HandleEvent handles events for the CRCSF use case
func (e *MRT) HandleEvent(payload spineapi.EventPayload) {
	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if internal.IsEntityConnected(payload) {
		e.deviceConnected(payload.Entity)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	switch payload.Data.(type) {
	case *model.MeasurementDescriptionListDataType:
		e.deviceMeasurementDescriptionDataUpdate(payload.Entity)

	case *model.MeasurementListDataType:
		e.deviceMeasurementDataUpdate(payload)
	}
}

// process required steps when a device is connected
func (e *MRT) deviceConnected(entity spineapi.EntityRemoteInterface) {
	if measurement, err := client.NewMeasurement(e.LocalEntity, entity); err == nil {
		if !measurement.HasSubscription() {
			if _, err := measurement.Subscribe(); err != nil {
				logging.Log().Error(err)
			}
		}

		// Get measurement parameters
		if _, err := measurement.RequestDescriptions(nil, nil); err != nil {
			logging.Log().Error(err)
		}

		if _, err := measurement.RequestConstraints(nil, nil); err != nil {
			logging.Log().Error(err)
		}
	}
}

// The measurement description data of a device was updated
func (e *MRT) deviceMeasurementDescriptionDataUpdate(entity spineapi.EntityRemoteInterface) {
	if measurement, err := client.NewMeasurement(e.LocalEntity, entity); err == nil {
		// measurement descriptions received, now get the data
		if _, err := measurement.RequestData(nil, nil); err != nil {
			logging.Log().Error("Error getting measurement list values:", err)
		}
	}
}

// The measurement data of a device was updated
func (e *MRT) deviceMeasurementDataUpdate(payload spineapi.EventPayload) {
	if measurement, err := client.NewMeasurement(e.LocalEntity, payload.Entity); err == nil {
		// Scenario 1
		filter := model.MeasurementDescriptionDataType{
			ScopeType: util.Ptr(model.ScopeTypeTypeRoomAirTemperature),
		}
		if measurement.CheckEventPayloadDataForFilter(payload.Data, filter) && e.EventCB != nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateRoomTemperature)
		}
	}
}
