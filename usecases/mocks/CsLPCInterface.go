// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	api "github.com/enbility/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	spine_goapi "github.com/enbility/spine-go/api"

	time "time"
)

// CsLPCInterface is an autogenerated mock type for the CsLPCInterface type
type CsLPCInterface struct {
	mock.Mock
}

type CsLPCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CsLPCInterface) EXPECT() *CsLPCInterface_Expecter {
	return &CsLPCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CsLPCInterface) AddFeatures() {
	_m.Called()
}

// CsLPCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CsLPCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) AddFeatures() *CsLPCInterface_AddFeatures_Call {
	return &CsLPCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CsLPCInterface_AddFeatures_Call) Run(run func()) *CsLPCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_AddFeatures_Call) Return() *CsLPCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_AddFeatures_Call) RunAndReturn(run func()) *CsLPCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CsLPCInterface) AddUseCase() {
	_m.Called()
}

// CsLPCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CsLPCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) AddUseCase() *CsLPCInterface_AddUseCase_Call {
	return &CsLPCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CsLPCInterface_AddUseCase_Call) Run(run func()) *CsLPCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_AddUseCase_Call) Return() *CsLPCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_AddUseCase_Call) RunAndReturn(run func()) *CsLPCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// ApproveOrDenyConsumptionLimit provides a mock function with given fields: msgCounter, approve, reason
func (_m *CsLPCInterface) ApproveOrDenyConsumptionLimit(msgCounter model.MsgCounterType, approve bool, reason string) {
	_m.Called(msgCounter, approve, reason)
}

// CsLPCInterface_ApproveOrDenyConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApproveOrDenyConsumptionLimit'
type CsLPCInterface_ApproveOrDenyConsumptionLimit_Call struct {
	*mock.Call
}

// ApproveOrDenyConsumptionLimit is a helper method to define mock.On call
//   - msgCounter model.MsgCounterType
//   - approve bool
//   - reason string
func (_e *CsLPCInterface_Expecter) ApproveOrDenyConsumptionLimit(msgCounter interface{}, approve interface{}, reason interface{}) *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call {
	return &CsLPCInterface_ApproveOrDenyConsumptionLimit_Call{Call: _e.mock.On("ApproveOrDenyConsumptionLimit", msgCounter, approve, reason)}
}

func (_c *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call) Run(run func(msgCounter model.MsgCounterType, approve bool, reason string)) *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MsgCounterType), args[1].(bool), args[2].(string))
	})
	return _c
}

func (_c *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call) Return() *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call) RunAndReturn(run func(model.MsgCounterType, bool, string)) *CsLPCInterface_ApproveOrDenyConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CsLPCInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
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

// CsLPCInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CsLPCInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CsLPCInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CsLPCInterface_AvailableScenariosForEntity_Call {
	return &CsLPCInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CsLPCInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CsLPCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CsLPCInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CsLPCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CsLPCInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// ConsumptionLimit provides a mock function with given fields:
func (_m *CsLPCInterface) ConsumptionLimit() (api.LoadLimit, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConsumptionLimit")
	}

	var r0 api.LoadLimit
	var r1 error
	if rf, ok := ret.Get(0).(func() (api.LoadLimit, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() api.LoadLimit); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.LoadLimit)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CsLPCInterface_ConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumptionLimit'
type CsLPCInterface_ConsumptionLimit_Call struct {
	*mock.Call
}

// ConsumptionLimit is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) ConsumptionLimit() *CsLPCInterface_ConsumptionLimit_Call {
	return &CsLPCInterface_ConsumptionLimit_Call{Call: _e.mock.On("ConsumptionLimit")}
}

func (_c *CsLPCInterface_ConsumptionLimit_Call) Run(run func()) *CsLPCInterface_ConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_ConsumptionLimit_Call) Return(_a0 api.LoadLimit, _a1 error) *CsLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CsLPCInterface_ConsumptionLimit_Call) RunAndReturn(run func() (api.LoadLimit, error)) *CsLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// ConsumptionNominalMax provides a mock function with given fields:
func (_m *CsLPCInterface) ConsumptionNominalMax() (float64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConsumptionNominalMax")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func() (float64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CsLPCInterface_ConsumptionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumptionNominalMax'
type CsLPCInterface_ConsumptionNominalMax_Call struct {
	*mock.Call
}

// ConsumptionNominalMax is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) ConsumptionNominalMax() *CsLPCInterface_ConsumptionNominalMax_Call {
	return &CsLPCInterface_ConsumptionNominalMax_Call{Call: _e.mock.On("ConsumptionNominalMax")}
}

