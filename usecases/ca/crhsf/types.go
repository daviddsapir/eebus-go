package crcsf

import "github.com/enbility/eebus-go/api"

const (
	// Update of the list of remote entities supporting the Use Case
	//
	// Use `RemoteEntities` to get the current data
	UseCaseSupportUpdate api.EventType = "ca-crcsf-UseCaseSupportUpdate"

	// Operation modes data updated
	//
	// Use `OperationModes` to get the current data
	//
	// Use Case CRCSF, Scenario 1
	DataUpdateOperationModes api.EventType = "ca-crcsf-DataUpdateOperationModes"
)
