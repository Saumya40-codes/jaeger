// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	processor "github.com/jaegertracing/jaeger/cmd/ingester/app/processor"
	mock "github.com/stretchr/testify/mock"
)

// SpanProcessor is an autogenerated mock type for the SpanProcessor type
type SpanProcessor struct {
	mock.Mock
}

// Close provides a mock function with no fields
func (_m *SpanProcessor) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Process provides a mock function with given fields: input
func (_m *SpanProcessor) Process(input processor.Message) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Process")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(processor.Message) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSpanProcessor creates a new instance of SpanProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpanProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpanProcessor {
	mock := &SpanProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
