package internal

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

type HvacCommon struct {
	featureLocal  spineapi.FeatureLocalInterface
	featureRemote spineapi.FeatureRemoteInterface
}

// NewLocalHvac creates a new HvacCommon helper for local entities
func NewLocalHvac(featureLocal spineapi.FeatureLocalInterface) *HvacCommon {
	return &HvacCommon{
		featureLocal: featureLocal,
	}
}

// NewRemoteHvac creates a new HvacCommon helper for remote entities
func NewRemoteHvac(featureRemote spineapi.FeatureRemoteInterface) *HvacCommon {
	return &HvacCommon{
		featureRemote: featureRemote,
	}
}

// GetHvacOperationModeDescriptions returns the operation mode descriptions
func (h *HvacCommon) GetHvacOperationModeDescriptions() ([]model.HvacOperationModeDescriptionDataType, error) {
	function := model.FunctionTypeHvacOperationModeDescriptionListData
	operationModeDescriptions := make([]model.HvacOperationModeDescriptionDataType, 0)

	data, err := featureDataCopyOfType[model.HvacOperationModeDescriptionListDataType](h.featureLocal, h.featureRemote, function)
	if err == nil || data != nil {
		operationModeDescriptions = append(operationModeDescriptions, data.HvacOperationModeDescriptionData...)
	}

	return operationModeDescriptions, nil
}

// GetHvacSystemFunctionSetpointRelations returns the operation mode relations (used to map operation modes to setpoints)
func (h *HvacCommon) GetHvacSystemFunctionSetpointRelations() ([]model.HvacSystemFunctionSetpointRelationDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionSetPointRelationListData
	relations := make([]model.HvacSystemFunctionSetpointRelationDataType, 0)

	data, err := featureDataCopyOfType[model.HvacSystemFunctionSetpointRelationListDataType](h.featureLocal, h.featureRemote, function)
	if err == nil || data != nil {
		relations = append(relations, data.HvacSystemFunctionSetpointRelationData...)
	}

	return relations, nil
}

// GetHvacSystemFunctionOperationModeRelations returns the operation mode relations
func (h *HvacCommon) GetHvacSystemFunctionOperationModeRelations() ([]model.HvacSystemFunctionOperationModeRelationDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionOperationModeRelationListData
	relations := make([]model.HvacSystemFunctionOperationModeRelationDataType, 0)

	data, err := featureDataCopyOfType[model.HvacSystemFunctionOperationModeRelationListDataType](h.featureLocal, h.featureRemote, function)
	if err == nil || data != nil {
		relations = append(relations, data.HvacSystemFunctionOperationModeRelationData...)
	}

	return relations, nil
}

// GetHvacSystemFunctionSetpointRelations returns the operation mode relations for a given system function id
func (h *HvacCommon) GetHvacSystemFunctionSetpointRelationsForSystemFunctionId(
	id model.HvacSystemFunctionIdType,
) ([]model.HvacSystemFunctionSetpointRelationDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionSetPointRelationListData
	filter := model.HvacSystemFunctionSetpointRelationDataType{
		SystemFunctionId: &id,
	}

	data, err := featureDataCopyOfType[model.HvacSystemFunctionSetpointRelationListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacSystemFunctionSetpointRelationData == nil {
		return nil, api.ErrDataNotAvailable
	}

	result := searchFilterInList[model.HvacSystemFunctionSetpointRelationDataType](data.HvacSystemFunctionSetpointRelationData, filter)
	if len(result) == 0 {
		return nil, api.ErrDataNotAvailable
	}

	return result, nil
}

// GetHvacSystemFunctionDescriptions returns the system function descriptions for a given filter
// GetHvacSystemFunctionDescriptionsForFilter returns the system function descriptions for a given filter
func (h *HvacCommon) GetHvacSystemFunctionDescriptionsForFilter(
	filter model.HvacSystemFunctionDescriptionDataType,
) ([]model.HvacSystemFunctionDescriptionDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionDescriptionListData

	data, err := featureDataCopyOfType[model.HvacSystemFunctionDescriptionListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacSystemFunctionDescriptionData == nil {
		return nil, api.ErrDataNotAvailable
	}

	descriptions := searchFilterInList[model.HvacSystemFunctionDescriptionDataType](data.HvacSystemFunctionDescriptionData, filter)

	return descriptions, nil
}

// GetHvacSystemFunctionForId returns the system function for a given ID
func (h *HvacCommon) GetHvacSystemFunctionForId(id model.HvacSystemFunctionIdType) (*model.HvacSystemFunctionDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionListData
	filter := model.HvacSystemFunctionDataType{
		SystemFunctionId: util.Ptr(id),
	}

	data, err := featureDataCopyOfType[model.HvacSystemFunctionListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacSystemFunctionData == nil {
		return nil, api.ErrDataNotAvailable
	}

	systemFunction := searchFilterInList[model.HvacSystemFunctionDataType](data.HvacSystemFunctionData, filter)

	return util.Ptr(systemFunction[0]), nil
}

// GetHvacSystemFunctions returns the system functions
func (h *HvacCommon) GetHvacSystemFunctions() ([]model.HvacSystemFunctionDataType, error) {
	function := model.FunctionTypeHvacSystemFunctionListData

	data, err := featureDataCopyOfType[model.HvacSystemFunctionListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacSystemFunctionData == nil {
		return nil, api.ErrDataNotAvailable
	}

	return data.HvacSystemFunctionData, nil
}

func (h *HvacCommon) GetHvacOverrunDescriptionsForFilter(
	filter model.HvacOverrunDescriptionDataType,
) ([]model.HvacOverrunDescriptionDataType, error) {
	function := model.FunctionTypeHvacOverrunDescriptionListData

	data, err := featureDataCopyOfType[model.HvacOverrunDescriptionListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacOverrunDescriptionData == nil {
		return nil, api.ErrDataNotAvailable
	}

	descriptions := searchFilterInList[model.HvacOverrunDescriptionDataType](data.HvacOverrunDescriptionData, filter)

	return descriptions, nil
}

func (h *HvacCommon) GetHvacOverrunForId(
	id model.HvacOverrunIdType,
) (*model.HvacOverrunDataType, error) {
	function := model.FunctionTypeHvacOverrunListData

	data, err := featureDataCopyOfType[model.HvacOverrunListDataType](h.featureLocal, h.featureRemote, function)
	if err != nil || data == nil || data.HvacOverrunData == nil {
		return nil, api.ErrDataNotAvailable
	}

	filter := model.HvacOverrunDataType{
		OverrunId: util.Ptr(id),
	}

	overruns := searchFilterInList[model.HvacOverrunDataType](data.HvacOverrunData, filter)
	if len(overruns) != 1 {
		return nil, api.ErrDataNotAvailable
	}

	return util.Ptr(overruns[0]), nil
}
