package mrhsf

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	ucsapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/ship-go/logging"
	"github.com/enbility/ship-go/util"
	spineapi "github.com/enbility/spine-go/api"
)

// OperationMode returns the current operation mode
//
// Parameters:
//   - entity: The entity to get the operation mode for.
//
// Possible errors:
//   - ErrDataNotAvailable: If the operation mode is not (yet) available.
//   - Other: Any other errors encountered during the process.
func (e *MRHSF) OperationMode(entity spineapi.EntityRemoteInterface) (*ucsapi.HvacOperationModeType, error) {
	hvac, err := client.NewHvac(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	systemFunction, err := hvac.GetHvacSystemFunctionForId(*e.heatingSystemFunctionID)
	if err != nil {
		logging.Log().Debug(err)
		return nil, err
	}

	modeId := systemFunction.CurrentOperationModeId
	if modeId == nil {
		return nil, api.ErrDataNotAvailable
	}

	mode, found := e.operationModeForOperationModeId[*modeId]
	if !found {
		return nil, api.ErrDataNotAvailable
	}

	return util.Ptr(ucsapi.HvacOperationModeType(mode)), nil
}
