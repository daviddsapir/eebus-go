package mrhsf

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

// MDSF - Monitoring of DHW System Function Use Case
type MRHSF struct {
	*usecase.UseCaseBase

	heatingSystemFunctionID         *model.HvacSystemFunctionIdType
	operationModeForOperationModeId map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType
}

var _ ucapi.MaMRHSFInterface = (*MRHSF)(nil)

// Create a new Monitoring of DHW System Function Use Case
func NewMRHSF(
	localEntity spineapi.EntityLocalInterface,
	eventCB api.EntityEventCallback,
) *MRHSF {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeHVACRoom,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeHvacRoom,
	}
	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario:  model.UseCaseScenarioSupportType(1),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeHvac,
			},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeMonitoringAppliance,
		model.UseCaseNameTypeMonitoringOfRoomHeatingSystemFunction,
		"1.0.0",
		"release",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes,
	)

	uc := &MRHSF{
		UseCaseBase:                     usecase,
		operationModeForOperationModeId: make(map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType),
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

func (e *MRHSF) AddFeatures() {
	// client features
	var clientFeatures = []model.FeatureTypeType{}
	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
