package crht

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

type CRHT struct {
	*usecase.UseCaseBase

	setpointIdsForMode map[model.HvacOperationModeTypeType]model.SetpointIdType
}

var _ ucapi.CaCRHTInterface = (*CRHT)(nil)

// NewCDT creates a new CRHT use case
func NewCRHT(
	localEntity spineapi.EntityLocalInterface,
	eventCB api.EntityEventCallback,
) *CRHT {
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
				model.FeatureTypeTypeSetpoint,
			},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeConfigurationAppliance,
		model.UseCaseNameTypeConfigurationOfDhwTemperature,
		"1.0.0",
		"release",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes,
	)

	uc := &CRHT{
		UseCaseBase:        usecase,
		setpointIdsForMode: make(map[model.HvacOperationModeTypeType]model.SetpointIdType),
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

// AddFeatures adds the features required for the CRHT use case
func (e *CRHT) AddFeatures() {
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeHvac,
		model.FeatureTypeTypeSetpoint,
	}

	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
