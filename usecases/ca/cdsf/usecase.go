package cdsf

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

// Optimization of Heat Pump Compressor Function
type CDSF struct {
	*usecase.UseCaseBase

	dhwSystemFunctionId            *model.HvacSystemFunctionIdType
	oneTimeDhwOverrunId            *model.HvacOverrunIdType
	operationModeByOperationModeId map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType
	operationModeIdByOperationMode map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType
}

var _ ucapi.CaCDSFInterface = (*CDSF)(nil)

// NewCDT creates a new CDT use case
func NewCDSF(
	localEntity spineapi.EntityLocalInterface,
	eventCB api.EntityEventCallback,
) *CDSF {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeDHWCircuit,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeDHWCircuit,
	}
	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario:  model.UseCaseScenarioSupportType(1),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeHvac,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(2),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeHvac,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(3),
			Mandatory: false, // Recommended
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

	uc := &CDSF{
		UseCaseBase:                    usecase,
		operationModeByOperationModeId: make(map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType),
		operationModeIdByOperationMode: make(map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType),
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

// AddFeatures adds the features required for the CDT use case
func (e *CDSF) AddFeatures() {
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeHvac,
	}

	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
