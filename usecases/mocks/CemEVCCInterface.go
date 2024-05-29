// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	api "github.com/enbility/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	spine_goapi "github.com/enbility/spine-go/api"
)

// CemEVCCInterface is an autogenerated mock type for the CemEVCCInterface type
type CemEVCCInterface struct {
	mock.Mock
}

type CemEVCCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemEVCCInterface) EXPECT() *CemEVCCInterface_Expecter {
	return &CemEVCCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemEVCCInterface) AddFeatures() {
	_m.Called()
}

// CemEVCCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemEVCCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemEVCCInterface_Expecter) AddFeatures() *CemEVCCInterface_AddFeatures_Call {
	return &CemEVCCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemEVCCInterface_AddFeatures_Call) Run(run func()) *CemEVCCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCCInterface_AddFeatures_Call) Return() *CemEVCCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCCInterface_AddFeatures_Call) RunAndReturn(run func()) *CemEVCCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemEVCCInterface) AddUseCase() {
	_m.Called()
}

// CemEVCCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemEVCCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemEVCCInterface_Expecter) AddUseCase() *CemEVCCInterface_AddUseCase_Call {
	return &CemEVCCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemEVCCInterface_AddUseCase_Call) Run(run func()) *CemEVCCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCCInterface_AddUseCase_Call) Return() *CemEVCCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCCInterface_AddUseCase_Call) RunAndReturn(run func()) *CemEVCCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AsymmetricChargingSupport provides a mock function with given fields: entity
func (_m *CemEVCCInterface) AsymmetricChargingSupport(entity spine_goapi.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for AsymmetricChargingSupport")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (bool, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_AsymmetricChargingSupport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AsymmetricChargingSupport'
type CemEVCCInterface_AsymmetricChargingSupport_Call struct {
	*mock.Call
}

// AsymmetricChargingSupport is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) AsymmetricChargingSupport(entity interface{}) *CemEVCCInterface_AsymmetricChargingSupport_Call {
	return &CemEVCCInterface_AsymmetricChargingSupport_Call{Call: _e.mock.On("AsymmetricChargingSupport", entity)}
}

func (_c *CemEVCCInterface_AsymmetricChargingSupport_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_AsymmetricChargingSupport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_AsymmetricChargingSupport_Call) Return(_a0 bool, _a1 error) *CemEVCCInterface_AsymmetricChargingSupport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_AsymmetricChargingSupport_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (bool, error)) *CemEVCCInterface_AsymmetricChargingSupport_Call {
	_c.Call.Return(run)
	return _c
}

// ChargeState provides a mock function with given fields: entity
func (_m *CemEVCCInterface) ChargeState(entity spine_goapi.EntityRemoteInterface) (api.EVChargeStateType, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ChargeState")
	}

	var r0 api.EVChargeStateType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (api.EVChargeStateType, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) api.EVChargeStateType); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(api.EVChargeStateType)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_ChargeState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChargeState'
type CemEVCCInterface_ChargeState_Call struct {
	*mock.Call
}

// ChargeState is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) ChargeState(entity interface{}) *CemEVCCInterface_ChargeState_Call {
	return &CemEVCCInterface_ChargeState_Call{Call: _e.mock.On("ChargeState", entity)}
}

func (_c *CemEVCCInterface_ChargeState_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_ChargeState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_ChargeState_Call) Return(_a0 api.EVChargeStateType, _a1 error) *CemEVCCInterface_ChargeState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_ChargeState_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (api.EVChargeStateType, error)) *CemEVCCInterface_ChargeState_Call {
	_c.Call.Return(run)
	return _c
}

// ChargingPowerLimits provides a mock function with given fields: entity
func (_m *CemEVCCInterface) ChargingPowerLimits(entity spine_goapi.EntityRemoteInterface) (float64, float64, float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ChargingPowerLimits")
	}

	var r0 float64
	var r1 float64
	var r2 float64
	var r3 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, float64, float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Get(1).(float64)
	}

	if rf, ok := ret.Get(2).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r2 = rf(entity)
	} else {
		r2 = ret.Get(2).(float64)
	}

	if rf, ok := ret.Get(3).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r3 = rf(entity)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// CemEVCCInterface_ChargingPowerLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChargingPowerLimits'
