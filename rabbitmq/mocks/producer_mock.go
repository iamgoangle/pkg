// Code generated by MockGen. DO NOT EDIT.
// Source: ./producer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
	reflect "reflect"
)

// MockProducer is a mock of Producer interface
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// UseWithConfig mocks base method
func (m *MockProducer) UseWithConfig(config ...rabbitmq.ProducerConfigHandler) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range config {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UseWithConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseWithConfig indicates an expected call of UseWithConfig
func (mr *MockProducerMockRecorder) UseWithConfig(config ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseWithConfig", reflect.TypeOf((*MockProducer)(nil).UseWithConfig), config...)
}

// Publish mocks base method
func (m *MockProducer) Publish(body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockProducerMockRecorder) Publish(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockProducer)(nil).Publish), body)
}
