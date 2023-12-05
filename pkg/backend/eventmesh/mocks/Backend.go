// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	cleaner "github.com/kyma-project/eventing-manager/pkg/backend/cleaner"
	env "github.com/kyma-project/eventing-manager/pkg/env"

	mock "github.com/stretchr/testify/mock"

	v1alpha2 "github.com/kyma-project/eventing-manager/api/eventing/v1alpha2"

	v1beta1 "github.com/kyma-project/api-gateway/apis/gateway/v1beta1"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// DeleteSubscription provides a mock function with given fields: subscription
func (_m *Backend) DeleteSubscription(subscription *v1alpha2.Subscription) error {
	ret := _m.Called(subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.Subscription) error); ok {
		r0 = rf(subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: cfg
func (_m *Backend) Initialize(cfg env.Config) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(env.Config) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SyncSubscription provides a mock function with given fields: subscription, _a1, apiRule
func (_m *Backend) SyncSubscription(subscription *v1alpha2.Subscription, _a1 cleaner.Cleaner, apiRule *v1beta1.APIRule) (bool, error) {
	ret := _m.Called(subscription, _a1, apiRule)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*v1alpha2.Subscription, cleaner.Cleaner, *v1beta1.APIRule) bool); ok {
		r0 = rf(subscription, _a1, apiRule)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha2.Subscription, cleaner.Cleaner, *v1beta1.APIRule) error); ok {
		r1 = rf(subscription, _a1, apiRule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewBackend creates a new instance of Backend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBackend(t mockConstructorTestingTNewBackend) *Backend {
	mock := &Backend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