type CemEVCCInterface_ChargingPowerLimits_Call struct {
	*mock.Call
}

// ChargingPowerLimits is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) ChargingPowerLimits(entity interface{}) *CemEVCCInterface_ChargingPowerLimits_Call {
	return &CemEVCCInterface_ChargingPowerLimits_Call{Call: _e.mock.On("ChargingPowerLimits", entity)}
}

func (_c *CemEVCCInterface_ChargingPowerLimits_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_ChargingPowerLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_ChargingPowerLimits_Call) Return(_a0 float64, _a1 float64, _a2 float64, _a3 error) *CemEVCCInterface_ChargingPowerLimits_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *CemEVCCInterface_ChargingPowerLimits_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, float64, float64, error)) *CemEVCCInterface_ChargingPowerLimits_Call {
	_c.Call.Return(run)
	return _c
}

// CommunicationStandard provides a mock function with given fields: entity
func (_m *CemEVCCInterface) CommunicationStandard(entity spine_goapi.EntityRemoteInterface) (model.DeviceConfigurationKeyValueStringType, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for CommunicationStandard")
	}

	var r0 model.DeviceConfigurationKeyValueStringType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (model.DeviceConfigurationKeyValueStringType, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) model.DeviceConfigurationKeyValueStringType); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(model.DeviceConfigurationKeyValueStringType)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_CommunicationStandard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommunicationStandard'
type CemEVCCInterface_CommunicationStandard_Call struct {
	*mock.Call
}

// CommunicationStandard is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) CommunicationStandard(entity interface{}) *CemEVCCInterface_CommunicationStandard_Call {
	return &CemEVCCInterface_CommunicationStandard_Call{Call: _e.mock.On("CommunicationStandard", entity)}
}

func (_c *CemEVCCInterface_CommunicationStandard_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_CommunicationStandard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_CommunicationStandard_Call) Return(_a0 model.DeviceConfigurationKeyValueStringType, _a1 error) *CemEVCCInterface_CommunicationStandard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_CommunicationStandard_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (model.DeviceConfigurationKeyValueStringType, error)) *CemEVCCInterface_CommunicationStandard_Call {
	_c.Call.Return(run)
	return _c
}

// EVConnected provides a mock function with given fields: entity
func (_m *CemEVCCInterface) EVConnected(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EVConnected")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVCCInterface_EVConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EVConnected'
type CemEVCCInterface_EVConnected_Call struct {
	*mock.Call
}

// EVConnected is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) EVConnected(entity interface{}) *CemEVCCInterface_EVConnected_Call {
	return &CemEVCCInterface_EVConnected_Call{Call: _e.mock.On("EVConnected", entity)}
}

func (_c *CemEVCCInterface_EVConnected_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_EVConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_EVConnected_Call) Return(_a0 bool) *CemEVCCInterface_EVConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCCInterface_EVConnected_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemEVCCInterface_EVConnected_Call {
	_c.Call.Return(run)
	return _c
}

// Identifications provides a mock function with given fields: entity
func (_m *CemEVCCInterface) Identifications(entity spine_goapi.EntityRemoteInterface) ([]api.IdentificationItem, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for Identifications")
	}

	var r0 []api.IdentificationItem
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) ([]api.IdentificationItem, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []api.IdentificationItem); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.IdentificationItem)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_Identifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Identifications'
type CemEVCCInterface_Identifications_Call struct {
	*mock.Call
}

// Identifications is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) Identifications(entity interface{}) *CemEVCCInterface_Identifications_Call {
	return &CemEVCCInterface_Identifications_Call{Call: _e.mock.On("Identifications", entity)}
}

func (_c *CemEVCCInterface_Identifications_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_Identifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_Identifications_Call) Return(_a0 []api.IdentificationItem, _a1 error) *CemEVCCInterface_Identifications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_Identifications_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) ([]api.IdentificationItem, error)) *CemEVCCInterface_Identifications_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntity provides a mock function with given fields: entity
func (_m *CemEVCCInterface) IsCompatibleEntity(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVCCInterface_IsCompatibleEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntity'
type CemEVCCInterface_IsCompatibleEntity_Call struct {
	*mock.Call
}

// IsCompatibleEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) IsCompatibleEntity(entity interface{}) *CemEVCCInterface_IsCompatibleEntity_Call {
	return &CemEVCCInterface_IsCompatibleEntity_Call{Call: _e.mock.On("IsCompatibleEntity", entity)}
}

