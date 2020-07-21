// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/module (interfaces: ManagerInterface)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	module "github.com/vmware-tanzu/octant/internal/module"
	octant "github.com/vmware-tanzu/octant/internal/octant"
)

// MockManagerInterface is a mock of ManagerInterface interface
type MockManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockManagerInterfaceMockRecorder
}

// MockManagerInterfaceMockRecorder is the mock recorder for MockManagerInterface
type MockManagerInterfaceMockRecorder struct {
	mock *MockManagerInterface
}

// NewMockManagerInterface creates a new mock instance
func NewMockManagerInterface(ctrl *gomock.Controller) *MockManagerInterface {
	mock := &MockManagerInterface{ctrl: ctrl}
	mock.recorder = &MockManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerInterface) EXPECT() *MockManagerInterfaceMockRecorder {
	return m.recorder
}

// ClientRequestHandlers mocks base method
func (m *MockManagerInterface) ClientRequestHandlers() []octant.ClientRequestHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRequestHandlers")
	ret0, _ := ret[0].([]octant.ClientRequestHandler)
	return ret0
}

// ClientRequestHandlers indicates an expected call of ClientRequestHandlers
func (mr *MockManagerInterfaceMockRecorder) ClientRequestHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRequestHandlers", reflect.TypeOf((*MockManagerInterface)(nil).ClientRequestHandlers))
}

// GetNamespace mocks base method
func (m *MockManagerInterface) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockManagerInterfaceMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockManagerInterface)(nil).GetNamespace))
}

// ModuleForContentPath mocks base method
func (m *MockManagerInterface) ModuleForContentPath(arg0 string) (module.Module, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModuleForContentPath", arg0)
	ret0, _ := ret[0].(module.Module)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModuleForContentPath indicates an expected call of ModuleForContentPath
func (mr *MockManagerInterfaceMockRecorder) ModuleForContentPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModuleForContentPath", reflect.TypeOf((*MockManagerInterface)(nil).ModuleForContentPath), arg0)
}

// Modules mocks base method
func (m *MockManagerInterface) Modules() []module.Module {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modules")
	ret0, _ := ret[0].([]module.Module)
	return ret0
}

// Modules indicates an expected call of Modules
func (mr *MockManagerInterfaceMockRecorder) Modules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modules", reflect.TypeOf((*MockManagerInterface)(nil).Modules))
}

// ObjectPath mocks base method
func (m *MockManagerInterface) ObjectPath(arg0, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectPath", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPath indicates an expected call of ObjectPath
func (mr *MockManagerInterfaceMockRecorder) ObjectPath(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPath", reflect.TypeOf((*MockManagerInterface)(nil).ObjectPath), arg0, arg1, arg2, arg3)
}

// Register mocks base method
func (m *MockManagerInterface) Register(arg0 module.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockManagerInterfaceMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockManagerInterface)(nil).Register), arg0)
}

// SetNamespace mocks base method
func (m *MockManagerInterface) SetNamespace(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", arg0)
}

// SetNamespace indicates an expected call of SetNamespace
func (mr *MockManagerInterfaceMockRecorder) SetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockManagerInterface)(nil).SetNamespace), arg0)
}

// Unregister mocks base method
func (m *MockManagerInterface) Unregister(arg0 module.Module) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unregister", arg0)
}

// Unregister indicates an expected call of Unregister
func (mr *MockManagerInterfaceMockRecorder) Unregister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockManagerInterface)(nil).Unregister), arg0)
}

// UpdateContext mocks base method
func (m *MockManagerInterface) UpdateContext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContext indicates an expected call of UpdateContext
func (mr *MockManagerInterfaceMockRecorder) UpdateContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContext", reflect.TypeOf((*MockManagerInterface)(nil).UpdateContext), arg0, arg1)
}
