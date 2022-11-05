// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	mock "github.com/stretchr/testify/mock"

	models "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: data
func (_m *AuthRepository) ChangePassword(data *dto.ChangePReq) int {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(*dto.ChangePReq) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetCode provides a mock function with given fields: email
func (_m *AuthRepository) GetCode(email string) (*models.User, int, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MatchingCredential provides a mock function with given fields: email, password
func (_m *AuthRepository) MatchingCredential(email string, password string) (*models.User, error, bool) {
	ret := _m.Called(email, password)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string, string) *models.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func(string, string) bool); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: user, cr
func (_m *AuthRepository) Register(user *models.User, cr int) (*models.User, error) {
	ret := _m.Called(user, cr)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*models.User, int) *models.User); ok {
		r0 = rf(user, cr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User, int) error); ok {
		r1 = rf(user, cr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository(t mockConstructorTestingTNewAuthRepository) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
