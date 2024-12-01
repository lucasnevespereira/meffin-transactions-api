// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "meffin-transactions-api/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// ICategoryService is an autogenerated mock type for the ICategoryService type
type ICategoryService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *ICategoryService) Create(ctx context.Context, request models.CreateCategoryRequest) (*models.Category, error) {
	ret := _m.Called(ctx, request)

	var r0 *models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateCategoryRequest) (*models.Category, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateCategoryRequest) *models.Category); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.CreateCategoryRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCategory provides a mock function with given fields: ctx, categoryID
func (_m *ICategoryService) DeleteCategory(ctx context.Context, categoryID string) error {
	ret := _m.Called(ctx, categoryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserCategories provides a mock function with given fields: ctx, userId
func (_m *ICategoryService) GetUserCategories(ctx context.Context, userId string) ([]*models.Category, error) {
	ret := _m.Called(ctx, userId)

	var r0 []*models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.Category, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.Category); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: ctx, category
func (_m *ICategoryService) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	ret := _m.Called(ctx, category)

	var r0 *models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Category) (*models.Category, error)); ok {
		return rf(ctx, category)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Category) *models.Category); ok {
		r0 = rf(ctx, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Category) error); ok {
		r1 = rf(ctx, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICategoryService creates a new instance of ICategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICategoryService {
	mock := &ICategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