func (_c *CsLPCInterface_ConsumptionNominalMax_Call) Run(run func()) *CsLPCInterface_ConsumptionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_ConsumptionNominalMax_Call) Return(_a0 float64, _a1 error) *CsLPCInterface_ConsumptionNominalMax_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CsLPCInterface_ConsumptionNominalMax_Call) RunAndReturn(run func() (float64, error)) *CsLPCInterface_ConsumptionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeConsumptionActivePowerLimit provides a mock function with given fields:
func (_m *CsLPCInterface) FailsafeConsumptionActivePowerLimit() (float64, bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FailsafeConsumptionActivePowerLimit")
	}

	var r0 float64
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() (float64, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeConsumptionActivePowerLimit'
type CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// FailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) FailsafeConsumptionActivePowerLimit() *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	return &CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("FailsafeConsumptionActivePowerLimit")}
}

func (_c *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Run(run func()) *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Return(value float64, isChangeable bool, resultErr error) *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(value, isChangeable, resultErr)
	return _c
}

func (_c *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func() (float64, bool, error)) *CsLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeDurationMinimum provides a mock function with given fields:
func (_m *CsLPCInterface) FailsafeDurationMinimum() (time.Duration, bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FailsafeDurationMinimum")
	}

	var r0 time.Duration
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() (time.Duration, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CsLPCInterface_FailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeDurationMinimum'
type CsLPCInterface_FailsafeDurationMinimum_Call struct {
	*mock.Call
}

// FailsafeDurationMinimum is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) FailsafeDurationMinimum() *CsLPCInterface_FailsafeDurationMinimum_Call {
	return &CsLPCInterface_FailsafeDurationMinimum_Call{Call: _e.mock.On("FailsafeDurationMinimum")}
}

func (_c *CsLPCInterface_FailsafeDurationMinimum_Call) Run(run func()) *CsLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_FailsafeDurationMinimum_Call) Return(duration time.Duration, isChangeable bool, resultErr error) *CsLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(duration, isChangeable, resultErr)
	return _c
}

func (_c *CsLPCInterface_FailsafeDurationMinimum_Call) RunAndReturn(run func() (time.Duration, bool, error)) *CsLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CsLPCInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
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

// CsLPCInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CsLPCInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CsLPCInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CsLPCInterface_IsCompatibleEntityType_Call {
	return &CsLPCInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CsLPCInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CsLPCInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CsLPCInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CsLPCInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CsLPCInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsHeartbeatWithinDuration provides a mock function with given fields:
func (_m *CsLPCInterface) IsHeartbeatWithinDuration() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsHeartbeatWithinDuration")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CsLPCInterface_IsHeartbeatWithinDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsHeartbeatWithinDuration'
type CsLPCInterface_IsHeartbeatWithinDuration_Call struct {
	*mock.Call
}

// IsHeartbeatWithinDuration is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) IsHeartbeatWithinDuration() *CsLPCInterface_IsHeartbeatWithinDuration_Call {
	return &CsLPCInterface_IsHeartbeatWithinDuration_Call{Call: _e.mock.On("IsHeartbeatWithinDuration")}
}

func (_c *CsLPCInterface_IsHeartbeatWithinDuration_Call) Run(run func()) *CsLPCInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_IsHeartbeatWithinDuration_Call) Return(_a0 bool) *CsLPCInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_IsHeartbeatWithinDuration_Call) RunAndReturn(run func() bool) *CsLPCInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CsLPCInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
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

// CsLPCInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CsLPCInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CsLPCInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CsLPCInterface_IsScenarioAvailableAtEntity_Call {
	return &CsLPCInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CsLPCInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CsLPCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CsLPCInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CsLPCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CsLPCInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// PendingConsumptionLimits provides a mock function with given fields:
func (_m *CsLPCInterface) PendingConsumptionLimits() map[model.MsgCounterType]api.LoadLimit {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PendingConsumptionLimits")
	}

	var r0 map[model.MsgCounterType]api.LoadLimit
	if rf, ok := ret.Get(0).(func() map[model.MsgCounterType]api.LoadLimit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.MsgCounterType]api.LoadLimit)
		}
	}

	return r0
}

