// Code generated by MockGen. DO NOT EDIT.
// Source: ./consumer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
	amqp "github.com/streadway/amqp"
	reflect "reflect"
)

// MockConsumerHandler is a mock of ConsumerHandler interface
type MockConsumerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerHandlerMockRecorder
}

// MockConsumerHandlerMockRecorder is the mock recorder for MockConsumerHandler
type MockConsumerHandlerMockRecorder struct {
	mock *MockConsumerHandler
}

// NewMockConsumerHandler creates a new mock instance
func NewMockConsumerHandler(ctrl *gomock.Controller) *MockConsumerHandler {
	mock := &MockConsumerHandler{ctrl: ctrl}
	mock.recorder = &MockConsumerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumerHandler) EXPECT() *MockConsumerHandlerMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockConsumerHandler) Do(msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockConsumerHandlerMockRecorder) Do(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockConsumerHandler)(nil).Do), msg)
}

// OnSuccess mocks base method
func (m_2 *MockConsumerHandler) OnSuccess(m amqp.Delivery) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "OnSuccess", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnSuccess indicates an expected call of OnSuccess
func (mr *MockConsumerHandlerMockRecorder) OnSuccess(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSuccess", reflect.TypeOf((*MockConsumerHandler)(nil).OnSuccess), m)
}

// OnError mocks base method
func (m_2 *MockConsumerHandler) OnError(m amqp.Delivery, err error) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "OnError", m, err)
}

// OnError indicates an expected call of OnError
func (mr *MockConsumerHandlerMockRecorder) OnError(m, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockConsumerHandler)(nil).OnError), m, err)
}

// MockConsumer is a mock of Consumer interface
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// WithConfigs mocks base method
func (m *MockConsumer) WithConfigs(configs ...rabbitmq.ConsumerConfigHandler) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range configs {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "WithConfigs", varargs...)
}

// WithConfigs indicates an expected call of WithConfigs
func (mr *MockConsumerMockRecorder) WithConfigs(configs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithConfigs", reflect.TypeOf((*MockConsumer)(nil).WithConfigs), configs...)
}

// WithDeadLetterQueue mocks base method
func (m *MockConsumer) WithDeadLetterQueue(configs ...rabbitmq.ConsumerConfigDLQHandler) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range configs {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "WithDeadLetterQueue", varargs...)
}

// WithDeadLetterQueue indicates an expected call of WithDeadLetterQueue
func (mr *MockConsumerMockRecorder) WithDeadLetterQueue(configs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDeadLetterQueue", reflect.TypeOf((*MockConsumer)(nil).WithDeadLetterQueue), configs...)
}

// Use mocks base method
func (m *MockConsumer) Use(handler rabbitmq.ConsumerHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Use", handler)
}

// Use indicates an expected call of Use
func (mr *MockConsumerMockRecorder) Use(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockConsumer)(nil).Use), handler)
}

// Consume mocks base method
func (m *MockConsumer) Consume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume")
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume
func (mr *MockConsumerMockRecorder) Consume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockConsumer)(nil).Consume))
}

// ConsumeWithRetry mocks base method
func (m *MockConsumer) ConsumeWithRetry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsumeWithRetry")
}

// ConsumeWithRetry indicates an expected call of ConsumeWithRetry
func (mr *MockConsumerMockRecorder) ConsumeWithRetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeWithRetry", reflect.TypeOf((*MockConsumer)(nil).ConsumeWithRetry))
}
