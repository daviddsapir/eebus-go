package cdsf

import (
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/usecases/internal"
	"github.com/enbility/ship-go/logging"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// HandleEvent handles events for the CDSF use case.
func (e *CDSF) HandleEvent(payload spineapi.EventPayload) {
	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if internal.IsEntityConnected(payload) {
		e.dhwCircuitconnected(payload.Entity)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	switch payload.Data.(type) {
	case *model.HvacSystemFunctionDescriptionListDataType:
		e.systemFunctionsDescriptionsUpdate(payload)

	case *model.HvacSystemFunctionOperationModeRelationListDataType:
		e.systemFunctionOperationModeRelationsUpdate(payload)

	case *model.HvacOperationModeDescriptionListDataType,
		*model.HvacSystemFunctionListDataType:
		e.updateOperationModes(payload)

	case *model.HvacOverrunDescriptionListDataType:
		e.updateOverrunDescriptions(payload)

	case *model.HvacOverrunListDataType:
		e.updateOverruns(payload)
	}
}

// updateOverruns handles the overruns update event.
func (e *CDSF) updateOverruns(payload spineapi.EventPayload) {
	if e.overrunId == nil {
		return
	}

	overruns, _ := payload.Data.(*model.HvacOverrunListDataType)
	for _, overrun := range overruns.HvacOverrunData {
		if overrun.OverrunId == nil || *overrun.OverrunId == *e.overrunId {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOverrunStatus)
		}
	}
}

func (e *CDSF) systemFunctionOperationModeRelationsUpdate(payload spineapi.EventPayload) {
	hvac, err := client.NewHvac(e.LocalEntity, payload.Entity)
	if err != nil {
		logging.Log().Debug(err)
		return
	}

	// According to the specification, we need to perform a partial read of hvacOperationModeDescriptionListData
	// using operationModeId derived from hvacSystemFunctionOperationModeRelationListData.
	relations, _ := payload.Data.(*model.HvacSystemFunctionOperationModeRelationListDataType)
	for _, relation := range relations.HvacSystemFunctionOperationModeRelationData {
		for _, operationModeId := range relation.OperationModeId {
			selector := &model.HvacOperationModeDescriptionListDataSelectorsType{
				OperationModeId: &operationModeId,
			}
			if _, err := hvac.RequestHvacOperationModeDescriptions(selector, nil); err != nil {
				logging.Log().Debug(err)
			}
		}
	}
}

func (e *CDSF) updateOverrunDescriptions(payload spineapi.EventPayload) {
	hvac, err := client.NewHvac(e.LocalEntity, payload.Entity)
	if err != nil {
		logging.Log().Debug(err)
		return
	}

	filter := model.HvacOverrunDescriptionDataType{
		OverrunType: util.Ptr(model.HvacOverrunTypeTypeOneTimeDhw),
	}
	descriptions, err := hvac.GetHvacOverrunDescriptionsForFilter(filter)
	if err != nil {
		logging.Log().Debug(err)
		return
	}
	if len(descriptions) != 1 {
		logging.Log().Errorf("Expected exactly one DHW overrun description for oneTimeDhw system function. Got %d", len(descriptions))
		return
	}

	// Store the overrun ID for the "oneTimeDhw" overrun type.
	e.overrunId = descriptions[0].OverrunId

	selector := &model.HvacOverrunListDataSelectorsType{
		OverrunId: e.overrunId,
	}
	if _, err := hvac.RequestHvacOverruns(selector, nil); err != nil {
		logging.Log().Debug(err)
	}
}

func (e *CDSF) updateOperationModes(payload spineapi.EventPayload) {
	hvac, err := client.NewHvac(e.LocalEntity, payload.Entity)
	if err != nil {
		logging.Log().Debug(err)
		return
	}

	relations, _ := hvac.GetHvacSystemFunctionOperationModeRelations()
	descriptions, _ := hvac.GetHvacOperationModeDescriptions()
	systemFunctions, _ := hvac.GetHvacSystemFunctions()

	if len(descriptions) == 0 || len(systemFunctions) == 0 {
		return
	}

	operationModeById := make(map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType)
	for _, description := range descriptions {
		operationModeById[*description.OperationModeId] = *description.OperationModeType
	}

	clear(e.operationModeByOperationModeId)
	for _, relation := range relations {
		for _, operationModeId := range relation.OperationModeId {
			if operationMode, found := operationModeById[operationModeId]; found {
				if _, found := e.operationModeByOperationModeId[operationModeId]; !found {
					// Store the operation mode by operation mode ID
					e.operationModeByOperationModeId[operationModeId] = operationMode
				}

				if _, found := e.operationModeIdByOperationMode[operationMode]; !found {
					// Store the operation mode ID by operation mode
					e.operationModeIdByOperationMode[operationMode] = operationModeId
				}
			}
		}
	}

	e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOperationModes)
}

