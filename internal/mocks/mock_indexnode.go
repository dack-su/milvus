// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	commonpb "github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	clientv3 "go.etcd.io/etcd/client/v3"

	indexpb "github.com/milvus-io/milvus/internal/proto/indexpb"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"

	mock "github.com/stretchr/testify/mock"
)

// MockIndexNode is an autogenerated mock type for the IndexNodeComponent type
type MockIndexNode struct {
	mock.Mock
}

type MockIndexNode_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIndexNode) EXPECT() *MockIndexNode_Expecter {
	return &MockIndexNode_Expecter{mock: &_m.Mock}
}

// CreateJob provides a mock function with given fields: _a0, _a1
func (_m *MockIndexNode) CreateJob(_a0 context.Context, _a1 *indexpb.CreateJobRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.CreateJobRequest) (*commonpb.Status, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.CreateJobRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.CreateJobRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type MockIndexNode_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *indexpb.CreateJobRequest
func (_e *MockIndexNode_Expecter) CreateJob(_a0 interface{}, _a1 interface{}) *MockIndexNode_CreateJob_Call {
	return &MockIndexNode_CreateJob_Call{Call: _e.mock.On("CreateJob", _a0, _a1)}
}

func (_c *MockIndexNode_CreateJob_Call) Run(run func(_a0 context.Context, _a1 *indexpb.CreateJobRequest)) *MockIndexNode_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.CreateJobRequest))
	})
	return _c
}

func (_c *MockIndexNode_CreateJob_Call) Return(_a0 *commonpb.Status, _a1 error) *MockIndexNode_CreateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_CreateJob_Call) RunAndReturn(run func(context.Context, *indexpb.CreateJobRequest) (*commonpb.Status, error)) *MockIndexNode_CreateJob_Call {
	_c.Call.Return(run)
	return _c
}

// DropJobs provides a mock function with given fields: _a0, _a1
func (_m *MockIndexNode) DropJobs(_a0 context.Context, _a1 *indexpb.DropJobsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.DropJobsRequest) (*commonpb.Status, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.DropJobsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.DropJobsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_DropJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropJobs'
type MockIndexNode_DropJobs_Call struct {
	*mock.Call
}

// DropJobs is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *indexpb.DropJobsRequest
func (_e *MockIndexNode_Expecter) DropJobs(_a0 interface{}, _a1 interface{}) *MockIndexNode_DropJobs_Call {
	return &MockIndexNode_DropJobs_Call{Call: _e.mock.On("DropJobs", _a0, _a1)}
}

func (_c *MockIndexNode_DropJobs_Call) Run(run func(_a0 context.Context, _a1 *indexpb.DropJobsRequest)) *MockIndexNode_DropJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.DropJobsRequest))
	})
	return _c
}

func (_c *MockIndexNode_DropJobs_Call) Return(_a0 *commonpb.Status, _a1 error) *MockIndexNode_DropJobs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_DropJobs_Call) RunAndReturn(run func(context.Context, *indexpb.DropJobsRequest) (*commonpb.Status, error)) *MockIndexNode_DropJobs_Call {
	_c.Call.Return(run)
	return _c
}

// GetAddress provides a mock function with given fields:
func (_m *MockIndexNode) GetAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockIndexNode_GetAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddress'
type MockIndexNode_GetAddress_Call struct {
	*mock.Call
}

// GetAddress is a helper method to define mock.On call
func (_e *MockIndexNode_Expecter) GetAddress() *MockIndexNode_GetAddress_Call {
	return &MockIndexNode_GetAddress_Call{Call: _e.mock.On("GetAddress")}
}

func (_c *MockIndexNode_GetAddress_Call) Run(run func()) *MockIndexNode_GetAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexNode_GetAddress_Call) Return(_a0 string) *MockIndexNode_GetAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIndexNode_GetAddress_Call) RunAndReturn(run func() string) *MockIndexNode_GetAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetComponentStates provides a mock function with given fields: ctx
func (_m *MockIndexNode) GetComponentStates(ctx context.Context) (*milvuspb.ComponentStates, error) {
	ret := _m.Called(ctx)

	var r0 *milvuspb.ComponentStates
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*milvuspb.ComponentStates, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *milvuspb.ComponentStates); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.ComponentStates)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_GetComponentStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComponentStates'
type MockIndexNode_GetComponentStates_Call struct {
	*mock.Call
}

