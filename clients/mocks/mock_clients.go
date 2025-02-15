// Code generated by MockGen. DO NOT EDIT.
// Source: clients/interfaces.go

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	context "context"
	reflect "reflect"
	time "time"

	types "github.com/docker/docker/api/types"
	domain "github.com/forta-network/forta-core-go/domain"
	docker "github.com/forta-network/forta-node/clients/docker"
	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
)

// MockDockerClient is a mock of DockerClient interface.
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerClientMockRecorder
}

// MockDockerClientMockRecorder is the mock recorder for MockDockerClient.
type MockDockerClientMockRecorder struct {
	mock *MockDockerClient
}

// NewMockDockerClient creates a new mock instance.
func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &MockDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerClient) EXPECT() *MockDockerClientMockRecorder {
	return m.recorder
}

// AttachNetwork mocks base method.
func (m *MockDockerClient) AttachNetwork(ctx context.Context, containerID, networkID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachNetwork", ctx, containerID, networkID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachNetwork indicates an expected call of AttachNetwork.
func (mr *MockDockerClientMockRecorder) AttachNetwork(ctx, containerID, networkID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachNetwork", reflect.TypeOf((*MockDockerClient)(nil).AttachNetwork), ctx, containerID, networkID)
}

// DetachNetwork mocks base method.
func (m *MockDockerClient) DetachNetwork(ctx context.Context, containerID, networkID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachNetwork", ctx, containerID, networkID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachNetwork indicates an expected call of DetachNetwork.
func (mr *MockDockerClientMockRecorder) DetachNetwork(ctx, containerID, networkID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachNetwork", reflect.TypeOf((*MockDockerClient)(nil).DetachNetwork), ctx, containerID, networkID)
}

// EnsureInternalNetwork mocks base method.
func (m *MockDockerClient) EnsureInternalNetwork(ctx context.Context, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureInternalNetwork", ctx, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureInternalNetwork indicates an expected call of EnsureInternalNetwork.
func (mr *MockDockerClientMockRecorder) EnsureInternalNetwork(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureInternalNetwork", reflect.TypeOf((*MockDockerClient)(nil).EnsureInternalNetwork), ctx, name)
}

// EnsureLocalImage mocks base method.
func (m *MockDockerClient) EnsureLocalImage(ctx context.Context, name, ref string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureLocalImage", ctx, name, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureLocalImage indicates an expected call of EnsureLocalImage.
func (mr *MockDockerClientMockRecorder) EnsureLocalImage(ctx, name, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureLocalImage", reflect.TypeOf((*MockDockerClient)(nil).EnsureLocalImage), ctx, name, ref)
}

// EnsureLocalImages mocks base method.
func (m *MockDockerClient) EnsureLocalImages(ctx context.Context, timeoutPerPull time.Duration, imagePulls []docker.ImagePull) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureLocalImages", ctx, timeoutPerPull, imagePulls)
	ret0, _ := ret[0].([]error)
	return ret0
}

// EnsureLocalImages indicates an expected call of EnsureLocalImages.
func (mr *MockDockerClientMockRecorder) EnsureLocalImages(ctx, timeoutPerPull, imagePulls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureLocalImages", reflect.TypeOf((*MockDockerClient)(nil).EnsureLocalImages), ctx, timeoutPerPull, imagePulls)
}

// EnsurePublicNetwork mocks base method.
func (m *MockDockerClient) EnsurePublicNetwork(ctx context.Context, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsurePublicNetwork", ctx, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePublicNetwork indicates an expected call of EnsurePublicNetwork.
func (mr *MockDockerClientMockRecorder) EnsurePublicNetwork(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePublicNetwork", reflect.TypeOf((*MockDockerClient)(nil).EnsurePublicNetwork), ctx, name)
}

// GetContainerByID mocks base method.
func (m *MockDockerClient) GetContainerByID(ctx context.Context, id string) (*types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerByID", ctx, id)
	ret0, _ := ret[0].(*types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerByID indicates an expected call of GetContainerByID.
func (mr *MockDockerClientMockRecorder) GetContainerByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerByID", reflect.TypeOf((*MockDockerClient)(nil).GetContainerByID), ctx, id)
}

// GetContainerByName mocks base method.
func (m *MockDockerClient) GetContainerByName(ctx context.Context, name string) (*types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerByName", ctx, name)
	ret0, _ := ret[0].(*types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerByName indicates an expected call of GetContainerByName.
func (mr *MockDockerClientMockRecorder) GetContainerByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerByName", reflect.TypeOf((*MockDockerClient)(nil).GetContainerByName), ctx, name)
}

// GetContainerFromRemoteAddr mocks base method.
func (m *MockDockerClient) GetContainerFromRemoteAddr(ctx context.Context, hostPort string) (*types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerFromRemoteAddr", ctx, hostPort)
	ret0, _ := ret[0].(*types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerFromRemoteAddr indicates an expected call of GetContainerFromRemoteAddr.
func (mr *MockDockerClientMockRecorder) GetContainerFromRemoteAddr(ctx, hostPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerFromRemoteAddr", reflect.TypeOf((*MockDockerClient)(nil).GetContainerFromRemoteAddr), ctx, hostPort)
}

// GetContainerLogs mocks base method.
func (m *MockDockerClient) GetContainerLogs(ctx context.Context, containerID, tail string, truncate int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerLogs", ctx, containerID, tail, truncate)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerLogs indicates an expected call of GetContainerLogs.
func (mr *MockDockerClientMockRecorder) GetContainerLogs(ctx, containerID, tail, truncate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerLogs", reflect.TypeOf((*MockDockerClient)(nil).GetContainerLogs), ctx, containerID, tail, truncate)
}

// GetContainers mocks base method.
func (m *MockDockerClient) GetContainers(ctx context.Context) (docker.ContainerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainers", ctx)
	ret0, _ := ret[0].(docker.ContainerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers.
func (mr *MockDockerClientMockRecorder) GetContainers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockDockerClient)(nil).GetContainers), ctx)
}

// GetContainersByLabel mocks base method.
func (m *MockDockerClient) GetContainersByLabel(ctx context.Context, name, value string) (docker.ContainerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainersByLabel", ctx, name, value)
	ret0, _ := ret[0].(docker.ContainerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainersByLabel indicates an expected call of GetContainersByLabel.
func (mr *MockDockerClientMockRecorder) GetContainersByLabel(ctx, name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainersByLabel", reflect.TypeOf((*MockDockerClient)(nil).GetContainersByLabel), ctx, name, value)
}

// GetFortaServiceContainers mocks base method.
func (m *MockDockerClient) GetFortaServiceContainers(ctx context.Context) (docker.ContainerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFortaServiceContainers", ctx)
	ret0, _ := ret[0].(docker.ContainerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFortaServiceContainers indicates an expected call of GetFortaServiceContainers.
func (mr *MockDockerClientMockRecorder) GetFortaServiceContainers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFortaServiceContainers", reflect.TypeOf((*MockDockerClient)(nil).GetFortaServiceContainers), ctx)
}

// HasLocalImage mocks base method.
func (m *MockDockerClient) HasLocalImage(ctx context.Context, ref string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasLocalImage", ctx, ref)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasLocalImage indicates an expected call of HasLocalImage.
func (mr *MockDockerClientMockRecorder) HasLocalImage(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLocalImage", reflect.TypeOf((*MockDockerClient)(nil).HasLocalImage), ctx, ref)
}

// InspectContainer mocks base method.
func (m *MockDockerClient) InspectContainer(ctx context.Context, id string) (*types.ContainerJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectContainer", ctx, id)
	ret0, _ := ret[0].(*types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer.
func (mr *MockDockerClientMockRecorder) InspectContainer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainer", reflect.TypeOf((*MockDockerClient)(nil).InspectContainer), ctx, id)
}

// InterruptContainer mocks base method.
func (m *MockDockerClient) InterruptContainer(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterruptContainer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// InterruptContainer indicates an expected call of InterruptContainer.
func (mr *MockDockerClientMockRecorder) InterruptContainer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterruptContainer", reflect.TypeOf((*MockDockerClient)(nil).InterruptContainer), ctx, id)
}

// Nuke mocks base method.
func (m *MockDockerClient) Nuke(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nuke", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Nuke indicates an expected call of Nuke.
func (mr *MockDockerClientMockRecorder) Nuke(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nuke", reflect.TypeOf((*MockDockerClient)(nil).Nuke), ctx)
}

// Prune mocks base method.
func (m *MockDockerClient) Prune(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prune", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prune indicates an expected call of Prune.
func (mr *MockDockerClientMockRecorder) Prune(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockDockerClient)(nil).Prune), ctx)
}

// PullImage mocks base method.
func (m *MockDockerClient) PullImage(ctx context.Context, refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", ctx, refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullImage indicates an expected call of PullImage.
func (mr *MockDockerClientMockRecorder) PullImage(ctx, refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockDockerClient)(nil).PullImage), ctx, refStr)
}

// RemoveContainer mocks base method.
func (m *MockDockerClient) RemoveContainer(ctx context.Context, containerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveContainer", ctx, containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer.
func (mr *MockDockerClientMockRecorder) RemoveContainer(ctx, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockDockerClient)(nil).RemoveContainer), ctx, containerID)
}

// RemoveImage mocks base method.
func (m *MockDockerClient) RemoveImage(ctx context.Context, refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", ctx, refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockDockerClientMockRecorder) RemoveImage(ctx, refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockDockerClient)(nil).RemoveImage), ctx, refStr)
}

// RemoveNetworkByName mocks base method.
func (m *MockDockerClient) RemoveNetworkByName(ctx context.Context, networkName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNetworkByName", ctx, networkName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetworkByName indicates an expected call of RemoveNetworkByName.
func (mr *MockDockerClientMockRecorder) RemoveNetworkByName(ctx, networkName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNetworkByName", reflect.TypeOf((*MockDockerClient)(nil).RemoveNetworkByName), ctx, networkName)
}

// SetImagePullCooldown mocks base method.
func (m *MockDockerClient) SetImagePullCooldown(threshold int, cooldownDuration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetImagePullCooldown", threshold, cooldownDuration)
}

// SetImagePullCooldown indicates an expected call of SetImagePullCooldown.
func (mr *MockDockerClientMockRecorder) SetImagePullCooldown(threshold, cooldownDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetImagePullCooldown", reflect.TypeOf((*MockDockerClient)(nil).SetImagePullCooldown), threshold, cooldownDuration)
}

// StartContainer mocks base method.
func (m *MockDockerClient) StartContainer(ctx context.Context, config docker.ContainerConfig) (*docker.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", ctx, config)
	ret0, _ := ret[0].(*docker.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartContainer indicates an expected call of StartContainer.
func (mr *MockDockerClientMockRecorder) StartContainer(ctx, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockDockerClient)(nil).StartContainer), ctx, config)
}

// StartContainerWithID mocks base method.
func (m *MockDockerClient) StartContainerWithID(ctx context.Context, containerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainerWithID", ctx, containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainerWithID indicates an expected call of StartContainerWithID.
func (mr *MockDockerClientMockRecorder) StartContainerWithID(ctx, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainerWithID", reflect.TypeOf((*MockDockerClient)(nil).StartContainerWithID), ctx, containerID)
}

// StopContainer mocks base method.
func (m *MockDockerClient) StopContainer(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer.
func (mr *MockDockerClientMockRecorder) StopContainer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockDockerClient)(nil).StopContainer), ctx, id)
}

// TerminateContainer mocks base method.
func (m *MockDockerClient) TerminateContainer(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateContainer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateContainer indicates an expected call of TerminateContainer.
func (mr *MockDockerClientMockRecorder) TerminateContainer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateContainer", reflect.TypeOf((*MockDockerClient)(nil).TerminateContainer), ctx, id)
}

// WaitContainerExit mocks base method.
func (m *MockDockerClient) WaitContainerExit(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitContainerExit", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitContainerExit indicates an expected call of WaitContainerExit.
func (mr *MockDockerClientMockRecorder) WaitContainerExit(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitContainerExit", reflect.TypeOf((*MockDockerClient)(nil).WaitContainerExit), ctx, id)
}

// WaitContainerPrune mocks base method.
func (m *MockDockerClient) WaitContainerPrune(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitContainerPrune", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitContainerPrune indicates an expected call of WaitContainerPrune.
func (mr *MockDockerClientMockRecorder) WaitContainerPrune(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitContainerPrune", reflect.TypeOf((*MockDockerClient)(nil).WaitContainerPrune), ctx, id)
}

// WaitContainerStart mocks base method.
func (m *MockDockerClient) WaitContainerStart(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitContainerStart", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitContainerStart indicates an expected call of WaitContainerStart.
func (mr *MockDockerClientMockRecorder) WaitContainerStart(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitContainerStart", reflect.TypeOf((*MockDockerClient)(nil).WaitContainerStart), ctx, id)
}

// MockMessageClient is a mock of MessageClient interface.
type MockMessageClient struct {
	ctrl     *gomock.Controller
	recorder *MockMessageClientMockRecorder
}

// MockMessageClientMockRecorder is the mock recorder for MockMessageClient.
type MockMessageClientMockRecorder struct {
	mock *MockMessageClient
}

// NewMockMessageClient creates a new mock instance.
func NewMockMessageClient(ctrl *gomock.Controller) *MockMessageClient {
	mock := &MockMessageClient{ctrl: ctrl}
	mock.recorder = &MockMessageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageClient) EXPECT() *MockMessageClientMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockMessageClient) Publish(subject string, payload interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", subject, payload)
}

// Publish indicates an expected call of Publish.
func (mr *MockMessageClientMockRecorder) Publish(subject, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockMessageClient)(nil).Publish), subject, payload)
}

// PublishProto mocks base method.
func (m *MockMessageClient) PublishProto(subject string, payload proto.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishProto", subject, payload)
}

// PublishProto indicates an expected call of PublishProto.
func (mr *MockMessageClientMockRecorder) PublishProto(subject, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishProto", reflect.TypeOf((*MockMessageClient)(nil).PublishProto), subject, payload)
}

// Subscribe mocks base method.
func (m *MockMessageClient) Subscribe(subject string, handler interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", subject, handler)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockMessageClientMockRecorder) Subscribe(subject, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockMessageClient)(nil).Subscribe), subject, handler)
}

// MockAlertAPIClient is a mock of AlertAPIClient interface.
type MockAlertAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockAlertAPIClientMockRecorder
}

// MockAlertAPIClientMockRecorder is the mock recorder for MockAlertAPIClient.
type MockAlertAPIClientMockRecorder struct {
	mock *MockAlertAPIClient
}

// NewMockAlertAPIClient creates a new mock instance.
func NewMockAlertAPIClient(ctrl *gomock.Controller) *MockAlertAPIClient {
	mock := &MockAlertAPIClient{ctrl: ctrl}
	mock.recorder = &MockAlertAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertAPIClient) EXPECT() *MockAlertAPIClientMockRecorder {
	return m.recorder
}

// PostBatch mocks base method.
func (m *MockAlertAPIClient) PostBatch(batch *domain.AlertBatchRequest, token string) (*domain.AlertBatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBatch", batch, token)
	ret0, _ := ret[0].(*domain.AlertBatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBatch indicates an expected call of PostBatch.
func (mr *MockAlertAPIClientMockRecorder) PostBatch(batch, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBatch", reflect.TypeOf((*MockAlertAPIClient)(nil).PostBatch), batch, token)
}

// MockIPAuthenticator is a mock of IPAuthenticator interface.
type MockIPAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockIPAuthenticatorMockRecorder
}

// MockIPAuthenticatorMockRecorder is the mock recorder for MockIPAuthenticator.
type MockIPAuthenticatorMockRecorder struct {
	mock *MockIPAuthenticator
}

// NewMockIPAuthenticator creates a new mock instance.
func NewMockIPAuthenticator(ctrl *gomock.Controller) *MockIPAuthenticator {
	mock := &MockIPAuthenticator{ctrl: ctrl}
	mock.recorder = &MockIPAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPAuthenticator) EXPECT() *MockIPAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockIPAuthenticator) Authenticate(ctx context.Context, hostPort string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, hostPort)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockIPAuthenticatorMockRecorder) Authenticate(ctx, hostPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockIPAuthenticator)(nil).Authenticate), ctx, hostPort)
}

