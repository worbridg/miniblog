// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package mock is a generated GoMock package.
package mock

import (
	template "html/template"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/worbridg/miniblog/domain"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// BlogID mocks base method.
func (m *MockContext) BlogID() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlogID")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlogID indicates an expected call of BlogID.
func (mr *MockContextMockRecorder) BlogID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogID", reflect.TypeOf((*MockContext)(nil).BlogID))
}

// BlogPage mocks base method.
func (m *MockContext) BlogPage(blog domain.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogPage", blog)
}

// BlogPage indicates an expected call of BlogPage.
func (mr *MockContextMockRecorder) BlogPage(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogPage", reflect.TypeOf((*MockContext)(nil).BlogPage), blog)
}

// CreateBlog mocks base method.
func (m *MockContext) CreateBlog() (domain.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlog")
	ret0, _ := ret[0].(domain.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlog indicates an expected call of CreateBlog.
func (mr *MockContextMockRecorder) CreateBlog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlog", reflect.TypeOf((*MockContext)(nil).CreateBlog))
}

// EditPage mocks base method.
func (m *MockContext) EditPage(id int, blog domain.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EditPage", id, blog)
}

// EditPage indicates an expected call of EditPage.
func (mr *MockContextMockRecorder) EditPage(id, blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPage", reflect.TypeOf((*MockContext)(nil).EditPage), id, blog)
}

// ErrorPage mocks base method.
func (m *MockContext) ErrorPage(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ErrorPage", err)
}

// ErrorPage indicates an expected call of ErrorPage.
func (mr *MockContextMockRecorder) ErrorPage(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorPage", reflect.TypeOf((*MockContext)(nil).ErrorPage), err)
}

// File mocks base method.
func (m *MockContext) File(filename string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "File", filename)
}

// File indicates an expected call of File.
func (mr *MockContextMockRecorder) File(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockContext)(nil).File), filename)
}

// HTML mocks base method.
func (m *MockContext) HTML(file []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HTML", file)
}

// HTML indicates an expected call of HTML.
func (mr *MockContextMockRecorder) HTML(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTML", reflect.TypeOf((*MockContext)(nil).HTML), file)
}

// Header mocks base method.
func (m *MockContext) Header(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Header", key, value)
}

// Header indicates an expected call of Header.
func (mr *MockContextMockRecorder) Header(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockContext)(nil).Header), key, value)
}

// IndexPage mocks base method.
func (m *MockContext) IndexPage(blogs map[int]domain.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IndexPage", blogs)
}

// IndexPage indicates an expected call of IndexPage.
func (mr *MockContextMockRecorder) IndexPage(blogs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexPage", reflect.TypeOf((*MockContext)(nil).IndexPage), blogs)
}

// Method mocks base method.
func (m *MockContext) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockContextMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockContext)(nil).Method))
}

// Redirect mocks base method.
func (m *MockContext) Redirect(url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Redirect", url)
}

// Redirect indicates an expected call of Redirect.
func (mr *MockContextMockRecorder) Redirect(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockContext)(nil).Redirect), url)
}

// Request mocks base method.
func (m *MockContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockContext)(nil).Request))
}

// ResponseWriter mocks base method.
func (m *MockContext) ResponseWriter() http.ResponseWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseWriter")
	ret0, _ := ret[0].(http.ResponseWriter)
	return ret0
}

// ResponseWriter indicates an expected call of ResponseWriter.
func (mr *MockContextMockRecorder) ResponseWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseWriter", reflect.TypeOf((*MockContext)(nil).ResponseWriter))
}

// Status mocks base method.
func (m *MockContext) Status(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Status", code)
}

// Status indicates an expected call of Status.
func (mr *MockContextMockRecorder) Status(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockContext)(nil).Status), code)
}

// Template mocks base method.
func (m *MockContext) Template(t *template.Template, data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Template", t, data)
}

// Template indicates an expected call of Template.
func (mr *MockContextMockRecorder) Template(t, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockContext)(nil).Template), t, data)
}
