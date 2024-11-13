package crht

import (
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/usecases/internal"
	"github.com/enbility/ship-go/logging"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"

	"github.com/enbility/spine-go/model"
)

// HandleEvent handles events for the CRHT use case.
func (e *CRHT) HandleEvent(payload spineapi.EventPayload) {
	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if internal.IsEntityConnected(payload) {
		e.hvacRoomConnected(payload.Entity)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	switch payload.Data.(type) {
	case *model.SetpointDescriptionListDataType:
		e.setpointDescriptionsUpdate(payload)

	case *model.SetpointConstraintsListDataType:
		e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateSetpointConstraints)

	case *model.SetpointListDataType:
		e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateSetpoints)

	case *model.HvacSystemFunctionSetpointRelationListDataType,
		*model.HvacOperationModeDescriptionListDataType:
		e.mapSetpointsToOperationModes(payload)
	}
}

func (e *CRHT) mapSetpointsToOperationModes(payload spineapi.EventPayload) {
	hvac, err := client.NewHvac(e.LocalEntity, payload.Entity)
	if err != nil {
		logging.Log().Debug(err)
		return
	}

	// Get the heating system functionId.
	filter := model.HvacSystemFunctionDescriptionDataType{
		SystemFunctionType: util.Ptr(model.HvacSystemFunctionTypeTypeHeating),
	}
	functions, _ := hvac.GetHvacSystemFunctionDescriptionsForFilter(filter)
	if len(functions) != 1 {
		logging.Log().Debug("Expected one heating system function")
		return
	}

	heatingFunctionId := *functions[0].SystemFunctionId

	relations, _ := hvac.GetHvacSystemFunctionSetpointRelationsForSystemFunctionId(heatingFunctionId)
	descriptions, _ := hvac.GetHvacOperationModeDescriptions()
	if len(relations) == 0 || len(descriptions) == 0 {
		return
	}

	modeForModeId := make(map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType)
	for _, description := range descriptions {
		modeForModeId[*description.OperationModeId] = *description.OperationModeType
	}

	// Map the setpoints to their respective operation modes.
	for _, relation := range relations {
		if mode, found := modeForModeId[*relation.OperationModeId]; found {
			if len(relation.SetpointId) == 0 {
				// Only the 'Off' operation mode can have no setpoint associated with it.
				if mode != model.HvacOperationModeTypeTypeOff {
					logging.Log().Errorf("Operation mode '%s' has no setpoints", mode)
				}
			} else if len(relation.SetpointId) == 1 {
				// Store the unique setpoint for the operation mode.
				e.setpointIdsForMode[mode] = relation.SetpointId[0]
			} else if mode != model.HvacOperationModeTypeTypeAuto {
				// Only the 'Auto' operation mode can have multiple setpoints (1 to 4).
				// Since 'Auto' mode is not user-controllable, we do not store the setpoints.
				logging.Log().Errorf("Operation mode '%s' has multiple setpoints", mode)
			}
		}
	}
}

// setpointDescriptionsUpdate processes the necessary steps when setpoint descriptions are updated.
func (e *CRHT) setpointDescriptionsUpdate(payload spineapi.EventPayload) {
	setPoint, err := client.NewSetpoint(e.LocalEntity, payload.Entity)
	if err != nil {
		logging.Log().Debug(err)
		return
	}

	// The setpointConstraintsListData and setpointListData reads should
	// be partial, using setpointId from setpointDescriptionListData.
	setpointDescriptions := payload.Data.(*model.SetpointDescriptionListDataType).SetpointDescriptionData
	for _, setpointDescription := range setpointDescriptions {
		constraintsSelector := &model.SetpointConstraintsListDataSelectorsType{
			SetpointId: setpointDescription.SetpointId,
		}
		if _, err := setPoint.RequestSetPointConstraints(constraintsSelector, nil); err != nil {
			logging.Log().Debug(err)
		}

		setpointSelector := &model.SetpointListDataSelectorsType{
			SetpointId: setpointDescription.SetpointId,
		}
		if _, err := setPoint.RequestSetPoints(setpointSelector, nil); err != nil {
			logging.Log().Debug(err)
		}
	}

	// Request setpoint relations to map operation modes to setpoints.
	if hvac, err := client.NewHvac(e.LocalEntity, payload.Entity); err == nil {
		if _, err := hvac.RequestHvacSystemFunctionSetPointRelations(nil, nil); err != nil {
			logging.Log().Debug(err)
		}
	}
}

// hvacRoomConnected processes required steps when a HVAC Room entity is hvacRoomConnected.
func (e *CRHT) hvacRoomConnected(entity spineapi.EntityRemoteInterface) {
	if hvac, err := client.NewHvac(e.LocalEntity, entity); err == nil {
		if !hvac.HasSubscription() {
			if _, err := hvac.Subscribe(); err != nil {
				logging.Log().Debug(err)
			}
		}
	}

	if setPoint, err := client.NewSetpoint(e.LocalEntity, entity); err == nil {
		if !setPoint.HasSubscription() {
			if _, err := setPoint.Subscribe(); err != nil {
				logging.Log().Debug(err)
			}
		}

		selector := &model.SetpointDescriptionListDataSelectorsType{
			ScopeType: util.Ptr(model.ScopeTypeTypeRoomAirTemperature),
		}
		if _, err := setPoint.RequestSetPointDescriptions(selector, nil); err != nil {
			logging.Log().Debug(err)
		}
	}
}
