// Code generated by MockGen. DO NOT EDIT.
// Source: go.medium.engineering/picchu/pkg/controller/releasemanager (interfaces: Deployment)

// Package releasemanager is a generated GoMock package.
package releasemanager

import (
	context "context"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1"
	reflect "reflect"
)

// MockDeployment is a mock of Deployment interface
type MockDeployment struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentMockRecorder
}

// MockDeploymentMockRecorder is the mock recorder for MockDeployment
type MockDeploymentMockRecorder struct {
	mock *MockDeployment
}

// NewMockDeployment creates a new mock instance
func NewMockDeployment(ctrl *gomock.Controller) *MockDeployment {
	mock := &MockDeployment{ctrl: ctrl}
	mock.recorder = &MockDeploymentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeployment) EXPECT() *MockDeploymentMockRecorder {
	return m.recorder
}

// currentPercent mocks base method
func (m *MockDeployment) currentPercent() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "currentPercent")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// currentPercent indicates an expected call of currentPercent
func (mr *MockDeploymentMockRecorder) currentPercent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "currentPercent", reflect.TypeOf((*MockDeployment)(nil).currentPercent))
}

// del mocks base method
func (m *MockDeployment) del(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "del", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// del indicates an expected call of del
func (mr *MockDeploymentMockRecorder) del(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "del", reflect.TypeOf((*MockDeployment)(nil).del), arg0)
}

// deleteCanaryRules mocks base method
func (m *MockDeployment) deleteCanaryRules(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "deleteCanaryRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteCanaryRules indicates an expected call of deleteCanaryRules
func (mr *MockDeploymentMockRecorder) deleteCanaryRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteCanaryRules", reflect.TypeOf((*MockDeployment)(nil).deleteCanaryRules), arg0)
}

// deleteSLIRules mocks base method
func (m *MockDeployment) deleteSLIRules(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "deleteSLIRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteSLIRules indicates an expected call of deleteSLIRules
func (mr *MockDeploymentMockRecorder) deleteSLIRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteSLIRules", reflect.TypeOf((*MockDeployment)(nil).deleteSLIRules), arg0)
}

// getExternalTestStatus mocks base method
func (m *MockDeployment) getExternalTestStatus() ExternalTestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getExternalTestStatus")
	ret0, _ := ret[0].(ExternalTestStatus)
	return ret0
}

// getExternalTestStatus indicates an expected call of getExternalTestStatus
func (mr *MockDeploymentMockRecorder) getExternalTestStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getExternalTestStatus", reflect.TypeOf((*MockDeployment)(nil).getExternalTestStatus))
}

// getLog mocks base method
func (m *MockDeployment) getLog() logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLog")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// getLog indicates an expected call of getLog
func (mr *MockDeploymentMockRecorder) getLog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLog", reflect.TypeOf((*MockDeployment)(nil).getLog))
}

// getStatus mocks base method
func (m *MockDeployment) getStatus() *v1alpha1.ReleaseManagerRevisionStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getStatus")
	ret0, _ := ret[0].(*v1alpha1.ReleaseManagerRevisionStatus)
	return ret0
}

// getStatus indicates an expected call of getStatus
func (mr *MockDeploymentMockRecorder) getStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getStatus", reflect.TypeOf((*MockDeployment)(nil).getStatus))
}

// hasRevision mocks base method
func (m *MockDeployment) hasRevision() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasRevision")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasRevision indicates an expected call of hasRevision
func (mr *MockDeploymentMockRecorder) hasRevision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasRevision", reflect.TypeOf((*MockDeployment)(nil).hasRevision))
}

// isCanaryPending mocks base method
func (m *MockDeployment) isCanaryPending() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isCanaryPending")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isCanaryPending indicates an expected call of isCanaryPending
func (mr *MockDeploymentMockRecorder) isCanaryPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isCanaryPending", reflect.TypeOf((*MockDeployment)(nil).isCanaryPending))
}

// isDeployed mocks base method
func (m *MockDeployment) isDeployed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isDeployed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isDeployed indicates an expected call of isDeployed
func (mr *MockDeploymentMockRecorder) isDeployed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDeployed", reflect.TypeOf((*MockDeployment)(nil).isDeployed))
}

// isReleaseEligible mocks base method
func (m *MockDeployment) isReleaseEligible() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isReleaseEligible")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isReleaseEligible indicates an expected call of isReleaseEligible
func (mr *MockDeploymentMockRecorder) isReleaseEligible() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isReleaseEligible", reflect.TypeOf((*MockDeployment)(nil).isReleaseEligible))
}

// markedAsFailed mocks base method
func (m *MockDeployment) markedAsFailed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "markedAsFailed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// markedAsFailed indicates an expected call of markedAsFailed
func (mr *MockDeploymentMockRecorder) markedAsFailed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "markedAsFailed", reflect.TypeOf((*MockDeployment)(nil).markedAsFailed))
}

// peakPercent mocks base method
func (m *MockDeployment) peakPercent() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "peakPercent")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// peakPercent indicates an expected call of peakPercent
func (mr *MockDeploymentMockRecorder) peakPercent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "peakPercent", reflect.TypeOf((*MockDeployment)(nil).peakPercent))
}

// retire mocks base method
func (m *MockDeployment) retire(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "retire", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// retire indicates an expected call of retire
func (mr *MockDeploymentMockRecorder) retire(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "retire", reflect.TypeOf((*MockDeployment)(nil).retire), arg0)
}

// schedulePermitsRelease mocks base method
func (m *MockDeployment) schedulePermitsRelease() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "schedulePermitsRelease")
	ret0, _ := ret[0].(bool)
	return ret0
}

// schedulePermitsRelease indicates an expected call of schedulePermitsRelease
func (mr *MockDeploymentMockRecorder) schedulePermitsRelease() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "schedulePermitsRelease", reflect.TypeOf((*MockDeployment)(nil).schedulePermitsRelease))
}

// setState mocks base method
func (m *MockDeployment) setState(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setState", arg0)
}

// setState indicates an expected call of setState
func (mr *MockDeploymentMockRecorder) setState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setState", reflect.TypeOf((*MockDeployment)(nil).setState), arg0)
}

// sync mocks base method
func (m *MockDeployment) sync(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// sync indicates an expected call of sync
func (mr *MockDeploymentMockRecorder) sync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sync", reflect.TypeOf((*MockDeployment)(nil).sync), arg0)
}

// syncCanaryRules mocks base method
func (m *MockDeployment) syncCanaryRules(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "syncCanaryRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// syncCanaryRules indicates an expected call of syncCanaryRules
func (mr *MockDeploymentMockRecorder) syncCanaryRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "syncCanaryRules", reflect.TypeOf((*MockDeployment)(nil).syncCanaryRules), arg0)
}

// syncSLIRules mocks base method
func (m *MockDeployment) syncSLIRules(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "syncSLIRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// syncSLIRules indicates an expected call of syncSLIRules
func (mr *MockDeploymentMockRecorder) syncSLIRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "syncSLIRules", reflect.TypeOf((*MockDeployment)(nil).syncSLIRules), arg0)
}
