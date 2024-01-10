// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/incredible-squaring-avs/core/chainio (interfaces: AvsWriterer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_writer.go -package=mocks github.com/Layr-Labs/incredible-squaring-avs/core/chainio AvsWriterer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractBLSPubkeyRegistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	contractBLSRegistryCoordinatorWithIndices "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	contractIncredibleSquaringTaskManager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsWriterer is a mock of AvsWriterer interface.
type MockAvsWriterer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsWritererMockRecorder
}

// MockAvsWritererMockRecorder is the mock recorder for MockAvsWriterer.
type MockAvsWritererMockRecorder struct {
	mock *MockAvsWriterer
}

// NewMockAvsWriterer creates a new mock instance.
func NewMockAvsWriterer(ctrl *gomock.Controller) *MockAvsWriterer {
	mock := &MockAvsWriterer{ctrl: ctrl}
	mock.recorder = &MockAvsWritererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsWriterer) EXPECT() *MockAvsWritererMockRecorder {
	return m.recorder
}

// DeregisterOperator mocks base method.
func (m *MockAvsWriterer) DeregisterOperator(arg0 context.Context, arg1 common.Address, arg2 []byte, arg3 contractBLSPubkeyRegistry.BN254G1Point) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterOperator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterOperator indicates an expected call of DeregisterOperator.
func (mr *MockAvsWritererMockRecorder) DeregisterOperator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterOperator", reflect.TypeOf((*MockAvsWriterer)(nil).DeregisterOperator), arg0, arg1, arg2, arg3)
}

// RaiseChallenge mocks base method.
func (m *MockAvsWriterer) RaiseChallenge(arg0 context.Context, arg1 contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTask, arg2 contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTaskResponse, arg3 contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTaskResponseMetadata, arg4 []contractIncredibleSquaringTaskManager.BN254G1Point) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RaiseChallenge", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RaiseChallenge indicates an expected call of RaiseChallenge.
func (mr *MockAvsWritererMockRecorder) RaiseChallenge(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RaiseChallenge", reflect.TypeOf((*MockAvsWriterer)(nil).RaiseChallenge), arg0, arg1, arg2, arg3, arg4)
}

// RegisterOperatorWithAVSRegistryCoordinator mocks base method.
func (m *MockAvsWriterer) RegisterOperatorWithAVSRegistryCoordinator(arg0 context.Context, arg1 []byte, arg2 contractBLSRegistryCoordinatorWithIndices.BN254G1Point, arg3 string) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOperatorWithAVSRegistryCoordinator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOperatorWithAVSRegistryCoordinator indicates an expected call of RegisterOperatorWithAVSRegistryCoordinator.
func (mr *MockAvsWritererMockRecorder) RegisterOperatorWithAVSRegistryCoordinator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOperatorWithAVSRegistryCoordinator", reflect.TypeOf((*MockAvsWriterer)(nil).RegisterOperatorWithAVSRegistryCoordinator), arg0, arg1, arg2, arg3)
}

// SendAggregatedResponse mocks base method.
func (m *MockAvsWriterer) SendAggregatedResponse(arg0 context.Context, arg1 contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTask, arg2 contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTaskResponse, arg3 contractIncredibleSquaringTaskManager.IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAggregatedResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendAggregatedResponse indicates an expected call of SendAggregatedResponse.
func (mr *MockAvsWritererMockRecorder) SendAggregatedResponse(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAggregatedResponse", reflect.TypeOf((*MockAvsWriterer)(nil).SendAggregatedResponse), arg0, arg1, arg2, arg3)
}

// SendNewEchoTask mocks base method.
func (m *MockAvsWriterer) SendNewEchoTask(arg0 context.Context, arg1 []byte, arg2 uint32, arg3 []byte) (contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTask, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNewEchoTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTask)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendNewEchoTask indicates an expected call of SendNewEchoTask.
func (mr *MockAvsWritererMockRecorder) SendNewEchoTask(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNewEchoTask", reflect.TypeOf((*MockAvsWriterer)(nil).SendNewEchoTask), arg0, arg1, arg2, arg3)
}

// UpdateStakes mocks base method.
func (m *MockAvsWriterer) UpdateStakes(arg0 context.Context, arg1 []common.Address) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStakes", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStakes indicates an expected call of UpdateStakes.
func (mr *MockAvsWritererMockRecorder) UpdateStakes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStakes", reflect.TypeOf((*MockAvsWriterer)(nil).UpdateStakes), arg0, arg1)
}