func (_c *CemEVCCInterface_IsCompatibleEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_IsCompatibleEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_IsCompatibleEntity_Call) Return(_a0 bool) *CemEVCCInterface_IsCompatibleEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCCInterface_IsCompatibleEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemEVCCInterface_IsCompatibleEntity_Call {
	_c.Call.Return(run)
	return _c
}

// IsInSleepMode provides a mock function with given fields: entity
func (_m *CemEVCCInterface) IsInSleepMode(entity spine_goapi.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsInSleepMode")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (bool, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_IsInSleepMode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsInSleepMode'
type CemEVCCInterface_IsInSleepMode_Call struct {
	*mock.Call
}

// IsInSleepMode is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) IsInSleepMode(entity interface{}) *CemEVCCInterface_IsInSleepMode_Call {
	return &CemEVCCInterface_IsInSleepMode_Call{Call: _e.mock.On("IsInSleepMode", entity)}
}

func (_c *CemEVCCInterface_IsInSleepMode_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_IsInSleepMode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_IsInSleepMode_Call) Return(_a0 bool, _a1 error) *CemEVCCInterface_IsInSleepMode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_IsInSleepMode_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (bool, error)) *CemEVCCInterface_IsInSleepMode_Call {
	_c.Call.Return(run)
	return _c
}

// IsUseCaseSupported provides a mock function with given fields: remoteEntity
func (_m *CemEVCCInterface) IsUseCaseSupported(remoteEntity spine_goapi.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(remoteEntity)

	if len(ret) == 0 {
		panic("no return value specified for IsUseCaseSupported")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (bool, error)); ok {
		return rf(remoteEntity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(remoteEntity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(remoteEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_IsUseCaseSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUseCaseSupported'
type CemEVCCInterface_IsUseCaseSupported_Call struct {
	*mock.Call
}

// IsUseCaseSupported is a helper method to define mock.On call
//   - remoteEntity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) IsUseCaseSupported(remoteEntity interface{}) *CemEVCCInterface_IsUseCaseSupported_Call {
	return &CemEVCCInterface_IsUseCaseSupported_Call{Call: _e.mock.On("IsUseCaseSupported", remoteEntity)}
}

func (_c *CemEVCCInterface_IsUseCaseSupported_Call) Run(run func(remoteEntity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_IsUseCaseSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_IsUseCaseSupported_Call) Return(_a0 bool, _a1 error) *CemEVCCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_IsUseCaseSupported_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (bool, error)) *CemEVCCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(run)
	return _c
}

// ManufacturerData provides a mock function with given fields: entity
func (_m *CemEVCCInterface) ManufacturerData(entity spine_goapi.EntityRemoteInterface) (eebus_goapi.ManufacturerData, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ManufacturerData")
	}

	var r0 eebus_goapi.ManufacturerData
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (eebus_goapi.ManufacturerData, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) eebus_goapi.ManufacturerData); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(eebus_goapi.ManufacturerData)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCCInterface_ManufacturerData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ManufacturerData'
type CemEVCCInterface_ManufacturerData_Call struct {
	*mock.Call
}

// ManufacturerData is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCCInterface_Expecter) ManufacturerData(entity interface{}) *CemEVCCInterface_ManufacturerData_Call {
	return &CemEVCCInterface_ManufacturerData_Call{Call: _e.mock.On("ManufacturerData", entity)}
}

func (_c *CemEVCCInterface_ManufacturerData_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCCInterface_ManufacturerData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCCInterface_ManufacturerData_Call) Return(_a0 eebus_goapi.ManufacturerData, _a1 error) *CemEVCCInterface_ManufacturerData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCCInterface_ManufacturerData_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (eebus_goapi.ManufacturerData, error)) *CemEVCCInterface_ManufacturerData_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemEVCCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemEVCCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemEVCCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemEVCCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemEVCCInterface_UpdateUseCaseAvailability_Call {
	return &CemEVCCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemEVCCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemEVCCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemEVCCInterface_UpdateUseCaseAvailability_Call) Return() *CemEVCCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemEVCCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemEVCCInterface creates a new instance of CemEVCCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemEVCCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemEVCCInterface {
	mock := &CemEVCCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}