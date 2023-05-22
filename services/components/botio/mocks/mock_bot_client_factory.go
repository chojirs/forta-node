// Code generated by MockGen. DO NOT EDIT.
// Source: services/components/botio/bot_client_factory.go

// Package mock_botio is a generated GoMock package.
package mock_botio

import (
	context "context"
	reflect "reflect"

	config "github.com/forta-network/forta-node/config"
	botio "github.com/forta-network/forta-node/services/components/botio"
	gomock "github.com/golang/mock/gomock"
)

// MockBotClientFactory is a mock of BotClientFactory interface.
type MockBotClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockBotClientFactoryMockRecorder
}

// MockBotClientFactoryMockRecorder is the mock recorder for MockBotClientFactory.
type MockBotClientFactoryMockRecorder struct {
	mock *MockBotClientFactory
}

// NewMockBotClientFactory creates a new mock instance.
func NewMockBotClientFactory(ctrl *gomock.Controller) *MockBotClientFactory {
	mock := &MockBotClientFactory{ctrl: ctrl}
	mock.recorder = &MockBotClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBotClientFactory) EXPECT() *MockBotClientFactoryMockRecorder {
	return m.recorder
}

// NewBotClient mocks base method.
func (m *MockBotClientFactory) NewBotClient(ctx context.Context, botConfig config.AgentConfig) botio.BotClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBotClient", ctx, botConfig)
	ret0, _ := ret[0].(botio.BotClient)
	return ret0
}

// NewBotClient indicates an expected call of NewBotClient.
func (mr *MockBotClientFactoryMockRecorder) NewBotClient(ctx, botConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBotClient", reflect.TypeOf((*MockBotClientFactory)(nil).NewBotClient), ctx, botConfig)
}
