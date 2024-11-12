package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
)

// Actor: Monitoring Appliance
// UseCase: Monitoring of Outdoor Temperature
type MaMOTInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// return the momentary outdoor temperature
	//
	// parameters:
	//   - entity: the entity of the device (e.g. TemperatureSensor)
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	Temperature(entity spineapi.EntityRemoteInterface) (float64, error)
}