// CsLPCInterface_PendingConsumptionLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingConsumptionLimits'
type CsLPCInterface_PendingConsumptionLimits_Call struct {
	*mock.Call
}

// PendingConsumptionLimits is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) PendingConsumptionLimits() *CsLPCInterface_PendingConsumptionLimits_Call {
	return &CsLPCInterface_PendingConsumptionLimits_Call{Call: _e.mock.On("PendingConsumptionLimits")}
}

func (_c *CsLPCInterface_PendingConsumptionLimits_Call) Run(run func()) *CsLPCInterface_PendingConsumptionLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_PendingConsumptionLimits_Call) Return(_a0 map[model.MsgCounterType]api.LoadLimit) *CsLPCInterface_PendingConsumptionLimits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_PendingConsumptionLimits_Call) RunAndReturn(run func() map[model.MsgCounterType]api.LoadLimit) *CsLPCInterface_PendingConsumptionLimits_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CsLPCInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
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

// CsLPCInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CsLPCInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) RemoteEntitiesScenarios() *CsLPCInterface_RemoteEntitiesScenarios_Call {
	return &CsLPCInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CsLPCInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CsLPCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CsLPCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPCInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CsLPCInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CsLPCInterface) RemoveUseCase() {
	_m.Called()
}

// CsLPCInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CsLPCInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) RemoveUseCase() *CsLPCInterface_RemoveUseCase_Call {
	return &CsLPCInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CsLPCInterface_RemoveUseCase_Call) Run(run func()) *CsLPCInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_RemoveUseCase_Call) Return() *CsLPCInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CsLPCInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// SetConsumptionLimit provides a mock function with given fields: limit
func (_m *CsLPCInterface) SetConsumptionLimit(limit api.LoadLimit) error {
	ret := _m.Called(limit)

	if len(ret) == 0 {
		panic("no return value specified for SetConsumptionLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.LoadLimit) error); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPCInterface_SetConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConsumptionLimit'
type CsLPCInterface_SetConsumptionLimit_Call struct {
	*mock.Call
}

// SetConsumptionLimit is a helper method to define mock.On call
//   - limit api.LoadLimit
func (_e *CsLPCInterface_Expecter) SetConsumptionLimit(limit interface{}) *CsLPCInterface_SetConsumptionLimit_Call {
	return &CsLPCInterface_SetConsumptionLimit_Call{Call: _e.mock.On("SetConsumptionLimit", limit)}
}

func (_c *CsLPCInterface_SetConsumptionLimit_Call) Run(run func(limit api.LoadLimit)) *CsLPCInterface_SetConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.LoadLimit))
	})
	return _c
}

func (_c *CsLPCInterface_SetConsumptionLimit_Call) Return(resultErr error) *CsLPCInterface_SetConsumptionLimit_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPCInterface_SetConsumptionLimit_Call) RunAndReturn(run func(api.LoadLimit) error) *CsLPCInterface_SetConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// SetConsumptionNominalMax provides a mock function with given fields: value
func (_m *CsLPCInterface) SetConsumptionNominalMax(value float64) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for SetConsumptionNominalMax")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPCInterface_SetConsumptionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConsumptionNominalMax'
type CsLPCInterface_SetConsumptionNominalMax_Call struct {
	*mock.Call
}

// SetConsumptionNominalMax is a helper method to define mock.On call
//   - value float64
func (_e *CsLPCInterface_Expecter) SetConsumptionNominalMax(value interface{}) *CsLPCInterface_SetConsumptionNominalMax_Call {
	return &CsLPCInterface_SetConsumptionNominalMax_Call{Call: _e.mock.On("SetConsumptionNominalMax", value)}
}