// FindAgentByContainerName mocks base method.
func (m *MockIPAuthenticator) FindAgentByContainerName(containerName string) (*config.AgentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAgentByContainerName", containerName)
	ret0, _ := ret[0].(*config.AgentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAgentByContainerName indicates an expected call of FindAgentByContainerName.
func (mr *MockIPAuthenticatorMockRecorder) FindAgentByContainerName(containerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAgentByContainerName", reflect.TypeOf((*MockIPAuthenticator)(nil).FindAgentByContainerName), containerName)
}

// FindAgentFromRemoteAddr mocks base method.
func (m *MockIPAuthenticator) FindAgentFromRemoteAddr(hostPort string) (*config.AgentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAgentFromRemoteAddr", hostPort)
	ret0, _ := ret[0].(*config.AgentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAgentFromRemoteAddr indicates an expected call of FindAgentFromRemoteAddr.
func (mr *MockIPAuthenticatorMockRecorder) FindAgentFromRemoteAddr(hostPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAgentFromRemoteAddr", reflect.TypeOf((*MockIPAuthenticator)(nil).FindAgentFromRemoteAddr), hostPort)
}

// FindContainerNameFromRemoteAddr mocks base method.
func (m *MockIPAuthenticator) FindContainerNameFromRemoteAddr(ctx context.Context, hostPort string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindContainerNameFromRemoteAddr", ctx, hostPort)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindContainerNameFromRemoteAddr indicates an expected call of FindContainerNameFromRemoteAddr.
func (mr *MockIPAuthenticatorMockRecorder) FindContainerNameFromRemoteAddr(ctx, hostPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindContainerNameFromRemoteAddr", reflect.TypeOf((*MockIPAuthenticator)(nil).FindContainerNameFromRemoteAddr), ctx, hostPort)
}
