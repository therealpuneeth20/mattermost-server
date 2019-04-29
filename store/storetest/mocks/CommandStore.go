// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// CommandStore is an autogenerated mock type for the CommandStore type
type CommandStore struct {
	mock.Mock
}

// AnalyticsCommandCount provides a mock function with given fields: teamId
func (_m *CommandStore) AnalyticsCommandCount(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: commandId, time
func (_m *CommandStore) Delete(commandId string, time int64) store.StoreChannel {
	ret := _m.Called(commandId, time)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(commandId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *CommandStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByTeam provides a mock function with given fields: teamId
func (_m *CommandStore) GetByTeam(teamId string) ([]*model.Command, *model.AppError) {
	ret := _m.Called(teamId)

	var r0 []*model.Command
	if rf, ok := ret.Get(0).(func(string) []*model.Command); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Command)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(teamId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByTrigger provides a mock function with given fields: teamId, trigger
func (_m *CommandStore) GetByTrigger(teamId string, trigger string) (*model.Command, *model.AppError) {
	ret := _m.Called(teamId, trigger)

	var r0 *model.Command
	if rf, ok := ret.Get(0).(func(string, string) *model.Command); ok {
		r0 = rf(teamId, trigger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Command)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(teamId, trigger)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// PermanentDeleteByTeam provides a mock function with given fields: teamId
func (_m *CommandStore) PermanentDeleteByTeam(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *CommandStore) PermanentDeleteByUser(userId string) *model.AppError {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Save provides a mock function with given fields: webhook
func (_m *CommandStore) Save(webhook *model.Command) store.StoreChannel {
	ret := _m.Called(webhook)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Command) store.StoreChannel); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: hook
func (_m *CommandStore) Update(hook *model.Command) store.StoreChannel {
	ret := _m.Called(hook)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Command) store.StoreChannel); ok {
		r0 = rf(hook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
