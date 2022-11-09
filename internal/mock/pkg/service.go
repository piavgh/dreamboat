// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_relay is a generated GoMock package.
package mock_relay

import (
	context "context"
	structs "github.com/blocknative/dreamboat/pkg/structs"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRelayService is a mock of RelayService interface
type MockRelayService struct {
	ctrl     *gomock.Controller
	recorder *MockRelayServiceMockRecorder
}

// MockRelayServiceMockRecorder is the mock recorder for MockRelayService
type MockRelayServiceMockRecorder struct {
	mock *MockRelayService
}

// NewMockRelayService creates a new mock instance
func NewMockRelayService(ctrl *gomock.Controller) *MockRelayService {
	mock := &MockRelayService{ctrl: ctrl}
	mock.recorder = &MockRelayServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelayService) EXPECT() *MockRelayServiceMockRecorder {
	return m.recorder
}

// RegisterValidator mocks base method
func (m *MockRelayService) RegisterValidator(arg0 context.Context, arg1 []structs.SignedValidatorRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidator", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterValidator indicates an expected call of RegisterValidator
func (mr *MockRelayServiceMockRecorder) RegisterValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidator", reflect.TypeOf((*MockRelayService)(nil).RegisterValidator), arg0, arg1)
}

// GetHeader mocks base method
func (m *MockRelayService) GetHeader(arg0 context.Context, arg1 structs.HeaderRequest) (*types.GetHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(*types.GetHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockRelayServiceMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockRelayService)(nil).GetHeader), arg0, arg1)
}

// GetPayload mocks base method
func (m *MockRelayService) GetPayload(arg0 context.Context, arg1 *types.SignedBlindedBeaconBlock) (*types.GetPayloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1)
	ret0, _ := ret[0].(*types.GetPayloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayload indicates an expected call of GetPayload
func (mr *MockRelayServiceMockRecorder) GetPayload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockRelayService)(nil).GetPayload), arg0, arg1)
}

// SubmitBlock mocks base method
func (m *MockRelayService) SubmitBlock(arg0 context.Context, arg1 *types.BuilderSubmitBlockRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitBlock indicates an expected call of SubmitBlock
func (mr *MockRelayServiceMockRecorder) SubmitBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBlock", reflect.TypeOf((*MockRelayService)(nil).SubmitBlock), arg0, arg1)
}

// GetValidators mocks base method
func (m *MockRelayService) GetValidators() structs.BuilderGetValidatorsResponseEntrySlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators")
	ret0, _ := ret[0].(structs.BuilderGetValidatorsResponseEntrySlice)
	return ret0
}

// GetValidators indicates an expected call of GetValidators
func (mr *MockRelayServiceMockRecorder) GetValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockRelayService)(nil).GetValidators))
}

// GetPayloadDelivered mocks base method
func (m *MockRelayService) GetPayloadDelivered(arg0 context.Context, arg1 structs.TraceQuery) ([]structs.BidTraceExtended, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayloadDelivered", arg0, arg1)
	ret0, _ := ret[0].([]structs.BidTraceExtended)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayloadDelivered indicates an expected call of GetPayloadDelivered
func (mr *MockRelayServiceMockRecorder) GetPayloadDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayloadDelivered", reflect.TypeOf((*MockRelayService)(nil).GetPayloadDelivered), arg0, arg1)
}

// GetBlockReceived mocks base method
func (m *MockRelayService) GetBlockReceived(arg0 context.Context, arg1 structs.TraceQuery) ([]structs.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceived", arg0, arg1)
	ret0, _ := ret[0].([]structs.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockReceived indicates an expected call of GetBlockReceived
func (mr *MockRelayServiceMockRecorder) GetBlockReceived(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceived", reflect.TypeOf((*MockRelayService)(nil).GetBlockReceived), arg0, arg1)
}

// Registration mocks base method
func (m *MockRelayService) Registration(arg0 context.Context, arg1 types.PublicKey) (types.SignedValidatorRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", arg0, arg1)
	ret0, _ := ret[0].(types.SignedValidatorRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Registration indicates an expected call of Registration
func (mr *MockRelayServiceMockRecorder) Registration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockRelayService)(nil).Registration), arg0, arg1)
}
