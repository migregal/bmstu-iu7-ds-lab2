// Code generated by mockery v2.33.3. DO NOT EDIT.

package reservation

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// AddUserReservation provides a mock function with given fields: ctx, res
func (_m *MockClient) AddUserReservation(ctx context.Context, res Reservation) (string, error) {
	ret := _m.Called(ctx, res)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, Reservation) (string, error)); ok {
		return rf(ctx, res)
	}
	if rf, ok := ret.Get(0).(func(context.Context, Reservation) string); ok {
		r0 = rf(ctx, res)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, Reservation) error); ok {
		r1 = rf(ctx, res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_AddUserReservation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserReservation'
type MockClient_AddUserReservation_Call struct {
	*mock.Call
}

// AddUserReservation is a helper method to define mock.On call
//   - ctx context.Context
//   - res Reservation
func (_e *MockClient_Expecter) AddUserReservation(ctx interface{}, res interface{}) *MockClient_AddUserReservation_Call {
	return &MockClient_AddUserReservation_Call{Call: _e.mock.On("AddUserReservation", ctx, res)}
}

func (_c *MockClient_AddUserReservation_Call) Run(run func(ctx context.Context, res Reservation)) *MockClient_AddUserReservation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Reservation))
	})
	return _c
}

func (_c *MockClient_AddUserReservation_Call) Return(_a0 string, _a1 error) *MockClient_AddUserReservation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_AddUserReservation_Call) RunAndReturn(run func(context.Context, Reservation) (string, error)) *MockClient_AddUserReservation_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserReservations provides a mock function with given fields: ctx, username, status
func (_m *MockClient) GetUserReservations(ctx context.Context, username string, status string) ([]Reservation, error) {
	ret := _m.Called(ctx, username, status)

	var r0 []Reservation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]Reservation, error)); ok {
		return rf(ctx, username, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []Reservation); ok {
		r0 = rf(ctx, username, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Reservation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetUserReservations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserReservations'
type MockClient_GetUserReservations_Call struct {
	*mock.Call
}

// GetUserReservations is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - status string
func (_e *MockClient_Expecter) GetUserReservations(ctx interface{}, username interface{}, status interface{}) *MockClient_GetUserReservations_Call {
	return &MockClient_GetUserReservations_Call{Call: _e.mock.On("GetUserReservations", ctx, username, status)}
}

func (_c *MockClient_GetUserReservations_Call) Run(run func(ctx context.Context, username string, status string)) *MockClient_GetUserReservations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockClient_GetUserReservations_Call) Return(_a0 []Reservation, _a1 error) *MockClient_GetUserReservations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetUserReservations_Call) RunAndReturn(run func(context.Context, string, string) ([]Reservation, error)) *MockClient_GetUserReservations_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserReservationStatus provides a mock function with given fields: ctx, id, status
func (_m *MockClient) SetUserReservationStatus(ctx context.Context, id string, status string) error {
	ret := _m.Called(ctx, id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_SetUserReservationStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserReservationStatus'
type MockClient_SetUserReservationStatus_Call struct {
	*mock.Call
}

// SetUserReservationStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - status string
func (_e *MockClient_Expecter) SetUserReservationStatus(ctx interface{}, id interface{}, status interface{}) *MockClient_SetUserReservationStatus_Call {
	return &MockClient_SetUserReservationStatus_Call{Call: _e.mock.On("SetUserReservationStatus", ctx, id, status)}
}

func (_c *MockClient_SetUserReservationStatus_Call) Run(run func(ctx context.Context, id string, status string)) *MockClient_SetUserReservationStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockClient_SetUserReservationStatus_Call) Return(_a0 error) *MockClient_SetUserReservationStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_SetUserReservationStatus_Call) RunAndReturn(run func(context.Context, string, string) error) *MockClient_SetUserReservationStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
