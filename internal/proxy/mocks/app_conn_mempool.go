// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	abciclient "github.com/tendermint/tendermint/abci/client"

	mock "github.com/stretchr/testify/mock"

	types "github.com/tendermint/tendermint/abci/types"
)

// AppConnMempool is an autogenerated mock type for the AppConnMempool type
type AppConnMempool struct {
	mock.Mock
}

// CheckTx provides a mock function with given fields: _a0, _a1
func (_m *AppConnMempool) CheckTx(_a0 context.Context, _a1 types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseCheckTx
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestCheckTx) *types.ResponseCheckTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestCheckTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *AppConnMempool) CheckTxAsync(_a0 context.Context, _a1 types.RequestCheckTx) (*abciclient.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abciclient.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestCheckTx) *abciclient.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abciclient.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestCheckTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnMempool) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields: _a0
func (_m *AppConnMempool) Flush(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushAsync provides a mock function with given fields: _a0
func (_m *AppConnMempool) FlushAsync(_a0 context.Context) (*abciclient.ReqRes, error) {
	ret := _m.Called(_a0)

	var r0 *abciclient.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context) *abciclient.ReqRes); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abciclient.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetResponseCallback provides a mock function with given fields: _a0
func (_m *AppConnMempool) SetResponseCallback(_a0 abciclient.Callback) {
	_m.Called(_a0)
}
