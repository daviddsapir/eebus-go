package cdsf

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	usecasesapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/ship-go/logging"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// OperationModes returns the supported operation modes
//
// Parameters:
//   - entity: The entity to get the operation modes for.
//
// Returns:
//   - The supported operation modes.
//
// Possible errors:
//   - ErrDataNotAvailable: If the supported operation modes are not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) OperationModes(entity spineapi.EntityRemoteInterface) ([]usecasesapi.HvacOperationModeType, error) {
	if c.operationModeByOperationModeId == nil {
		return nil, api.ErrDataNotAvailable
	}

	operationModes := make([]usecasesapi.HvacOperationModeType, 0)
	for _, operationMode := range c.operationModeByOperationModeId {
		operationModes = append(operationModes, usecasesapi.HvacOperationModeType(operationMode))
	}

	return operationModes, nil
}

// OperationMode returns the current operation mode
//
// Parameters:
//   - entity: The entity to get the operation mode for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the operation mode is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) OperationMode(entity spineapi.EntityRemoteInterface) (*usecasesapi.HvacOperationModeType, error) {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	systemFunctionData, err := hvac.GetHvacSystemFunctionForId(*c.dhwSystemFunctionId)
	if err != nil {
		logging.Log().Debug(err)
		return nil, err
	}

	currentOperationModeId := systemFunctionData.CurrentOperationModeId
	if currentOperationModeId == nil {
		return nil, api.ErrDataNotAvailable
	}

	currentOperationMode, found := c.operationModeByOperationModeId[*currentOperationModeId]
	if !found {
		return nil, api.ErrDataNotAvailable
	}

	return util.Ptr(usecasesapi.HvacOperationModeType(currentOperationMode)), nil
}

// WriteOperationMode writes the operation mode
//
// Parameters:
//   - entity: The entity to write the operation mode for.
//   - operationMode: The operation mode to write.
//
// Possible errors:
//   - ErrDataNotAvailable: If the operation mode is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) WriteOperationMode(
	entity spineapi.EntityRemoteInterface,
	operationMode usecasesapi.HvacOperationModeType,
) error {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return err
	}

	operationModeId, found := c.operationModeIdByOperationMode[model.HvacOperationModeTypeType(operationMode)]
	if !found {
		return api.ErrDataNotAvailable
	}

	if c.dhwSystemFunctionId == nil {
		return api.ErrDataNotAvailable
	}

	systemFunctionData := &model.HvacSystemFunctionDataType{
		SystemFunctionId:       c.dhwSystemFunctionId,
		CurrentOperationModeId: &operationModeId,
	}

	data := []model.HvacSystemFunctionDataType{
		*systemFunctionData,
	}

	_, err = hvac.WriteHvacSystemFunctions(data)

	return err
}

// IsOverrunActive returns whether the overrun is active
//
// Parameters:
//   - entity: The entity to check the overrun status for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the overrun status is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) IsOverrunActive(entity spineapi.EntityRemoteInterface) bool {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return false
	}

	if c.dhwSystemFunctionId == nil {
		return false
	}

	systemFunctionData, err := hvac.GetHvacSystemFunctionForId(*c.dhwSystemFunctionId)
	if err == nil && systemFunctionData != nil && systemFunctionData.IsOverrunActive != nil {
		return *systemFunctionData.IsOverrunActive
	}

	return false
}

// OverrunStatus returns the overrun status
//
// Parameters:
//   - entity: The entity to get the overrun status for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the overrun status is not (yet) available.
//   - Other errors: Any other errors encountered during the process.
func (c *CDSF) OverrunStatus(
	entity spineapi.EntityRemoteInterface,
) (*usecasesapi.HvacOverrunStatusType, error) {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	if c.overrunId == nil {
		return nil, api.ErrDataNotAvailable
	}

	overrun, err := hvac.GetHvacOverrunForId(*c.overrunId)
	if err != nil {
		return nil, err
	}

	if overrun == nil || overrun.OverrunStatus == nil {
		return nil, api.ErrDataNotAvailable
	}

	return util.Ptr(usecasesapi.HvacOverrunStatusType(*overrun.OverrunStatus)), nil
}

// setOverrunState sets the overrun state
//
// Parameters:
//   - entity: The entity to set the overrun state for.
//   - state: The overrun state to set.
//
// Possible errors:
//   - ErrDataNotAvailable: If the overrun state is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) setOverrunState(
	entity spineapi.EntityRemoteInterface,
	state model.HvacOverrunStatusType,
) error {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return err
	}

	if c.overrunId == nil {
		return api.ErrDataNotAvailable
	}

	overrunData := model.HvacOverrunDataType{
		OverrunId:     c.overrunId,
		OverrunStatus: util.Ptr(model.HvacOverrunStatusType(state)),
	}

	data := []model.HvacOverrunDataType{
		overrunData,
	}

	_, err = hvac.WriteHvacOverruns(data)

	return err
}

// StartOverrun starts the overrun
//
// Parameters:
//   - entity: The entity to start the overrun for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the overrun state is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) StartOverrun(entity spineapi.EntityRemoteInterface) error {
	return c.setOverrunState(entity, model.HvacOverrunStatusTypeActive)
}

// StopOverrun stops the overrun
//
// Parameters:
//   - entity: The entity to stop the overrun for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the overrun state is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CDSF) StopOverrun(entity spineapi.EntityRemoteInterface) error {
	return c.setOverrunState(entity, model.HvacOverrunStatusTypeInactive)
}