// GetComponentStates is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIndexNode_Expecter) GetComponentStates(ctx interface{}) *MockIndexNode_GetComponentStates_Call {
	return &MockIndexNode_GetComponentStates_Call{Call: _e.mock.On("GetComponentStates", ctx)}
}

func (_c *MockIndexNode_GetComponentStates_Call) Run(run func(ctx context.Context)) *MockIndexNode_GetComponentStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIndexNode_GetComponentStates_Call) Return(_a0 *milvuspb.ComponentStates, _a1 error) *MockIndexNode_GetComponentStates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_GetComponentStates_Call) RunAndReturn(run func(context.Context) (*milvuspb.ComponentStates, error)) *MockIndexNode_GetComponentStates_Call {
	_c.Call.Return(run)
	return _c
}

// GetJobStats provides a mock function with given fields: _a0, _a1
func (_m *MockIndexNode) GetJobStats(_a0 context.Context, _a1 *indexpb.GetJobStatsRequest) (*indexpb.GetJobStatsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *indexpb.GetJobStatsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetJobStatsRequest) (*indexpb.GetJobStatsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetJobStatsRequest) *indexpb.GetJobStatsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.GetJobStatsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.GetJobStatsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_GetJobStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJobStats'
type MockIndexNode_GetJobStats_Call struct {
	*mock.Call
}

// GetJobStats is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *indexpb.GetJobStatsRequest
func (_e *MockIndexNode_Expecter) GetJobStats(_a0 interface{}, _a1 interface{}) *MockIndexNode_GetJobStats_Call {
	return &MockIndexNode_GetJobStats_Call{Call: _e.mock.On("GetJobStats", _a0, _a1)}
}

func (_c *MockIndexNode_GetJobStats_Call) Run(run func(_a0 context.Context, _a1 *indexpb.GetJobStatsRequest)) *MockIndexNode_GetJobStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.GetJobStatsRequest))
	})
	return _c
}

func (_c *MockIndexNode_GetJobStats_Call) Return(_a0 *indexpb.GetJobStatsResponse, _a1 error) *MockIndexNode_GetJobStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_GetJobStats_Call) RunAndReturn(run func(context.Context, *indexpb.GetJobStatsRequest) (*indexpb.GetJobStatsResponse, error)) *MockIndexNode_GetJobStats_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetrics provides a mock function with given fields: ctx, req
func (_m *MockIndexNode) GetMetrics(ctx context.Context, req *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *milvuspb.GetMetricsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest) *milvuspb.GetMetricsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.GetMetricsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetMetricsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_GetMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetrics'
type MockIndexNode_GetMetrics_Call struct {
	*mock.Call
}

// GetMetrics is a helper method to define mock.On call
//   - ctx context.Context
//   - req *milvuspb.GetMetricsRequest
func (_e *MockIndexNode_Expecter) GetMetrics(ctx interface{}, req interface{}) *MockIndexNode_GetMetrics_Call {
	return &MockIndexNode_GetMetrics_Call{Call: _e.mock.On("GetMetrics", ctx, req)}
}

func (_c *MockIndexNode_GetMetrics_Call) Run(run func(ctx context.Context, req *milvuspb.GetMetricsRequest)) *MockIndexNode_GetMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*milvuspb.GetMetricsRequest))
	})
	return _c
}

func (_c *MockIndexNode_GetMetrics_Call) Return(_a0 *milvuspb.GetMetricsResponse, _a1 error) *MockIndexNode_GetMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_GetMetrics_Call) RunAndReturn(run func(context.Context, *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error)) *MockIndexNode_GetMetrics_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatisticsChannel provides a mock function with given fields: ctx
func (_m *MockIndexNode) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	ret := _m.Called(ctx)

	var r0 *milvuspb.StringResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*milvuspb.StringResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *milvuspb.StringResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_GetStatisticsChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatisticsChannel'
type MockIndexNode_GetStatisticsChannel_Call struct {
	*mock.Call
}

// GetStatisticsChannel is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIndexNode_Expecter) GetStatisticsChannel(ctx interface{}) *MockIndexNode_GetStatisticsChannel_Call {
	return &MockIndexNode_GetStatisticsChannel_Call{Call: _e.mock.On("GetStatisticsChannel", ctx)}
}

