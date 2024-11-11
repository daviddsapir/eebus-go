package crcsf

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
// Possible errors:
//   - ErrDataNotAvailable: If the supported operation modes are not (yet) available.
//   - Other errors: Any other errors encountered during the process.
func (c *CRHSF) OperationModes(
	entity spineapi.EntityRemoteInterface,
) ([]usecasesapi.HvacOperationModeType, error) {
	if c.modeForModeId == nil {
		return nil, api.ErrDataNotAvailable
	}

	modes := []usecasesapi.HvacOperationModeType{}
	for _, mode := range c.modeForModeId {
		modes = append(modes, usecasesapi.HvacOperationModeType(mode))
	}

	return modes, nil
}

// OperationMode returns the current operation mode
//
// Parameters:
//   - entity: The entity to get the operation mode for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the operation mode is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (c *CRHSF) OperationMode(entity spineapi.EntityRemoteInterface) (*usecasesapi.HvacOperationModeType, error) {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	systemFunction, err := hvac.GetHvacSystemFunctionForId(*c.heatingSystemFunctionID)
	if err != nil {
		logging.Log().Debug(err)
		return nil, err
	}

	modeId := systemFunction.CurrentOperationModeId
	if modeId == nil {
		return nil, api.ErrDataNotAvailable
	}

	mode, found := c.modeForModeId[*modeId]
	if !found {
		return nil, api.ErrDataNotAvailable
	}

	return util.Ptr(usecasesapi.HvacOperationModeType(mode)), nil
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
func (c *CRHSF) WriteOperationMode(
	entity spineapi.EntityRemoteInterface,
	operationMode usecasesapi.HvacOperationModeType,
) error {
	hvac, err := client.NewHvac(c.LocalEntity, entity)
	if err != nil {
		return err
	}

	modeId, found := c.modeIdForMode[model.HvacOperationModeTypeType(operationMode)]
	if !found {
		return api.ErrDataNotAvailable
	}

	if c.heatingSystemFunctionID == nil {
		return api.ErrDataNotAvailable
	}

	systemFunction, err := hvac.GetHvacSystemFunctionForId(*c.heatingSystemFunctionID)
	if err != nil || systemFunction == nil {
		return api.ErrDataNotAvailable
	}

	systemFunction.SystemFunctionId = c.heatingSystemFunctionID
	systemFunction.CurrentOperationModeId = &modeId

	data := []model.HvacSystemFunctionDataType{*systemFunction}

	_, err = hvac.WriteHvacSystemFunctions(data)

	return err
}
