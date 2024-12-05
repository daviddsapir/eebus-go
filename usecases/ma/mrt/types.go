package mrhsf

import "github.com/enbility/eebus-go/api"

const (
	// Update of the list of remote entities supporting the Use Case
	//
	// Use `RemoteEntities` to get the current data
	UseCaseSupportUpdate api.EventType = "ma-mrhsf-UseCaseSupportUpdate"

	// Current operation mode
	//
	// Use `OperationMode` to get the current data
	//
	// Use Case MRHSF, Scenario 1
	DataUpdateOperationMode api.EventType = "ma-mrhsf-DataUpdateOperationMode"
)
