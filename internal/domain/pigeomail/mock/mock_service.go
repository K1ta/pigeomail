// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pigeomail "pigeomail/internal/domain/pigeomail"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// PrepareCreateEmail mocks base method
func (m *MockService) PrepareCreateEmail(ctx context.Context, chatID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCreateEmail", ctx, chatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareCreateEmail indicates an expected call of PrepareCreateEmail
func (mr *MockServiceMockRecorder) PrepareCreateEmail(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCreateEmail", reflect.TypeOf((*MockService)(nil).PrepareCreateEmail), ctx, chatID)
}

// CreateEmail mocks base method
func (m *MockService) CreateEmail(ctx context.Context, email pigeomail.EMail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmail indicates an expected call of CreateEmail
func (mr *MockServiceMockRecorder) CreateEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmail", reflect.TypeOf((*MockService)(nil).CreateEmail), ctx, email)
}

// GetChatIDByEmail mocks base method
func (m *MockService) GetChatIDByEmail(ctx context.Context, email string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatIDByEmail", ctx, email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatIDByEmail indicates an expected call of GetChatIDByEmail
func (mr *MockServiceMockRecorder) GetChatIDByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatIDByEmail", reflect.TypeOf((*MockService)(nil).GetChatIDByEmail), ctx, email)
}

// GetEmailByChatID mocks base method
func (m *MockService) GetEmailByChatID(ctx context.Context, chatID int64) (pigeomail.EMail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailByChatID", ctx, chatID)
	ret0, _ := ret[0].(pigeomail.EMail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailByChatID indicates an expected call of GetEmailByChatID
func (mr *MockServiceMockRecorder) GetEmailByChatID(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailByChatID", reflect.TypeOf((*MockService)(nil).GetEmailByChatID), ctx, chatID)
}

// PrepareDeleteEmail mocks base method
func (m *MockService) PrepareDeleteEmail(ctx context.Context, chatID int64) (pigeomail.EMail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareDeleteEmail", ctx, chatID)
	ret0, _ := ret[0].(pigeomail.EMail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareDeleteEmail indicates an expected call of PrepareDeleteEmail
func (mr *MockServiceMockRecorder) PrepareDeleteEmail(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareDeleteEmail", reflect.TypeOf((*MockService)(nil).PrepareDeleteEmail), ctx, chatID)
}

// CancelDeleteEmail mocks base method
func (m *MockService) CancelDeleteEmail(ctx context.Context, chatID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelDeleteEmail", ctx, chatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelDeleteEmail indicates an expected call of CancelDeleteEmail
func (mr *MockServiceMockRecorder) CancelDeleteEmail(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelDeleteEmail", reflect.TypeOf((*MockService)(nil).CancelDeleteEmail), ctx, chatID)
}

// DeleteEmail mocks base method
func (m *MockService) DeleteEmail(ctx context.Context, chatID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmail", ctx, chatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmail indicates an expected call of DeleteEmail
func (mr *MockServiceMockRecorder) DeleteEmail(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmail", reflect.TypeOf((*MockService)(nil).DeleteEmail), ctx, chatID)
}
