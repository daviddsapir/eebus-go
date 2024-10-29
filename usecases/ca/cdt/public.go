package cdt

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	usecasesapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/ship-go/logging"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// Setpoints retrieves the setpoints for various HVAC operation modes from a remote entity.
//
// Possible errors:
//   - ErrDataNotAvailable: If the mapping of operation modes to setpoints or the setpoints themselves are not available.
//   - Other errors: Any other errors encountered during the process.
func (e *CDT) Setpoints(entity spineapi.EntityRemoteInterface) ([]usecasesapi.Setpoint, error) {
	setpoints := make([]usecasesapi.Setpoint, 0)

	sp, err := client.NewSetpoint(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	for _, setpoint := range sp.GetSetpoints() {
		var value float64 = 0
		var minValue float64 = 0
		var maxValue float64 = 0
		var timePeriod model.TimePeriodType = model.TimePeriodType{}

		if setpoint.SetpointId == nil {
			logging.Log().Error("Setpoint ID is nil")
			continue
		}

		if setpoint.Value != nil {
			value = setpoint.Value.GetValue()
		}

		if setpoint.ValueMax != nil {
			maxValue = setpoint.ValueMax.GetValue()
		}

		if setpoint.ValueMin != nil {
			minValue = setpoint.ValueMin.GetValue()
		}

		if setpoint.TimePeriod != nil {
			timePeriod = *setpoint.TimePeriod
		}

		isActive := (setpoint.IsSetpointActive == nil || *setpoint.IsSetpointActive)
		isChangeable := (setpoint.IsSetpointChangeable == nil || *setpoint.IsSetpointChangeable)

		setpoints = append(setpoints,
			usecasesapi.Setpoint{
				Id:           uint(*setpoint.SetpointId),
				Value:        value,
				MinValue:     minValue,
				MaxValue:     maxValue,
				IsActive:     isActive,
				IsChangeable: isChangeable,
				TimePeriod:   timePeriod,
			},
		)
	}

	if len(setpoints) == 0 {
		return nil, api.ErrDataNotAvailable
	}

	return setpoints, nil
}

// SetpointConstraints retrieves the setpoint constraints for various HVAC operation modes from a remote entity.
//
// Possible errors:
//   - ErrDataNotAvailable: If the mapping of operation modes to setpoints or the setpoint constraints are not available.
//   - Other errors: Any other errors encountered during the process.
func (e *CDT) SetpointConstraints(entity spineapi.EntityRemoteInterface) ([]usecasesapi.SetpointConstraints, error) {
	setpointConstraints := make([]usecasesapi.SetpointConstraints, 0)

	sp, err := client.NewSetpoint(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	for _, constraints := range sp.GetSetpointConstraints() {
		var minValue float64 = 0
		var maxValue float64 = 0
		var setSize float64 = 0

		if constraints.SetpointId == nil {
			logging.Log().Error("Setpoint ID is nil")
			continue
		}

		if constraints.SetpointRangeMin != nil {
			minValue = constraints.SetpointRangeMin.GetValue()
		}

		if constraints.SetpointRangeMax != nil {
			maxValue = constraints.SetpointRangeMax.GetValue()
		}

		if constraints.SetpointStepSize != nil {
			setSize = constraints.SetpointStepSize.GetValue()
		}

		setpointConstraints = append(setpointConstraints,
			usecasesapi.SetpointConstraints{
				Id:       uint(*constraints.SetpointId),
				MinValue: minValue,
				MaxValue: maxValue,
				StepSize: setSize,
			},
		)
	}

	if len(setpointConstraints) == 0 {
		return nil, api.ErrDataNotAvailable
	}

	return setpointConstraints, nil
}

// WriteSetpoint sets the temperature setpoint for a specified HVAC operation mode on a remote entity.
//
// Possible errors:
//   - ErrDataNotAvailable: If the mapping of operation modes to setpoints is not available.
//   - ErrNotSupported: If the setpoint is not changeable.
//   - Other errors: Any other errors encountered during the process.
func (e *CDT) WriteSetpoint(
	entity spineapi.EntityRemoteInterface,
	mode model.HvacOperationModeTypeType,
	degC float64,
) error {
	if mode == model.HvacOperationModeTypeTypeAuto {
		return nil
	}

	setpointId, found := e.operationModeToSetpoint[mode]
	if !found {
		return api.ErrDataNotAvailable
	}

	setPoint, err := client.NewSetpoint(e.LocalEntity, entity)
	if err != nil {
		return err
	}

	setpointToWrite := []model.SetpointDataType{
		{
			SetpointId: &setpointId,
			Value:      model.NewScaledNumberType(degC),
		},
	}

	if _, err = setPoint.WriteSetPointListData(setpointToWrite); err != nil {
		return err
	}

	return nil
}
