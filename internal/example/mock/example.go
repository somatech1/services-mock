// Code generated by MockGen. DO NOT EDIT.
// Source: internal/example/example.go
//
// Generated by this command:
//
//	mockgen -source=internal/example/example.go -destination=internal/example/mock/example.go -package mock_example
//

// Package mock_example is a generated GoMock package.
package mock_example

import (
	context "context"
	reflect "reflect"

	example "github.com/somatech1/mocks/internal/example"
	gomock "go.uber.org/mock/gomock"
)

// MockExampleMock is a mock of ExampleMock interface.
type MockExampleMock struct {
	ctrl     *gomock.Controller
	recorder *MockExampleMockMockRecorder
}

// MockExampleMockMockRecorder is the mock recorder for MockExampleMock.
type MockExampleMockMockRecorder struct {
	mock *MockExampleMock
}

// NewMockExampleMock creates a new mock instance.
func NewMockExampleMock(ctrl *gomock.Controller) *MockExampleMock {
	mock := &MockExampleMock{ctrl: ctrl}
	mock.recorder = &MockExampleMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExampleMock) EXPECT() *MockExampleMockMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockExampleMock) Any(ctx context.Context, in *example.Example, options ...string) (*example.Example, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(*example.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Any indicates an expected call of Any.
func (mr *MockExampleMockMockRecorder) Any(ctx, in any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockExampleMock)(nil).Any), varargs...)
}

// GetByInt mocks base method.
func (m *MockExampleMock) GetByInt(ctx context.Context, i int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByInt", ctx, i)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByInt indicates an expected call of GetByInt.
func (mr *MockExampleMockMockRecorder) GetByInt(ctx, i any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByInt", reflect.TypeOf((*MockExampleMock)(nil).GetByInt), ctx, i)
}

// GetByString mocks base method.
func (m *MockExampleMock) GetByString(ctx context.Context, str string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByString", ctx, str)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByString indicates an expected call of GetByString.
func (mr *MockExampleMockMockRecorder) GetByString(ctx, str any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByString", reflect.TypeOf((*MockExampleMock)(nil).GetByString), ctx, str)
}

// GetWithVariadic mocks base method.
func (m *MockExampleMock) GetWithVariadic(ctx context.Context, id string, options ...string) (int, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithVariadic", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithVariadic indicates an expected call of GetWithVariadic.
func (mr *MockExampleMockMockRecorder) GetWithVariadic(ctx, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithVariadic", reflect.TypeOf((*MockExampleMock)(nil).GetWithVariadic), varargs...)
}

// SingleError mocks base method.
func (m *MockExampleMock) SingleError(ctx context.Context, id string, options ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SingleError", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SingleError indicates an expected call of SingleError.
func (mr *MockExampleMockMockRecorder) SingleError(ctx, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingleError", reflect.TypeOf((*MockExampleMock)(nil).SingleError), varargs...)
}

// WithDoAndReturn mocks base method.
func (m *MockExampleMock) WithDoAndReturn(ctx context.Context, in *example.Example) (*example.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDoAndReturn", ctx, in)
	ret0, _ := ret[0].(*example.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithDoAndReturn indicates an expected call of WithDoAndReturn.
func (mr *MockExampleMockMockRecorder) WithDoAndReturn(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDoAndReturn", reflect.TypeOf((*MockExampleMock)(nil).WithDoAndReturn), ctx, in)
}

// WithStruct mocks base method.
func (m *MockExampleMock) WithStruct(ctx context.Context, in *example.Example) (*example.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStruct", ctx, in)
	ret0, _ := ret[0].(*example.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithStruct indicates an expected call of WithStruct.
func (mr *MockExampleMockMockRecorder) WithStruct(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStruct", reflect.TypeOf((*MockExampleMock)(nil).WithStruct), ctx, in)
}
