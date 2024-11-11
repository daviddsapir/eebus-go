package crcsf

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

type CRHSF struct {
	*usecase.UseCaseBase

	heatingSystemFunctionID *model.HvacSystemFunctionIdType
	modeForModeId           map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType
	modeIdForMode           map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType
}

var _ ucapi.CaCRHSFInterface = (*CRHSF)(nil)

// NewCDT creates a new CRHSF use case
func NewCRHSF(
	localEntity spineapi.EntityLocalInterface,
	eventCB api.EntityEventCallback,
) *CRHSF {
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

	uc := &CRHSF{
		UseCaseBase:   usecase,
		modeForModeId: map[model.HvacOperationModeIdType]model.HvacOperationModeTypeType{},
		modeIdForMode: map[model.HvacOperationModeTypeType]model.HvacOperationModeIdType{},
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

// AddFeatures adds the features required for the CRCSF use case
func (e *CRHSF) AddFeatures() {
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeHvac,
	}

	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
