// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	s3 "github.com/igefined/nftique/pkg/s3"
)

// MockS3 is a mock of S3 interface.
type MockS3 struct {
	ctrl     *gomock.Controller
	recorder *MockS3MockRecorder
}

// MockS3MockRecorder is the mock recorder for MockS3.
type MockS3MockRecorder struct {
	mock *MockS3
}

// NewMockS3 creates a new mock instance.
func NewMockS3(ctrl *gomock.Controller) *MockS3 {
	mock := &MockS3{ctrl: ctrl}
	mock.recorder = &MockS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3) EXPECT() *MockS3MockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockS3) List(ctx context.Context, filename string) ([]*s3.Media, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filename)
	ret0, _ := ret[0].([]*s3.Media)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockS3MockRecorder) List(ctx, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockS3)(nil).List), ctx, filename)
}

// Store mocks base method.
func (m *MockS3) Store(ctx context.Context, filename string, contentBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, filename, contentBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockS3MockRecorder) Store(ctx, filename, contentBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockS3)(nil).Store), ctx, filename, contentBytes)
}
