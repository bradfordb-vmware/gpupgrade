// Code generated by MockGen. DO NOT EDIT.
// Source: idl/cli_to_hub.pb.go

// Package mock_idl is a generated GoMock package.
package mock_idl

import (
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gp_upgrade/idl"
	"reflect"
)

// MockCliToHubClient is a mock of CliToHubClient interface
type MockCliToHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCliToHubClientMockRecorder
}

// MockCliToHubClientMockRecorder is the mock recorder for MockCliToHubClient
type MockCliToHubClientMockRecorder struct {
	mock *MockCliToHubClient
}

// NewMockCliToHubClient creates a new mock instance
func NewMockCliToHubClient(ctrl *gomock.Controller) *MockCliToHubClient {
	mock := &MockCliToHubClient{ctrl: ctrl}
	mock.recorder = &MockCliToHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliToHubClient) EXPECT() *MockCliToHubClientMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockCliToHubClient) Ping(ctx context.Context, in *idl.PingRequest, opts ...grpc.CallOption) (*idl.PingReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*idl.PingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCliToHubClientMockRecorder) Ping(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCliToHubClient)(nil).Ping), varargs...)
}

// StatusUpgrade mocks base method
func (m *MockCliToHubClient) StatusUpgrade(ctx context.Context, in *idl.StatusUpgradeRequest, opts ...grpc.CallOption) (*idl.StatusUpgradeReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatusUpgrade", varargs...)
	ret0, _ := ret[0].(*idl.StatusUpgradeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusUpgrade indicates an expected call of StatusUpgrade
func (mr *MockCliToHubClientMockRecorder) StatusUpgrade(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusUpgrade", reflect.TypeOf((*MockCliToHubClient)(nil).StatusUpgrade), varargs...)
}

// StatusConversion mocks base method
func (m *MockCliToHubClient) StatusConversion(ctx context.Context, in *idl.StatusConversionRequest, opts ...grpc.CallOption) (*idl.StatusConversionReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatusConversion", varargs...)
	ret0, _ := ret[0].(*idl.StatusConversionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusConversion indicates an expected call of StatusConversion
func (mr *MockCliToHubClientMockRecorder) StatusConversion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusConversion", reflect.TypeOf((*MockCliToHubClient)(nil).StatusConversion), varargs...)
}

// CheckConfig mocks base method
func (m *MockCliToHubClient) CheckConfig(ctx context.Context, in *idl.CheckConfigRequest, opts ...grpc.CallOption) (*idl.CheckConfigReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckConfig", varargs...)
	ret0, _ := ret[0].(*idl.CheckConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConfig indicates an expected call of CheckConfig
func (mr *MockCliToHubClientMockRecorder) CheckConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConfig", reflect.TypeOf((*MockCliToHubClient)(nil).CheckConfig), varargs...)
}

// CheckSeginstall mocks base method
func (m *MockCliToHubClient) CheckSeginstall(ctx context.Context, in *idl.CheckSeginstallRequest, opts ...grpc.CallOption) (*idl.CheckSeginstallReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckSeginstall", varargs...)
	ret0, _ := ret[0].(*idl.CheckSeginstallReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSeginstall indicates an expected call of CheckSeginstall
func (mr *MockCliToHubClientMockRecorder) CheckSeginstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSeginstall", reflect.TypeOf((*MockCliToHubClient)(nil).CheckSeginstall), varargs...)
}

// CheckObjectCount mocks base method
func (m *MockCliToHubClient) CheckObjectCount(ctx context.Context, in *idl.CheckObjectCountRequest, opts ...grpc.CallOption) (*idl.CheckObjectCountReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckObjectCount", varargs...)
	ret0, _ := ret[0].(*idl.CheckObjectCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckObjectCount indicates an expected call of CheckObjectCount
func (mr *MockCliToHubClientMockRecorder) CheckObjectCount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckObjectCount", reflect.TypeOf((*MockCliToHubClient)(nil).CheckObjectCount), varargs...)
}

// CheckVersion mocks base method
func (m *MockCliToHubClient) CheckVersion(ctx context.Context, in *idl.CheckVersionRequest, opts ...grpc.CallOption) (*idl.CheckVersionReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckVersion", varargs...)
	ret0, _ := ret[0].(*idl.CheckVersionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion
func (mr *MockCliToHubClientMockRecorder) CheckVersion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockCliToHubClient)(nil).CheckVersion), varargs...)
}

// CheckDiskUsage mocks base method
func (m *MockCliToHubClient) CheckDiskUsage(ctx context.Context, in *idl.CheckDiskUsageRequest, opts ...grpc.CallOption) (*idl.CheckDiskUsageReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDiskUsage", varargs...)
	ret0, _ := ret[0].(*idl.CheckDiskUsageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsage indicates an expected call of CheckDiskUsage
func (mr *MockCliToHubClientMockRecorder) CheckDiskUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsage", reflect.TypeOf((*MockCliToHubClient)(nil).CheckDiskUsage), varargs...)
}

// PrepareInitCluster mocks base method
func (m *MockCliToHubClient) PrepareInitCluster(ctx context.Context, in *idl.PrepareInitClusterRequest, opts ...grpc.CallOption) (*idl.PrepareInitClusterReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareInitCluster", varargs...)
	ret0, _ := ret[0].(*idl.PrepareInitClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareInitCluster indicates an expected call of PrepareInitCluster
func (mr *MockCliToHubClientMockRecorder) PrepareInitCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareInitCluster", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareInitCluster), varargs...)
}

// PrepareShutdownClusters mocks base method
func (m *MockCliToHubClient) PrepareShutdownClusters(ctx context.Context, in *idl.PrepareShutdownClustersRequest, opts ...grpc.CallOption) (*idl.PrepareShutdownClustersReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareShutdownClusters", varargs...)
	ret0, _ := ret[0].(*idl.PrepareShutdownClustersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareShutdownClusters indicates an expected call of PrepareShutdownClusters
func (mr *MockCliToHubClientMockRecorder) PrepareShutdownClusters(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareShutdownClusters", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareShutdownClusters), varargs...)
}

// UpgradeConvertMaster mocks base method
func (m *MockCliToHubClient) UpgradeConvertMaster(ctx context.Context, in *idl.UpgradeConvertMasterRequest, opts ...grpc.CallOption) (*idl.UpgradeConvertMasterReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeConvertMaster", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeConvertMasterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertMaster indicates an expected call of UpgradeConvertMaster
func (mr *MockCliToHubClientMockRecorder) UpgradeConvertMaster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertMaster", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeConvertMaster), varargs...)
}

// PrepareStartAgents mocks base method
func (m *MockCliToHubClient) PrepareStartAgents(ctx context.Context, in *idl.PrepareStartAgentsRequest, opts ...grpc.CallOption) (*idl.PrepareStartAgentsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareStartAgents", varargs...)
	ret0, _ := ret[0].(*idl.PrepareStartAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareStartAgents indicates an expected call of PrepareStartAgents
func (mr *MockCliToHubClientMockRecorder) PrepareStartAgents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStartAgents", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareStartAgents), varargs...)
}

// UpgradeShareOids mocks base method
func (m *MockCliToHubClient) UpgradeShareOids(ctx context.Context, in *idl.UpgradeShareOidsRequest, opts ...grpc.CallOption) (*idl.UpgradeShareOidsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeShareOids", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeShareOidsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeShareOids indicates an expected call of UpgradeShareOids
func (mr *MockCliToHubClientMockRecorder) UpgradeShareOids(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeShareOids", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeShareOids), varargs...)
}

// UpgradeValidateStartCluster mocks base method
func (m *MockCliToHubClient) UpgradeValidateStartCluster(ctx context.Context, in *idl.UpgradeValidateStartClusterRequest, opts ...grpc.CallOption) (*idl.UpgradeValidateStartClusterReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeValidateStartCluster", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeValidateStartClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeValidateStartCluster indicates an expected call of UpgradeValidateStartCluster
func (mr *MockCliToHubClientMockRecorder) UpgradeValidateStartCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeValidateStartCluster", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeValidateStartCluster), varargs...)
}

// UpgradeConvertPrimaries mocks base method
func (m *MockCliToHubClient) UpgradeConvertPrimaries(ctx context.Context, in *idl.UpgradeConvertPrimariesRequest, opts ...grpc.CallOption) (*idl.UpgradeConvertPrimariesReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeConvertPrimaries", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimaries indicates an expected call of UpgradeConvertPrimaries
func (mr *MockCliToHubClientMockRecorder) UpgradeConvertPrimaries(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimaries", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeConvertPrimaries), varargs...)
}

// MockCliToHubServer is a mock of CliToHubServer interface
type MockCliToHubServer struct {
	ctrl     *gomock.Controller
	recorder *MockCliToHubServerMockRecorder
}

// MockCliToHubServerMockRecorder is the mock recorder for MockCliToHubServer
type MockCliToHubServerMockRecorder struct {
	mock *MockCliToHubServer
}

// NewMockCliToHubServer creates a new mock instance
func NewMockCliToHubServer(ctrl *gomock.Controller) *MockCliToHubServer {
	mock := &MockCliToHubServer{ctrl: ctrl}
	mock.recorder = &MockCliToHubServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliToHubServer) EXPECT() *MockCliToHubServerMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockCliToHubServer) Ping(arg0 context.Context, arg1 *idl.PingRequest) (*idl.PingReply, error) {
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(*idl.PingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCliToHubServerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCliToHubServer)(nil).Ping), arg0, arg1)
}

// StatusUpgrade mocks base method
func (m *MockCliToHubServer) StatusUpgrade(arg0 context.Context, arg1 *idl.StatusUpgradeRequest) (*idl.StatusUpgradeReply, error) {
	ret := m.ctrl.Call(m, "StatusUpgrade", arg0, arg1)
	ret0, _ := ret[0].(*idl.StatusUpgradeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusUpgrade indicates an expected call of StatusUpgrade
func (mr *MockCliToHubServerMockRecorder) StatusUpgrade(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusUpgrade", reflect.TypeOf((*MockCliToHubServer)(nil).StatusUpgrade), arg0, arg1)
}

// StatusConversion mocks base method
func (m *MockCliToHubServer) StatusConversion(arg0 context.Context, arg1 *idl.StatusConversionRequest) (*idl.StatusConversionReply, error) {
	ret := m.ctrl.Call(m, "StatusConversion", arg0, arg1)
	ret0, _ := ret[0].(*idl.StatusConversionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusConversion indicates an expected call of StatusConversion
func (mr *MockCliToHubServerMockRecorder) StatusConversion(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusConversion", reflect.TypeOf((*MockCliToHubServer)(nil).StatusConversion), arg0, arg1)
}

// CheckConfig mocks base method
func (m *MockCliToHubServer) CheckConfig(arg0 context.Context, arg1 *idl.CheckConfigRequest) (*idl.CheckConfigReply, error) {
	ret := m.ctrl.Call(m, "CheckConfig", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConfig indicates an expected call of CheckConfig
func (mr *MockCliToHubServerMockRecorder) CheckConfig(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConfig", reflect.TypeOf((*MockCliToHubServer)(nil).CheckConfig), arg0, arg1)
}

// CheckSeginstall mocks base method
func (m *MockCliToHubServer) CheckSeginstall(arg0 context.Context, arg1 *idl.CheckSeginstallRequest) (*idl.CheckSeginstallReply, error) {
	ret := m.ctrl.Call(m, "CheckSeginstall", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckSeginstallReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSeginstall indicates an expected call of CheckSeginstall
func (mr *MockCliToHubServerMockRecorder) CheckSeginstall(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSeginstall", reflect.TypeOf((*MockCliToHubServer)(nil).CheckSeginstall), arg0, arg1)
}

// CheckObjectCount mocks base method
func (m *MockCliToHubServer) CheckObjectCount(arg0 context.Context, arg1 *idl.CheckObjectCountRequest) (*idl.CheckObjectCountReply, error) {
	ret := m.ctrl.Call(m, "CheckObjectCount", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckObjectCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckObjectCount indicates an expected call of CheckObjectCount
func (mr *MockCliToHubServerMockRecorder) CheckObjectCount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckObjectCount", reflect.TypeOf((*MockCliToHubServer)(nil).CheckObjectCount), arg0, arg1)
}

// CheckVersion mocks base method
func (m *MockCliToHubServer) CheckVersion(arg0 context.Context, arg1 *idl.CheckVersionRequest) (*idl.CheckVersionReply, error) {
	ret := m.ctrl.Call(m, "CheckVersion", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckVersionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion
func (mr *MockCliToHubServerMockRecorder) CheckVersion(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockCliToHubServer)(nil).CheckVersion), arg0, arg1)
}

// CheckDiskUsage mocks base method
func (m *MockCliToHubServer) CheckDiskUsage(arg0 context.Context, arg1 *idl.CheckDiskUsageRequest) (*idl.CheckDiskUsageReply, error) {
	ret := m.ctrl.Call(m, "CheckDiskUsage", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckDiskUsageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsage indicates an expected call of CheckDiskUsage
func (mr *MockCliToHubServerMockRecorder) CheckDiskUsage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsage", reflect.TypeOf((*MockCliToHubServer)(nil).CheckDiskUsage), arg0, arg1)
}

// PrepareInitCluster mocks base method
func (m *MockCliToHubServer) PrepareInitCluster(arg0 context.Context, arg1 *idl.PrepareInitClusterRequest) (*idl.PrepareInitClusterReply, error) {
	ret := m.ctrl.Call(m, "PrepareInitCluster", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareInitClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareInitCluster indicates an expected call of PrepareInitCluster
func (mr *MockCliToHubServerMockRecorder) PrepareInitCluster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareInitCluster", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareInitCluster), arg0, arg1)
}

// PrepareShutdownClusters mocks base method
func (m *MockCliToHubServer) PrepareShutdownClusters(arg0 context.Context, arg1 *idl.PrepareShutdownClustersRequest) (*idl.PrepareShutdownClustersReply, error) {
	ret := m.ctrl.Call(m, "PrepareShutdownClusters", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareShutdownClustersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareShutdownClusters indicates an expected call of PrepareShutdownClusters
func (mr *MockCliToHubServerMockRecorder) PrepareShutdownClusters(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareShutdownClusters", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareShutdownClusters), arg0, arg1)
}

// UpgradeConvertMaster mocks base method
func (m *MockCliToHubServer) UpgradeConvertMaster(arg0 context.Context, arg1 *idl.UpgradeConvertMasterRequest) (*idl.UpgradeConvertMasterReply, error) {
	ret := m.ctrl.Call(m, "UpgradeConvertMaster", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeConvertMasterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertMaster indicates an expected call of UpgradeConvertMaster
func (mr *MockCliToHubServerMockRecorder) UpgradeConvertMaster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertMaster", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeConvertMaster), arg0, arg1)
}

// PrepareStartAgents mocks base method
func (m *MockCliToHubServer) PrepareStartAgents(arg0 context.Context, arg1 *idl.PrepareStartAgentsRequest) (*idl.PrepareStartAgentsReply, error) {
	ret := m.ctrl.Call(m, "PrepareStartAgents", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareStartAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareStartAgents indicates an expected call of PrepareStartAgents
func (mr *MockCliToHubServerMockRecorder) PrepareStartAgents(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStartAgents", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareStartAgents), arg0, arg1)
}

// UpgradeShareOids mocks base method
func (m *MockCliToHubServer) UpgradeShareOids(arg0 context.Context, arg1 *idl.UpgradeShareOidsRequest) (*idl.UpgradeShareOidsReply, error) {
	ret := m.ctrl.Call(m, "UpgradeShareOids", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeShareOidsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeShareOids indicates an expected call of UpgradeShareOids
func (mr *MockCliToHubServerMockRecorder) UpgradeShareOids(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeShareOids", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeShareOids), arg0, arg1)
}

// UpgradeValidateStartCluster mocks base method
func (m *MockCliToHubServer) UpgradeValidateStartCluster(arg0 context.Context, arg1 *idl.UpgradeValidateStartClusterRequest) (*idl.UpgradeValidateStartClusterReply, error) {
	ret := m.ctrl.Call(m, "UpgradeValidateStartCluster", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeValidateStartClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeValidateStartCluster indicates an expected call of UpgradeValidateStartCluster
func (mr *MockCliToHubServerMockRecorder) UpgradeValidateStartCluster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeValidateStartCluster", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeValidateStartCluster), arg0, arg1)
}

// UpgradeConvertPrimaries mocks base method
func (m *MockCliToHubServer) UpgradeConvertPrimaries(arg0 context.Context, arg1 *idl.UpgradeConvertPrimariesRequest) (*idl.UpgradeConvertPrimariesReply, error) {
	ret := m.ctrl.Call(m, "UpgradeConvertPrimaries", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimaries indicates an expected call of UpgradeConvertPrimaries
func (mr *MockCliToHubServerMockRecorder) UpgradeConvertPrimaries(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimaries", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeConvertPrimaries), arg0, arg1)
}
