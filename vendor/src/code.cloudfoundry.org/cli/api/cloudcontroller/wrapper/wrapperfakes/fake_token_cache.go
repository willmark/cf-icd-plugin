// This file was generated by counterfeiter
package wrapperfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/cloudcontroller/wrapper"
)

type FakeTokenCache struct {
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct{}
	refreshTokenReturns     struct {
		result1 string
	}
	SetAccessTokenStub        func(token string)
	setAccessTokenMutex       sync.RWMutex
	setAccessTokenArgsForCall []struct {
		token string
	}
	SetRefreshTokenStub        func(token string)
	setRefreshTokenMutex       sync.RWMutex
	setRefreshTokenArgsForCall []struct {
		token string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenCache) AccessToken() string {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1
	}
}

func (fake *FakeTokenCache) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeTokenCache) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	} else {
		return fake.refreshTokenReturns.result1
	}
}

func (fake *FakeTokenCache) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeTokenCache) RefreshTokenReturns(result1 string) {
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) SetAccessToken(token string) {
	fake.setAccessTokenMutex.Lock()
	fake.setAccessTokenArgsForCall = append(fake.setAccessTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("SetAccessToken", []interface{}{token})
	fake.setAccessTokenMutex.Unlock()
	if fake.SetAccessTokenStub != nil {
		fake.SetAccessTokenStub(token)
	}
}

func (fake *FakeTokenCache) SetAccessTokenCallCount() int {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return len(fake.setAccessTokenArgsForCall)
}

func (fake *FakeTokenCache) SetAccessTokenArgsForCall(i int) string {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return fake.setAccessTokenArgsForCall[i].token
}

func (fake *FakeTokenCache) SetRefreshToken(token string) {
	fake.setRefreshTokenMutex.Lock()
	fake.setRefreshTokenArgsForCall = append(fake.setRefreshTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("SetRefreshToken", []interface{}{token})
	fake.setRefreshTokenMutex.Unlock()
	if fake.SetRefreshTokenStub != nil {
		fake.SetRefreshTokenStub(token)
	}
}

func (fake *FakeTokenCache) SetRefreshTokenCallCount() int {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return len(fake.setRefreshTokenArgsForCall)
}

func (fake *FakeTokenCache) SetRefreshTokenArgsForCall(i int) string {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return fake.setRefreshTokenArgsForCall[i].token
}

func (fake *FakeTokenCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenCache) recordInvocation(key string, args []interface{}) {
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

var _ wrapper.TokenCache = new(FakeTokenCache)
