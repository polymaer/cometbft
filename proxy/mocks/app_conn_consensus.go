// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v2 "github.com/cometbft/cometbft/api/cometbft/abci/v2"

	v3 "github.com/cometbft/cometbft/api/cometbft/abci/v3"

	v4 "github.com/cometbft/cometbft/api/cometbft/abci/v4"
)

// AppConnConsensus is an autogenerated mock type for the AppConnConsensus type
type AppConnConsensus struct {
	mock.Mock
}

// Commit provides a mock function with given fields: _a0
func (_m *AppConnConsensus) Commit(_a0 context.Context) (*v3.ResponseCommit, error) {
	ret := _m.Called(_a0)

	var r0 *v3.ResponseCommit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v3.ResponseCommit, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v3.ResponseCommit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v3.ResponseCommit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnConsensus) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendVote provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) ExtendVote(_a0 context.Context, _a1 *v4.RequestExtendVote) (*v3.ResponseExtendVote, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v3.ResponseExtendVote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestExtendVote) (*v3.ResponseExtendVote, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestExtendVote) *v3.ResponseExtendVote); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v3.ResponseExtendVote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v4.RequestExtendVote) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeBlock provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) FinalizeBlock(_a0 context.Context, _a1 *v4.RequestFinalizeBlock) (*v3.ResponseFinalizeBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v3.ResponseFinalizeBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestFinalizeBlock) (*v3.ResponseFinalizeBlock, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestFinalizeBlock) *v3.ResponseFinalizeBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v3.ResponseFinalizeBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v4.RequestFinalizeBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChain provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) InitChain(_a0 context.Context, _a1 *v3.RequestInitChain) (*v3.ResponseInitChain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v3.ResponseInitChain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v3.RequestInitChain) (*v3.ResponseInitChain, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v3.RequestInitChain) *v3.ResponseInitChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v3.ResponseInitChain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v3.RequestInitChain) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareProposal provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) PrepareProposal(_a0 context.Context, _a1 *v4.RequestPrepareProposal) (*v2.ResponsePrepareProposal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v2.ResponsePrepareProposal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestPrepareProposal) (*v2.ResponsePrepareProposal, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestPrepareProposal) *v2.ResponsePrepareProposal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.ResponsePrepareProposal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v4.RequestPrepareProposal) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessProposal provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) ProcessProposal(_a0 context.Context, _a1 *v4.RequestProcessProposal) (*v4.ResponseProcessProposal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v4.ResponseProcessProposal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestProcessProposal) (*v4.ResponseProcessProposal, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v4.RequestProcessProposal) *v4.ResponseProcessProposal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.ResponseProcessProposal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v4.RequestProcessProposal) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyVoteExtension provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) VerifyVoteExtension(_a0 context.Context, _a1 *v3.RequestVerifyVoteExtension) (*v4.ResponseVerifyVoteExtension, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v4.ResponseVerifyVoteExtension
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v3.RequestVerifyVoteExtension) (*v4.ResponseVerifyVoteExtension, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v3.RequestVerifyVoteExtension) *v4.ResponseVerifyVoteExtension); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.ResponseVerifyVoteExtension)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v3.RequestVerifyVoteExtension) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAppConnConsensus creates a new instance of AppConnConsensus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppConnConsensus(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppConnConsensus {
	mock := &AppConnConsensus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
