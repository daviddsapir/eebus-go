package cdsf

import "github.com/enbility/eebus-go/api"

const (
	// Update of the list of remote entities supporting the Use Case
	//
	// Use `RemoteEntities` to get the current data
	UseCaseSupportUpdate api.EventType = "ca-cdsf-UseCaseSupportUpdate"

	// Operation modes data updated
	//
	// Use `OperationModes` to get the current data
	//
	// Use Case CDSF, Scenario 1
	DataUpdateOperationModes api.EventType = "ca-cdsf-DataUpdateOperationModes"

	// Overrun status data updated
	//
	// Use `OverrunState` to get the current data
	//
	// Use Case CDSF, Scenario 2 and 3
	DataUpdateOverrunStatus api.EventType = "ca-cdsf-DataUpdateOverrunStatus"
)
