// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_relay is a generated GoMock package.
package mock_relay

import (
	context "context"
	reflect "reflect"

	relay "github.com/blocknative/dreamboat/pkg"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
)

// MockRelayService is a mock of RelayService interface.
type MockRelayService struct {
	ctrl     *gomock.Controller
	recorder *MockRelayServiceMockRecorder
}

// MockRelayServiceMockRecorder is the mock recorder for MockRelayService.
type MockRelayServiceMockRecorder struct {
	mock *MockRelayService
}

// NewMockRelayService creates a new mock instance.
func NewMockRelayService(ctrl *gomock.Controller) *MockRelayService {
	mock := &MockRelayService{ctrl: ctrl}
	mock.recorder = &MockRelayServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelayService) EXPECT() *MockRelayServiceMockRecorder {
	return m.recorder
}

// GetBlockReceived mocks base method.
func (m *MockRelayService) GetBlockReceived(arg0 context.Context, arg1 relay.Slot) ([]relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceived", arg0, arg1)
	ret0, _ := ret[0].([]relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockReceived indicates an expected call of GetBlockReceived.
func (mr *MockRelayServiceMockRecorder) GetBlockReceived(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceived", reflect.TypeOf((*MockRelayService)(nil).GetBlockReceived), arg0, arg1)
}

// GetBlockReceivedByHash mocks base method.
func (m *MockRelayService) GetBlockReceivedByHash(arg0 context.Context, arg1 types.Hash) ([]relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceivedByHash", arg0, arg1)
	ret0, _ := ret[0].([]relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockReceivedByHash indicates an expected call of GetBlockReceivedByHash.
func (mr *MockRelayServiceMockRecorder) GetBlockReceivedByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceivedByHash", reflect.TypeOf((*MockRelayService)(nil).GetBlockReceivedByHash), arg0, arg1)
}

// GetBlockReceivedByNum mocks base method.
func (m *MockRelayService) GetBlockReceivedByNum(arg0 context.Context, arg1 uint64) ([]relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceivedByNum", arg0, arg1)
	ret0, _ := ret[0].([]relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockReceivedByNum indicates an expected call of GetBlockReceivedByNum.
func (mr *MockRelayServiceMockRecorder) GetBlockReceivedByNum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceivedByNum", reflect.TypeOf((*MockRelayService)(nil).GetBlockReceivedByNum), arg0, arg1)
}

// GetDelivered mocks base method.
func (m *MockRelayService) GetDelivered(arg0 context.Context, arg1 relay.Slot) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelivered", arg0, arg1)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelivered indicates an expected call of GetDelivered.
func (mr *MockRelayServiceMockRecorder) GetDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelivered", reflect.TypeOf((*MockRelayService)(nil).GetDelivered), arg0, arg1)
}

// GetDeliveredByHash mocks base method.
func (m *MockRelayService) GetDeliveredByHash(arg0 context.Context, arg1 types.Hash) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredByHash", arg0, arg1)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveredByHash indicates an expected call of GetDeliveredByHash.
func (mr *MockRelayServiceMockRecorder) GetDeliveredByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredByHash", reflect.TypeOf((*MockRelayService)(nil).GetDeliveredByHash), arg0, arg1)
}

// GetDeliveredByNum mocks base method.
func (m *MockRelayService) GetDeliveredByNum(arg0 context.Context, arg1 uint64) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredByNum", arg0, arg1)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveredByNum indicates an expected call of GetDeliveredByNum.
func (mr *MockRelayServiceMockRecorder) GetDeliveredByNum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredByNum", reflect.TypeOf((*MockRelayService)(nil).GetDeliveredByNum), arg0, arg1)
}

// GetDeliveredByPubKey mocks base method.
func (m *MockRelayService) GetDeliveredByPubKey(arg0 context.Context, arg1 types.PublicKey) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredByPubKey", arg0, arg1)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveredByPubKey indicates an expected call of GetDeliveredByPubKey.
func (mr *MockRelayServiceMockRecorder) GetDeliveredByPubKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredByPubKey", reflect.TypeOf((*MockRelayService)(nil).GetDeliveredByPubKey), arg0, arg1)
}