func (e *CDSF) systemFunctionsDescriptionsUpdate(payload spineapi.EventPayload) {
	if hvac, err := client.NewHvac(e.LocalEntity, payload.Entity); err == nil {
		filter := model.HvacSystemFunctionDescriptionDataType{
			SystemFunctionType: util.Ptr(model.HvacSystemFunctionTypeTypeDhw),
		}
		descriptions, err := hvac.GetHvacSystemFunctionDescriptionsForFilter(filter)
		if err != nil {
			logging.Log().Debug(err)
			return
		}

		systemFunctionIds := make([]model.HvacSystemFunctionIdType, 0)
		for _, description := range descriptions {
			systemFunctionIds = append(systemFunctionIds, *description.SystemFunctionId)
		}
		if len(systemFunctionIds) != 1 {
			logging.Log().Errorf("Expected exactly one DHW system function description. Got %d", len(systemFunctionIds))
			return
		}

		// Store the system function ID for the DHW system function.
		e.dhwSystemFunctionId = &systemFunctionIds[0]

		// According to the usecase specification, we need to perform a partial read of hvacSystemFunctionOperationModeRelationListData
		// using systemFunctionId derived from hvacSystemFunctionDescriptionListData.
		systemFunctionOperationModeRelationsSelector := &model.HvacSystemFunctionOperationModeRelationListDataSelectorsType{
			SystemFunctionId: e.dhwSystemFunctionId,
		}
		if _, err := hvac.RequestHvacSystemFunctionOperationModeRelations(systemFunctionOperationModeRelationsSelector, nil); err != nil {
			logging.Log().Debug(err)
		}

		// According to the specification, we need to perform a partial read of hvacSystemFunctionListData
		// using systemFunctionId from hvacSystemFunctionDescriptionListData.
		systemFunctionsSelector := &model.HvacSystemFunctionListDataSelectorsType{
			SystemFunctionId: e.dhwSystemFunctionId,
		}
		if _, err := hvac.RequestHvacSystemFunctions(systemFunctionsSelector, nil); err != nil {
			logging.Log().Debug(err)
		}

		// The hvacOverrunDescriptionListData read SHOULD be a "full" read operation.
		_, err = hvac.RequestHvacOverrunDescriptions(nil, nil)
		if err != nil {
			logging.Log().Debug(err)
		}
	}
}

// dhwCircuitconnected handles the DHW circuit connected event.
func (e *CDSF) dhwCircuitconnected(entity spineapi.EntityRemoteInterface) {
	if hvac, err := client.NewHvac(e.LocalEntity, entity); err == nil {
		if !hvac.HasSubscription() {
			if _, err := hvac.Subscribe(); err != nil {
				logging.Log().Debug(err)
			}
		}

		if _, err := hvac.RequestHvacSystemFunctionDescriptions(nil, nil); err != nil {
			logging.Log().Debug(err)
		}
	}
}
