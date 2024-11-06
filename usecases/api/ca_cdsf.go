package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
)

type CaCDSFInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// Return the supported operation modes
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
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
	//   - entity: the entity to get the setpoints data from
	//
	// return values:
	//   - operationMode: The current operation mode
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	OperationMode(entity spineapi.EntityRemoteInterface) (*HvacOperationModeType, error)

	// write the operation mode
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//   - operationMode: The operation mode to write
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	WriteOperationMode(entity spineapi.EntityRemoteInterface, operationMode HvacOperationModeType) error

	// Scenario 2

	// Returns if an HVAC overrun is currently active.
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//
	// possible errors:
	//   - ErrDataNotAvailable if the operation modes are not (yet) available
	//   - and others
	IsOverrunActive(entity spineapi.EntityRemoteInterface) bool

	// Returns the current HVAC overrun status.
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//
	// possible errors:
	//   - ErrDataNotAvailable if the overrun status is not (yet) available
	//   - and others
	OverrunStatus(entity spineapi.EntityRemoteInterface) (*HvacOverrunStatusType, error)

	// Start Overrun
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//
	// possible errors:
	//   - ErrDataNotAvailable if the overrun status is not (yet) available
	//   - and others
	StartOverrun(entity spineapi.EntityRemoteInterface) error

	// Scenario 3

	// Stop Overrun
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//
	// possible errors:
	//   - ErrDataNotAvailable if the overrun status is not (yet) available
	//   - and others
	StopOverrun(entity spineapi.EntityRemoteInterface) error
}
