// Code generated by MockGen. DO NOT EDIT.
// Source: go.medium.engineering/picchu/pkg/plan (interfaces: Plan)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockPlan is a mock of Plan interface
type MockPlan struct {
	ctrl     *gomock.Controller
	recorder *MockPlanMockRecorder
}

// MockPlanMockRecorder is the mock recorder for MockPlan
type MockPlanMockRecorder struct {
	mock *MockPlan
}

// NewMockPlan creates a new mock instance
func NewMockPlan(ctrl *gomock.Controller) *MockPlan {
	mock := &MockPlan{ctrl: ctrl}
	mock.recorder = &MockPlanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlan) EXPECT() *MockPlanMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockPlan) Apply(arg0 context.Context, arg1 client.Client, arg2 *v1alpha1.Cluster, arg3 logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockPlanMockRecorder) Apply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockPlan)(nil).Apply), arg0, arg1, arg2, arg3)
}