func (_c *MockIndexNode_GetStatisticsChannel_Call) Run(run func(ctx context.Context)) *MockIndexNode_GetStatisticsChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIndexNode_GetStatisticsChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockIndexNode_GetStatisticsChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_GetStatisticsChannel_Call) RunAndReturn(run func(context.Context) (*milvuspb.StringResponse, error)) *MockIndexNode_GetStatisticsChannel_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields:
func (_m *MockIndexNode) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexNode_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockIndexNode_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *MockIndexNode_Expecter) Init() *MockIndexNode_Init_Call {
	return &MockIndexNode_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *MockIndexNode_Init_Call) Run(run func()) *MockIndexNode_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexNode_Init_Call) Return(_a0 error) *MockIndexNode_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIndexNode_Init_Call) RunAndReturn(run func() error) *MockIndexNode_Init_Call {
	_c.Call.Return(run)
	return _c
}

// QueryJobs provides a mock function with given fields: _a0, _a1
func (_m *MockIndexNode) QueryJobs(_a0 context.Context, _a1 *indexpb.QueryJobsRequest) (*indexpb.QueryJobsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *indexpb.QueryJobsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.QueryJobsRequest) (*indexpb.QueryJobsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.QueryJobsRequest) *indexpb.QueryJobsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.QueryJobsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.QueryJobsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_QueryJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryJobs'
type MockIndexNode_QueryJobs_Call struct {
	*mock.Call
}

// QueryJobs is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *indexpb.QueryJobsRequest
func (_e *MockIndexNode_Expecter) QueryJobs(_a0 interface{}, _a1 interface{}) *MockIndexNode_QueryJobs_Call {
	return &MockIndexNode_QueryJobs_Call{Call: _e.mock.On("QueryJobs", _a0, _a1)}
}

func (_c *MockIndexNode_QueryJobs_Call) Run(run func(_a0 context.Context, _a1 *indexpb.QueryJobsRequest)) *MockIndexNode_QueryJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.QueryJobsRequest))
	})
	return _c
}

func (_c *MockIndexNode_QueryJobs_Call) Return(_a0 *indexpb.QueryJobsResponse, _a1 error) *MockIndexNode_QueryJobs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_QueryJobs_Call) RunAndReturn(run func(context.Context, *indexpb.QueryJobsRequest) (*indexpb.QueryJobsResponse, error)) *MockIndexNode_QueryJobs_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields:
func (_m *MockIndexNode) Register() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexNode_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockIndexNode_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
func (_e *MockIndexNode_Expecter) Register() *MockIndexNode_Register_Call {
	return &MockIndexNode_Register_Call{Call: _e.mock.On("Register")}
}

func (_c *MockIndexNode_Register_Call) Run(run func()) *MockIndexNode_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexNode_Register_Call) Return(_a0 error) *MockIndexNode_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIndexNode_Register_Call) RunAndReturn(run func() error) *MockIndexNode_Register_Call {
	_c.Call.Return(run)
	return _c
}

// SetAddress provides a mock function with given fields: address
func (_m *MockIndexNode) SetAddress(address string) {
	_m.Called(address)
}

// MockIndexNode_SetAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAddress'
type MockIndexNode_SetAddress_Call struct {
	*mock.Call
}

// SetAddress is a helper method to define mock.On call
//   - address string
func (_e *MockIndexNode_Expecter) SetAddress(address interface{}) *MockIndexNode_SetAddress_Call {
	return &MockIndexNode_SetAddress_Call{Call: _e.mock.On("SetAddress", address)}
}

func (_c *MockIndexNode_SetAddress_Call) Run(run func(address string)) *MockIndexNode_SetAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIndexNode_SetAddress_Call) Return() *MockIndexNode_SetAddress_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIndexNode_SetAddress_Call) RunAndReturn(run func(string)) *MockIndexNode_SetAddress_Call {
	_c.Call.Return(run)
	return _c
}

// SetEtcdClient provides a mock function with given fields: etcdClient
func (_m *MockIndexNode) SetEtcdClient(etcdClient *clientv3.Client) {
	_m.Called(etcdClient)
}

// MockIndexNode_SetEtcdClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEtcdClient'
type MockIndexNode_SetEtcdClient_Call struct {
	*mock.Call
}

// SetEtcdClient is a helper method to define mock.On call
//   - etcdClient *clientv3.Client
func (_e *MockIndexNode_Expecter) SetEtcdClient(etcdClient interface{}) *MockIndexNode_SetEtcdClient_Call {
	return &MockIndexNode_SetEtcdClient_Call{Call: _e.mock.On("SetEtcdClient", etcdClient)}
}

