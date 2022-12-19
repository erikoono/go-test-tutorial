// Code generated by MockGen. DO NOT EDIT.
// Source: ./chapters/chapter4/db/interface.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	structs "github.com/dip-dev/go-test-tutorial/mysql/structs"
	gomock "github.com/golang/mock/gomock"
)

// MockSelecters is a mock of Selecters interface.
type MockSelecters struct {
	ctrl     *gomock.Controller
	recorder *MockSelectersMockRecorder
}

// MockSelectersMockRecorder is the mock recorder for MockSelecters.
type MockSelectersMockRecorder struct {
	mock *MockSelecters
}

// NewMockSelecters creates a new mock instance.
func NewMockSelecters(ctrl *gomock.Controller) *MockSelecters {
	mock := &MockSelecters{ctrl: ctrl}
	mock.recorder = &MockSelectersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelecters) EXPECT() *MockSelectersMockRecorder {
	return m.recorder
}

// SelectPrefectures mocks base method.
func (m *MockSelecters) SelectPrefectures(ctx context.Context, area string) ([]structs.MPrefecture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPrefectures", ctx, area)
	ret0, _ := ret[0].([]structs.MPrefecture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectPrefectures indicates an expected call of SelectPrefectures.
func (mr *MockSelectersMockRecorder) SelectPrefectures(ctx, area interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPrefectures", reflect.TypeOf((*MockSelecters)(nil).SelectPrefectures), ctx, area)
}
