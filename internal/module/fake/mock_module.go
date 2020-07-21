// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/module (interfaces: Module)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"

	module "github.com/vmware-tanzu/octant/internal/module"
	octant "github.com/vmware-tanzu/octant/internal/octant"
	navigation "github.com/vmware-tanzu/octant/pkg/navigation"
	component "github.com/vmware-tanzu/octant/pkg/view/component"
)

// MockModule is a mock of Module interface
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// AddCRD mocks base method
func (m *MockModule) AddCRD(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCRD", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCRD indicates an expected call of AddCRD
func (mr *MockModuleMockRecorder) AddCRD(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCRD", reflect.TypeOf((*MockModule)(nil).AddCRD), arg0, arg1)
}

// ClientRequestHandlers mocks base method
func (m *MockModule) ClientRequestHandlers() []octant.ClientRequestHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRequestHandlers")
	ret0, _ := ret[0].([]octant.ClientRequestHandler)
	return ret0
}

// ClientRequestHandlers indicates an expected call of ClientRequestHandlers
func (mr *MockModuleMockRecorder) ClientRequestHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRequestHandlers", reflect.TypeOf((*MockModule)(nil).ClientRequestHandlers))
}

// Content mocks base method
func (m *MockModule) Content(arg0 context.Context, arg1 string, arg2 module.ContentOptions) (component.ContentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Content", arg0, arg1, arg2)
	ret0, _ := ret[0].(component.ContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Content indicates an expected call of Content
func (mr *MockModuleMockRecorder) Content(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockModule)(nil).Content), arg0, arg1, arg2)
}

// ContentPath mocks base method
func (m *MockModule) ContentPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentPath indicates an expected call of ContentPath
func (mr *MockModuleMockRecorder) ContentPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentPath", reflect.TypeOf((*MockModule)(nil).ContentPath))
}

// Generators mocks base method
func (m *MockModule) Generators() []octant.Generator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generators")
	ret0, _ := ret[0].([]octant.Generator)
	return ret0
}

// Generators indicates an expected call of Generators
func (mr *MockModuleMockRecorder) Generators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generators", reflect.TypeOf((*MockModule)(nil).Generators))
}

// GroupVersionKindPath mocks base method
func (m *MockModule) GroupVersionKindPath(arg0, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionKindPath", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupVersionKindPath indicates an expected call of GroupVersionKindPath
func (mr *MockModuleMockRecorder) GroupVersionKindPath(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionKindPath", reflect.TypeOf((*MockModule)(nil).GroupVersionKindPath), arg0, arg1, arg2, arg3)
}

// Name mocks base method
func (m *MockModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModule)(nil).Name))
}

// Navigation mocks base method
func (m *MockModule) Navigation(arg0 context.Context, arg1, arg2 string) ([]navigation.Navigation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Navigation", arg0, arg1, arg2)
	ret0, _ := ret[0].([]navigation.Navigation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Navigation indicates an expected call of Navigation
func (mr *MockModuleMockRecorder) Navigation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Navigation", reflect.TypeOf((*MockModule)(nil).Navigation), arg0, arg1, arg2)
}

// RemoveCRD mocks base method
func (m *MockModule) RemoveCRD(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCRD", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCRD indicates an expected call of RemoveCRD
func (mr *MockModuleMockRecorder) RemoveCRD(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCRD", reflect.TypeOf((*MockModule)(nil).RemoveCRD), arg0, arg1)
}

// ResetCRDs mocks base method
func (m *MockModule) ResetCRDs(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetCRDs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetCRDs indicates an expected call of ResetCRDs
func (mr *MockModuleMockRecorder) ResetCRDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCRDs", reflect.TypeOf((*MockModule)(nil).ResetCRDs), arg0)
}

// SetContext mocks base method
func (m *MockModule) SetContext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContext indicates an expected call of SetContext
func (mr *MockModuleMockRecorder) SetContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContext", reflect.TypeOf((*MockModule)(nil).SetContext), arg0, arg1)
}

// SetNamespace mocks base method
func (m *MockModule) SetNamespace(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNamespace indicates an expected call of SetNamespace
func (mr *MockModuleMockRecorder) SetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockModule)(nil).SetNamespace), arg0)
}

// Start mocks base method
func (m *MockModule) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockModuleMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockModule)(nil).Start))
}

// Stop mocks base method
func (m *MockModule) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockModuleMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockModule)(nil).Stop))
}

// SupportedGroupVersionKind mocks base method
func (m *MockModule) SupportedGroupVersionKind() []schema.GroupVersionKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportedGroupVersionKind")
	ret0, _ := ret[0].([]schema.GroupVersionKind)
	return ret0
}

// SupportedGroupVersionKind indicates an expected call of SupportedGroupVersionKind
func (mr *MockModuleMockRecorder) SupportedGroupVersionKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedGroupVersionKind", reflect.TypeOf((*MockModule)(nil).SupportedGroupVersionKind))
}
