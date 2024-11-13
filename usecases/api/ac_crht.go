package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
)

// Actor: Customer Energy Management
// UseCase: Coordinated EV Charging
type CaCRHTInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// Returns the setpoints.
	//
	// parameters:
	//   - entity: the entity to get the setpoints data from
	//
	// return values:
	//   - setpoints: A list of setpoints
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	Setpoints(entity spineapi.EntityRemoteInterface) ([]Setpoint, error)

	// Returns the setpoint constraints.
	//
	// parameters:
	//   - entity: the entity to get the setpoints constraints from
	//
	// return values:
	//   - setpointConstraints: A list of setpoint constraints
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	SetpointConstraints(entity spineapi.EntityRemoteInterface) ([]SetpointConstraints, error)

	// Write a setpoint
	//
	// parameters:
	//   - entity: the entity to write the setpoint to
	//   - mode: the mode to write the setpoint for
	//   - temperature: the temperature setpoint value to write
	WriteSetpoint(entity spineapi.EntityRemoteInterface, mode HvacOperationModeType, temperature float64) error
}
