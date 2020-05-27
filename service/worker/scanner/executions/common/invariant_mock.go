// The MIT License (MIT)
//
// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package common

import (
	mock "github.com/stretchr/testify/mock"
)

// MockInvariant is an autogenerated mock type for the Invariant type
type MockInvariant struct {
	mock.Mock
}

// Check provides a mock function with given fields: _a0, _a1
func (_m *MockInvariant) Check(_a0 Execution, _a1 *InvariantResourceBag) CheckResult {
	ret := _m.Called(_a0, _a1)

	var r0 CheckResult
	if rf, ok := ret.Get(0).(func(Execution, *InvariantResourceBag) CheckResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(CheckResult)
	}

	return r0
}

// Fix provides a mock function with given fields: _a0, _a1
func (_m *MockInvariant) Fix(_a0 Execution, _a1 *InvariantResourceBag) FixResult {
	ret := _m.Called(_a0, _a1)

	var r0 FixResult
	if rf, ok := ret.Get(0).(func(Execution, *InvariantResourceBag) FixResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(FixResult)
	}

	return r0
}

// InvariantType provides a mock function with given fields:
func (_m *MockInvariant) InvariantType() InvariantType {
	ret := _m.Called()

	var r0 InvariantType
	if rf, ok := ret.Get(0).(func() InvariantType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(InvariantType)
	}

	return r0
}