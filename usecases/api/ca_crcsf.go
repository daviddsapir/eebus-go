package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
)

type CaCRCSFInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// Return the supported operation modes
	//
	// parameters:
	//   - entity: the entity to get the supported operation modes from
	//
	// return values:
	//   - operationModes: A list of the supported operation modes
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	OperationModes(entity spineapi.EntityRemoteInterface) ([]HvacOperationModeType, error)

	// Return the current operation mode
	//
	// parameters:
	//   - entity: the entity to get the current operation mode from
	//
	// return values:
	//   - operationMode: The current operation mode
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	OperationMode(entity spineapi.EntityRemoteInterface) (*HvacOperationModeType, error)

	// Write the operation mode
	//
	// parameters:
	//   - entity: the entity to write the operation mode to
	//   - operationMode: The operation mode to write
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	WriteOperationMode(entity spineapi.EntityRemoteInterface, operationMode HvacOperationModeType) error
}
