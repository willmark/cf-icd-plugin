// This file was generated by counterfeiter
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeTasksActor struct {
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationTasksStub        func(appGUID string, sortOrder v3action.SortOrder) ([]v3action.Task, v3action.Warnings, error)
	getApplicationTasksMutex       sync.RWMutex
	getApplicationTasksArgsForCall []struct {
		appGUID   string
		sortOrder v3action.SortOrder
	}
	getApplicationTasksReturns struct {
		result1 []v3action.Task
		result2 v3action.Warnings
		result3 error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	} else {
		return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
	}
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) GetApplicationTasks(appGUID string, sortOrder v3action.SortOrder) ([]v3action.Task, v3action.Warnings, error) {
	fake.getApplicationTasksMutex.Lock()
	fake.getApplicationTasksArgsForCall = append(fake.getApplicationTasksArgsForCall, struct {
		appGUID   string
		sortOrder v3action.SortOrder
	}{appGUID, sortOrder})
	fake.recordInvocation("GetApplicationTasks", []interface{}{appGUID, sortOrder})
	fake.getApplicationTasksMutex.Unlock()
	if fake.GetApplicationTasksStub != nil {
		return fake.GetApplicationTasksStub(appGUID, sortOrder)
	} else {
		return fake.getApplicationTasksReturns.result1, fake.getApplicationTasksReturns.result2, fake.getApplicationTasksReturns.result3
	}
}

func (fake *FakeTasksActor) GetApplicationTasksCallCount() int {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return len(fake.getApplicationTasksArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationTasksArgsForCall(i int) (string, v3action.SortOrder) {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return fake.getApplicationTasksArgsForCall[i].appGUID, fake.getApplicationTasksArgsForCall[i].sortOrder
}

func (fake *FakeTasksActor) GetApplicationTasksReturns(result1 []v3action.Task, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationTasksStub = nil
	fake.getApplicationTasksReturns = struct {
		result1 []v3action.Task
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	} else {
		return fake.cloudControllerAPIVersionReturns.result1
	}
}

func (fake *FakeTasksActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeTasksActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTasksActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTasksActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.TasksActor = new(FakeTasksActor)
