// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	event "sigs.k8s.io/controller-runtime/pkg/event"
)

// Watcher is an autogenerated mock type for the Watcher type
type Watcher struct {
	mock.Mock
}

type Watcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Watcher) EXPECT() *Watcher_Expecter {
	return &Watcher_Expecter{mock: &_m.Mock}
}

// GetEventsChannel provides a mock function with given fields:
func (_m *Watcher) GetEventsChannel() <-chan event.GenericEvent {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEventsChannel")
	}

	var r0 <-chan event.GenericEvent
	if rf, ok := ret.Get(0).(func() <-chan event.GenericEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan event.GenericEvent)
		}
	}

	return r0
}

// Watcher_GetEventsChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventsChannel'
type Watcher_GetEventsChannel_Call struct {
	*mock.Call
}

// GetEventsChannel is a helper method to define mock.On call
func (_e *Watcher_Expecter) GetEventsChannel() *Watcher_GetEventsChannel_Call {
	return &Watcher_GetEventsChannel_Call{Call: _e.mock.On("GetEventsChannel")}
}

func (_c *Watcher_GetEventsChannel_Call) Run(run func()) *Watcher_GetEventsChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Watcher_GetEventsChannel_Call) Return(_a0 <-chan event.GenericEvent) *Watcher_GetEventsChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Watcher_GetEventsChannel_Call) RunAndReturn(run func() <-chan event.GenericEvent) *Watcher_GetEventsChannel_Call {
	_c.Call.Return(run)
	return _c
}

// IsStarted provides a mock function with given fields:
func (_m *Watcher) IsStarted() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsStarted")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Watcher_IsStarted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsStarted'
type Watcher_IsStarted_Call struct {
	*mock.Call
}

// IsStarted is a helper method to define mock.On call
func (_e *Watcher_Expecter) IsStarted() *Watcher_IsStarted_Call {
	return &Watcher_IsStarted_Call{Call: _e.mock.On("IsStarted")}
}

func (_c *Watcher_IsStarted_Call) Run(run func()) *Watcher_IsStarted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Watcher_IsStarted_Call) Return(_a0 bool) *Watcher_IsStarted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Watcher_IsStarted_Call) RunAndReturn(run func() bool) *Watcher_IsStarted_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *Watcher) Start() {
	_m.Called()
}

// Watcher_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Watcher_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *Watcher_Expecter) Start() *Watcher_Start_Call {
	return &Watcher_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *Watcher_Start_Call) Run(run func()) *Watcher_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Watcher_Start_Call) Return() *Watcher_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *Watcher_Start_Call) RunAndReturn(run func()) *Watcher_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Watcher) Stop() {
	_m.Called()
}

// Watcher_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Watcher_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Watcher_Expecter) Stop() *Watcher_Stop_Call {
	return &Watcher_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Watcher_Stop_Call) Run(run func()) *Watcher_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Watcher_Stop_Call) Return() *Watcher_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Watcher_Stop_Call) RunAndReturn(run func()) *Watcher_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewWatcher creates a new instance of Watcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Watcher {
	mock := &Watcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
