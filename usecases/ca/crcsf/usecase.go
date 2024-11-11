package crcsf

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

// Optimization of Heat Pump Compressor Function
type CRCSF struct {
	*usecase.UseCaseBase

	coolingSystemFunctionID *model.HvacSystemFunctionIdType
	modeForModeId           map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType
	modeIdForMode           map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType
}

var _ ucapi.CaCRCSFInterface = (*CRCSF)(nil)

// NewCDT creates a new CRCSF use case
func NewCRCSF(
	localEntity spineapi.EntityLocalInterface,
	eventCB api.EntityEventCallback,
) *CRCSF {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeHVACRoom,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeHvacRoom,
		model.EntityTypeTypeHeatingCircuit,
		model.EntityTypeTypeHeatingZone,
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

	uc := &CRCSF{
		UseCaseBase:   usecase,
		modeForModeId: map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType{},
		modeIdForMode: map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType{},
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

// AddFeatures adds the features required for the CRCSF use case
func (e *CRCSF) AddFeatures() {
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeHvac,
	}

	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