// GetHeader mocks base method.
func (m *MockRelayService) GetHeader(arg0 context.Context, arg1 relay.HeaderRequest) (*types.GetHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(*types.GetHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockRelayServiceMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockRelayService)(nil).GetHeader), arg0, arg1)
}

// GetPayload mocks base method.
func (m *MockRelayService) GetPayload(arg0 context.Context, arg1 *types.SignedBlindedBeaconBlock) (*types.GetPayloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1)
	ret0, _ := ret[0].(*types.GetPayloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayload indicates an expected call of GetPayload.
func (mr *MockRelayServiceMockRecorder) GetPayload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockRelayService)(nil).GetPayload), arg0, arg1)
}

// GetTailBlockReceived mocks base method.
func (m *MockRelayService) GetTailBlockReceived(arg0 context.Context, arg1 uint64) ([]relay.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTailBlockReceived", arg0, arg1)
	ret0, _ := ret[0].([]relay.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTailBlockReceived indicates an expected call of GetTailBlockReceived.
func (mr *MockRelayServiceMockRecorder) GetTailBlockReceived(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTailBlockReceived", reflect.TypeOf((*MockRelayService)(nil).GetTailBlockReceived), arg0, arg1)
}

// GetTailDelivered mocks base method.
func (m *MockRelayService) GetTailDelivered(arg0 context.Context, arg1 uint64) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTailDelivered", arg0, arg1)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTailDelivered indicates an expected call of GetTailDelivered.
func (mr *MockRelayServiceMockRecorder) GetTailDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTailDelivered", reflect.TypeOf((*MockRelayService)(nil).GetTailDelivered), arg0, arg1)
}

// GetTailDeliveredCursor mocks base method.
func (m *MockRelayService) GetTailDeliveredCursor(arg0 context.Context, arg1, arg2 uint64) ([]types.BidTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTailDeliveredCursor", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.BidTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTailDeliveredCursor indicates an expected call of GetTailDeliveredCursor.
func (mr *MockRelayServiceMockRecorder) GetTailDeliveredCursor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTailDeliveredCursor", reflect.TypeOf((*MockRelayService)(nil).GetTailDeliveredCursor), arg0, arg1, arg2)
}

// GetValidators mocks base method.
func (m *MockRelayService) GetValidators() relay.BuilderGetValidatorsResponseEntrySlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators")
	ret0, _ := ret[0].(relay.BuilderGetValidatorsResponseEntrySlice)
	return ret0
}

// GetValidators indicates an expected call of GetValidators.
func (mr *MockRelayServiceMockRecorder) GetValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockRelayService)(nil).GetValidators))
}

// RegisterValidator mocks base method.
func (m *MockRelayService) RegisterValidator(arg0 context.Context, arg1 []types.SignedValidatorRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidator", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterValidator indicates an expected call of RegisterValidator.
func (mr *MockRelayServiceMockRecorder) RegisterValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidator", reflect.TypeOf((*MockRelayService)(nil).RegisterValidator), arg0, arg1)
}

// Registration mocks base method.
func (m *MockRelayService) Registration(arg0 context.Context, arg1 types.PublicKey) (types.SignedValidatorRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", arg0, arg1)
	ret0, _ := ret[0].(types.SignedValidatorRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Registration indicates an expected call of Registration.
func (mr *MockRelayServiceMockRecorder) Registration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockRelayService)(nil).Registration), arg0, arg1)
}

// SubmitBlock mocks base method.
func (m *MockRelayService) SubmitBlock(arg0 context.Context, arg1 *types.BuilderSubmitBlockRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitBlock indicates an expected call of SubmitBlock.
func (mr *MockRelayServiceMockRecorder) SubmitBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBlock", reflect.TypeOf((*MockRelayService)(nil).SubmitBlock), arg0, arg1)
}
