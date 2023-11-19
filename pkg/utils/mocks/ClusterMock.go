// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/backplane-cli/pkg/utils (interfaces: ClusterUtils)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	utils "github.com/openshift/backplane-cli/pkg/utils"
)

// MockClusterUtils is a mock of ClusterUtils interface.
type MockClusterUtils struct {
	ctrl     *gomock.Controller
	recorder *MockClusterUtilsMockRecorder
}

// MockClusterUtilsMockRecorder is the mock recorder for MockClusterUtils.
type MockClusterUtilsMockRecorder struct {
	mock *MockClusterUtils
}

// NewMockClusterUtils creates a new mock instance.
func NewMockClusterUtils(ctrl *gomock.Controller) *MockClusterUtils {
	mock := &MockClusterUtils{ctrl: ctrl}
	mock.recorder = &MockClusterUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterUtils) EXPECT() *MockClusterUtilsMockRecorder {
	return m.recorder
}

// GetBackplaneCluster mocks base method.
func (m *MockClusterUtils) GetBackplaneCluster(arg0 ...string) (utils.BackplaneCluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackplaneCluster", varargs...)
	ret0, _ := ret[0].(utils.BackplaneCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneCluster indicates an expected call of GetBackplaneCluster.
func (mr *MockClusterUtilsMockRecorder) GetBackplaneCluster(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneCluster", reflect.TypeOf((*MockClusterUtils)(nil).GetBackplaneCluster), arg0...)
}

// GetBackplaneClusterFromClusterKey mocks base method.
func (m *MockClusterUtils) GetBackplaneClusterFromClusterKey(arg0 string) (utils.BackplaneCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackplaneClusterFromClusterKey", arg0)
	ret0, _ := ret[0].(utils.BackplaneCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneClusterFromClusterKey indicates an expected call of GetBackplaneClusterFromClusterKey.
func (mr *MockClusterUtilsMockRecorder) GetBackplaneClusterFromClusterKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneClusterFromClusterKey", reflect.TypeOf((*MockClusterUtils)(nil).GetBackplaneClusterFromClusterKey), arg0)
}

// GetBackplaneClusterFromConfig mocks base method.
func (m *MockClusterUtils) GetBackplaneClusterFromConfig() (utils.BackplaneCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackplaneClusterFromConfig")
	ret0, _ := ret[0].(utils.BackplaneCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneClusterFromConfig indicates an expected call of GetBackplaneClusterFromConfig.
func (mr *MockClusterUtilsMockRecorder) GetBackplaneClusterFromConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneClusterFromConfig", reflect.TypeOf((*MockClusterUtils)(nil).GetBackplaneClusterFromConfig))
}

// GetClusterIDAndHostFromClusterURL mocks base method.
func (m *MockClusterUtils) GetClusterIDAndHostFromClusterURL(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterIDAndHostFromClusterURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetClusterIDAndHostFromClusterURL indicates an expected call of GetClusterIDAndHostFromClusterURL.
func (mr *MockClusterUtilsMockRecorder) GetClusterIDAndHostFromClusterURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterIDAndHostFromClusterURL", reflect.TypeOf((*MockClusterUtils)(nil).GetClusterIDAndHostFromClusterURL), arg0)
}