func (_c *MockIndexNode_SetEtcdClient_Call) Run(run func(etcdClient *clientv3.Client)) *MockIndexNode_SetEtcdClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*clientv3.Client))
	})
	return _c
}

func (_c *MockIndexNode_SetEtcdClient_Call) Return() *MockIndexNode_SetEtcdClient_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIndexNode_SetEtcdClient_Call) RunAndReturn(run func(*clientv3.Client)) *MockIndexNode_SetEtcdClient_Call {
	_c.Call.Return(run)
	return _c
}

// ShowConfigurations provides a mock function with given fields: ctx, req
func (_m *MockIndexNode) ShowConfigurations(ctx context.Context, req *internalpb.ShowConfigurationsRequest) (*internalpb.ShowConfigurationsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *internalpb.ShowConfigurationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.ShowConfigurationsRequest) (*internalpb.ShowConfigurationsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.ShowConfigurationsRequest) *internalpb.ShowConfigurationsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.ShowConfigurationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.ShowConfigurationsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexNode_ShowConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShowConfigurations'
type MockIndexNode_ShowConfigurations_Call struct {
	*mock.Call
}

// ShowConfigurations is a helper method to define mock.On call
//   - ctx context.Context
//   - req *internalpb.ShowConfigurationsRequest
func (_e *MockIndexNode_Expecter) ShowConfigurations(ctx interface{}, req interface{}) *MockIndexNode_ShowConfigurations_Call {
	return &MockIndexNode_ShowConfigurations_Call{Call: _e.mock.On("ShowConfigurations", ctx, req)}
}

func (_c *MockIndexNode_ShowConfigurations_Call) Run(run func(ctx context.Context, req *internalpb.ShowConfigurationsRequest)) *MockIndexNode_ShowConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.ShowConfigurationsRequest))
	})
	return _c
}

func (_c *MockIndexNode_ShowConfigurations_Call) Return(_a0 *internalpb.ShowConfigurationsResponse, _a1 error) *MockIndexNode_ShowConfigurations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIndexNode_ShowConfigurations_Call) RunAndReturn(run func(context.Context, *internalpb.ShowConfigurationsRequest) (*internalpb.ShowConfigurationsResponse, error)) *MockIndexNode_ShowConfigurations_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockIndexNode) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexNode_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockIndexNode_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockIndexNode_Expecter) Start() *MockIndexNode_Start_Call {
	return &MockIndexNode_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockIndexNode_Start_Call) Run(run func()) *MockIndexNode_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexNode_Start_Call) Return(_a0 error) *MockIndexNode_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIndexNode_Start_Call) RunAndReturn(run func() error) *MockIndexNode_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockIndexNode) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexNode_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockIndexNode_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockIndexNode_Expecter) Stop() *MockIndexNode_Stop_Call {
	return &MockIndexNode_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockIndexNode_Stop_Call) Run(run func()) *MockIndexNode_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexNode_Stop_Call) Return(_a0 error) *MockIndexNode_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIndexNode_Stop_Call) RunAndReturn(run func() error) *MockIndexNode_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStateCode provides a mock function with given fields: stateCode
func (_m *MockIndexNode) UpdateStateCode(stateCode commonpb.StateCode) {
	_m.Called(stateCode)
}

// MockIndexNode_UpdateStateCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStateCode'
type MockIndexNode_UpdateStateCode_Call struct {
	*mock.Call
}

// UpdateStateCode is a helper method to define mock.On call
//   - stateCode commonpb.StateCode
func (_e *MockIndexNode_Expecter) UpdateStateCode(stateCode interface{}) *MockIndexNode_UpdateStateCode_Call {
	return &MockIndexNode_UpdateStateCode_Call{Call: _e.mock.On("UpdateStateCode", stateCode)}
}

func (_c *MockIndexNode_UpdateStateCode_Call) Run(run func(stateCode commonpb.StateCode)) *MockIndexNode_UpdateStateCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(commonpb.StateCode))
	})
	return _c
}

func (_c *MockIndexNode_UpdateStateCode_Call) Return() *MockIndexNode_UpdateStateCode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIndexNode_UpdateStateCode_Call) RunAndReturn(run func(commonpb.StateCode)) *MockIndexNode_UpdateStateCode_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIndexNode creates a new instance of MockIndexNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIndexNode(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIndexNode {
	mock := &MockIndexNode{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