func (_c *CsLPCInterface_SetConsumptionNominalMax_Call) Run(run func(value float64)) *CsLPCInterface_SetConsumptionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *CsLPCInterface_SetConsumptionNominalMax_Call) Return(resultErr error) *CsLPCInterface_SetConsumptionNominalMax_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPCInterface_SetConsumptionNominalMax_Call) RunAndReturn(run func(float64) error) *CsLPCInterface_SetConsumptionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// SetFailsafeConsumptionActivePowerLimit provides a mock function with given fields: value, changeable
func (_m *CsLPCInterface) SetFailsafeConsumptionActivePowerLimit(value float64, changeable bool) error {
	ret := _m.Called(value, changeable)

	if len(ret) == 0 {
		panic("no return value specified for SetFailsafeConsumptionActivePowerLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(float64, bool) error); ok {
		r0 = rf(value, changeable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFailsafeConsumptionActivePowerLimit'
type CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// SetFailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
//   - value float64
//   - changeable bool
func (_e *CsLPCInterface_Expecter) SetFailsafeConsumptionActivePowerLimit(value interface{}, changeable interface{}) *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call {
	return &CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("SetFailsafeConsumptionActivePowerLimit", value, changeable)}
}

func (_c *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call) Run(run func(value float64, changeable bool)) *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(bool))
	})
	return _c
}

func (_c *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call) Return(resultErr error) *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func(float64, bool) error) *CsLPCInterface_SetFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// SetFailsafeDurationMinimum provides a mock function with given fields: duration, changeable
func (_m *CsLPCInterface) SetFailsafeDurationMinimum(duration time.Duration, changeable bool) error {
	ret := _m.Called(duration, changeable)

	if len(ret) == 0 {
		panic("no return value specified for SetFailsafeDurationMinimum")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, bool) error); ok {
		r0 = rf(duration, changeable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPCInterface_SetFailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFailsafeDurationMinimum'
type CsLPCInterface_SetFailsafeDurationMinimum_Call struct {
	*mock.Call
}

// SetFailsafeDurationMinimum is a helper method to define mock.On call
//   - duration time.Duration
//   - changeable bool
func (_e *CsLPCInterface_Expecter) SetFailsafeDurationMinimum(duration interface{}, changeable interface{}) *CsLPCInterface_SetFailsafeDurationMinimum_Call {
	return &CsLPCInterface_SetFailsafeDurationMinimum_Call{Call: _e.mock.On("SetFailsafeDurationMinimum", duration, changeable)}
}

func (_c *CsLPCInterface_SetFailsafeDurationMinimum_Call) Run(run func(duration time.Duration, changeable bool)) *CsLPCInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].(bool))
	})
	return _c
}

func (_c *CsLPCInterface_SetFailsafeDurationMinimum_Call) Return(resultErr error) *CsLPCInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPCInterface_SetFailsafeDurationMinimum_Call) RunAndReturn(run func(time.Duration, bool) error) *CsLPCInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// StartHeartbeat provides a mock function with given fields:
func (_m *CsLPCInterface) StartHeartbeat() {
	_m.Called()
}

// CsLPCInterface_StartHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartHeartbeat'
type CsLPCInterface_StartHeartbeat_Call struct {
	*mock.Call
}

// StartHeartbeat is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) StartHeartbeat() *CsLPCInterface_StartHeartbeat_Call {
	return &CsLPCInterface_StartHeartbeat_Call{Call: _e.mock.On("StartHeartbeat")}
}

func (_c *CsLPCInterface_StartHeartbeat_Call) Run(run func()) *CsLPCInterface_StartHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_StartHeartbeat_Call) Return() *CsLPCInterface_StartHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_StartHeartbeat_Call) RunAndReturn(run func()) *CsLPCInterface_StartHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// StopHeartbeat provides a mock function with given fields:
func (_m *CsLPCInterface) StopHeartbeat() {
	_m.Called()
}

// CsLPCInterface_StopHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopHeartbeat'
type CsLPCInterface_StopHeartbeat_Call struct {
	*mock.Call
}

// StopHeartbeat is a helper method to define mock.On call
func (_e *CsLPCInterface_Expecter) StopHeartbeat() *CsLPCInterface_StopHeartbeat_Call {
	return &CsLPCInterface_StopHeartbeat_Call{Call: _e.mock.On("StopHeartbeat")}
}

func (_c *CsLPCInterface_StopHeartbeat_Call) Run(run func()) *CsLPCInterface_StopHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPCInterface_StopHeartbeat_Call) Return() *CsLPCInterface_StopHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_StopHeartbeat_Call) RunAndReturn(run func()) *CsLPCInterface_StopHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CsLPCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CsLPCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CsLPCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CsLPCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CsLPCInterface_UpdateUseCaseAvailability_Call {
	return &CsLPCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CsLPCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CsLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CsLPCInterface_UpdateUseCaseAvailability_Call) Return() *CsLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CsLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCsLPCInterface creates a new instance of CsLPCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCsLPCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CsLPCInterface {
	mock := &CsLPCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
