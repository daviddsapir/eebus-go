// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	api "github.com/enbility/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	spine_goapi "github.com/enbility/spine-go/api"
)

// CemEVSECCInterface is an autogenerated mock type for the CemEVSECCInterface type
type CemEVSECCInterface struct {
	mock.Mock
}

type CemEVSECCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemEVSECCInterface) EXPECT() *CemEVSECCInterface_Expecter {
	return &CemEVSECCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemEVSECCInterface) AddFeatures() {
	_m.Called()
}

// CemEVSECCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemEVSECCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemEVSECCInterface_Expecter) AddFeatures() *CemEVSECCInterface_AddFeatures_Call {
	return &CemEVSECCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemEVSECCInterface_AddFeatures_Call) Run(run func()) *CemEVSECCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVSECCInterface_AddFeatures_Call) Return() *CemEVSECCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVSECCInterface_AddFeatures_Call) RunAndReturn(run func()) *CemEVSECCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemEVSECCInterface) AddUseCase() {
	_m.Called()
}

// CemEVSECCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemEVSECCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemEVSECCInterface_Expecter) AddUseCase() *CemEVSECCInterface_AddUseCase_Call {
	return &CemEVSECCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemEVSECCInterface_AddUseCase_Call) Run(run func()) *CemEVSECCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVSECCInterface_AddUseCase_Call) Return() *CemEVSECCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVSECCInterface_AddUseCase_Call) RunAndReturn(run func()) *CemEVSECCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CemEVSECCInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for AvailableScenariosForEntity")
	}

	var r0 []uint
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []uint); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	return r0
}

// CemEVSECCInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CemEVSECCInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVSECCInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CemEVSECCInterface_AvailableScenariosForEntity_Call {
	return &CemEVSECCInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CemEVSECCInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVSECCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVSECCInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CemEVSECCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVSECCInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CemEVSECCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CemEVSECCInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntityType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVSECCInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CemEVSECCInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVSECCInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CemEVSECCInterface_IsCompatibleEntityType_Call {
	return &CemEVSECCInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CemEVSECCInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVSECCInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVSECCInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CemEVSECCInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVSECCInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemEVSECCInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CemEVSECCInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
	ret := _m.Called(entity, scenario)

	if len(ret) == 0 {
		panic("no return value specified for IsScenarioAvailableAtEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, uint) bool); ok {
		r0 = rf(entity, scenario)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVSECCInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CemEVSECCInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CemEVSECCInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call {
	return &CemEVSECCInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CemEVSECCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// ManufacturerData provides a mock function with given fields: entity
func (_m *CemEVSECCInterface) ManufacturerData(entity spine_goapi.EntityRemoteInterface) (api.ManufacturerData, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ManufacturerData")
	}

	var r0 api.ManufacturerData
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (api.ManufacturerData, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) api.ManufacturerData); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(api.ManufacturerData)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVSECCInterface_ManufacturerData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ManufacturerData'
type CemEVSECCInterface_ManufacturerData_Call struct {
	*mock.Call
}

// ManufacturerData is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVSECCInterface_Expecter) ManufacturerData(entity interface{}) *CemEVSECCInterface_ManufacturerData_Call {
	return &CemEVSECCInterface_ManufacturerData_Call{Call: _e.mock.On("ManufacturerData", entity)}
}

func (_c *CemEVSECCInterface_ManufacturerData_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVSECCInterface_ManufacturerData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVSECCInterface_ManufacturerData_Call) Return(_a0 api.ManufacturerData, _a1 error) *CemEVSECCInterface_ManufacturerData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVSECCInterface_ManufacturerData_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (api.ManufacturerData, error)) *CemEVSECCInterface_ManufacturerData_Call {
	_c.Call.Return(run)
	return _c
}

// OperatingState provides a mock function with given fields: entity
func (_m *CemEVSECCInterface) OperatingState(entity spine_goapi.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for OperatingState")
	}

	var r0 model.DeviceDiagnosisOperatingStateType
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) model.DeviceDiagnosisOperatingStateType); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(model.DeviceDiagnosisOperatingStateType)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) string); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r2 = rf(entity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CemEVSECCInterface_OperatingState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OperatingState'
type CemEVSECCInterface_OperatingState_Call struct {
	*mock.Call
}

// OperatingState is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVSECCInterface_Expecter) OperatingState(entity interface{}) *CemEVSECCInterface_OperatingState_Call {
	return &CemEVSECCInterface_OperatingState_Call{Call: _e.mock.On("OperatingState", entity)}
}

func (_c *CemEVSECCInterface_OperatingState_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVSECCInterface_OperatingState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVSECCInterface_OperatingState_Call) Return(_a0 model.DeviceDiagnosisOperatingStateType, _a1 string, _a2 error) *CemEVSECCInterface_OperatingState_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CemEVSECCInterface_OperatingState_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error)) *CemEVSECCInterface_OperatingState_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CemEVSECCInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteEntitiesScenarios")
	}

	var r0 []eebus_goapi.RemoteEntityScenarios
	if rf, ok := ret.Get(0).(func() []eebus_goapi.RemoteEntityScenarios); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]eebus_goapi.RemoteEntityScenarios)
		}
	}

	return r0
}

// CemEVSECCInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CemEVSECCInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CemEVSECCInterface_Expecter) RemoteEntitiesScenarios() *CemEVSECCInterface_RemoteEntitiesScenarios_Call {
	return &CemEVSECCInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CemEVSECCInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CemEVSECCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVSECCInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CemEVSECCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVSECCInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CemEVSECCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CemEVSECCInterface) RemoveUseCase() {
	_m.Called()
}

// CemEVSECCInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CemEVSECCInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CemEVSECCInterface_Expecter) RemoveUseCase() *CemEVSECCInterface_RemoveUseCase_Call {
	return &CemEVSECCInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CemEVSECCInterface_RemoveUseCase_Call) Run(run func()) *CemEVSECCInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVSECCInterface_RemoveUseCase_Call) Return() *CemEVSECCInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVSECCInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CemEVSECCInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemEVSECCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemEVSECCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemEVSECCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemEVSECCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemEVSECCInterface_UpdateUseCaseAvailability_Call {
	return &CemEVSECCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemEVSECCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemEVSECCInterface_UpdateUseCaseAvailability_Call) Return() *CemEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVSECCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemEVSECCInterface creates a new instance of CemEVSECCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemEVSECCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemEVSECCInterface {
	mock := &CemEVSECCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
