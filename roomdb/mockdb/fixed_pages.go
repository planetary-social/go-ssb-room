// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

// Code generated by counterfeiter. DO NOT EDIT.
package mockdb

import (
	"context"
	"sync"

	"github.com/ssbc/go-ssb-room/v2/roomdb"
)

type FakePinnedNoticesService struct {
	GetStub        func(context.Context, roomdb.PinnedNoticeName, string) (*roomdb.Notice, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 roomdb.PinnedNoticeName
		arg3 string
	}
	getReturns struct {
		result1 *roomdb.Notice
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *roomdb.Notice
		result2 error
	}
	ListStub        func(context.Context) (roomdb.PinnedNotices, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 roomdb.PinnedNotices
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 roomdb.PinnedNotices
		result2 error
	}
	SetStub        func(context.Context, roomdb.PinnedNoticeName, int64) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 context.Context
		arg2 roomdb.PinnedNoticeName
		arg3 int64
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePinnedNoticesService) Get(arg1 context.Context, arg2 roomdb.PinnedNoticeName, arg3 string) (*roomdb.Notice, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 roomdb.PinnedNoticeName
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePinnedNoticesService) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakePinnedNoticesService) GetCalls(stub func(context.Context, roomdb.PinnedNoticeName, string) (*roomdb.Notice, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakePinnedNoticesService) GetArgsForCall(i int) (context.Context, roomdb.PinnedNoticeName, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePinnedNoticesService) GetReturns(result1 *roomdb.Notice, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *roomdb.Notice
		result2 error
	}{result1, result2}
}

func (fake *FakePinnedNoticesService) GetReturnsOnCall(i int, result1 *roomdb.Notice, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *roomdb.Notice
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *roomdb.Notice
		result2 error
	}{result1, result2}
}

func (fake *FakePinnedNoticesService) List(arg1 context.Context) (roomdb.PinnedNotices, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePinnedNoticesService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePinnedNoticesService) ListCalls(stub func(context.Context) (roomdb.PinnedNotices, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakePinnedNoticesService) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePinnedNoticesService) ListReturns(result1 roomdb.PinnedNotices, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 roomdb.PinnedNotices
		result2 error
	}{result1, result2}
}

func (fake *FakePinnedNoticesService) ListReturnsOnCall(i int, result1 roomdb.PinnedNotices, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 roomdb.PinnedNotices
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 roomdb.PinnedNotices
		result2 error
	}{result1, result2}
}

func (fake *FakePinnedNoticesService) Set(arg1 context.Context, arg2 roomdb.PinnedNoticeName, arg3 int64) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 context.Context
		arg2 roomdb.PinnedNoticeName
		arg3 int64
	}{arg1, arg2, arg3})
	stub := fake.SetStub
	fakeReturns := fake.setReturns
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePinnedNoticesService) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakePinnedNoticesService) SetCalls(stub func(context.Context, roomdb.PinnedNoticeName, int64) error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakePinnedNoticesService) SetArgsForCall(i int) (context.Context, roomdb.PinnedNoticeName, int64) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePinnedNoticesService) SetReturns(result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePinnedNoticesService) SetReturnsOnCall(i int, result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePinnedNoticesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePinnedNoticesService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ roomdb.PinnedNoticesService = new(FakePinnedNoticesService)
