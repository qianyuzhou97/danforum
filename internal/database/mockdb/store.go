// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qianyuzhou97/danforum/internal/database (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/qianyuzhou97/danforum/internal/database"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockStore) Authenticate(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockStoreMockRecorder) Authenticate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockStore)(nil).Authenticate), arg0, arg1, arg2)
}

// CreateCommunity mocks base method.
func (m *MockStore) CreateCommunity(arg0 context.Context, arg1 database.NewCommunity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommunity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCommunity indicates an expected call of CreateCommunity.
func (mr *MockStoreMockRecorder) CreateCommunity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommunity", reflect.TypeOf((*MockStore)(nil).CreateCommunity), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 database.NewPost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 database.NewUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeletePostByID mocks base method.
func (m *MockStore) DeletePostByID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostByID indicates an expected call of DeletePostByID.
func (mr *MockStoreMockRecorder) DeletePostByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostByID", reflect.TypeOf((*MockStore)(nil).DeletePostByID), arg0, arg1)
}

// GetCommunityByID mocks base method.
func (m *MockStore) GetCommunityByID(arg0 context.Context, arg1 string) (*database.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommunityByID", arg0, arg1)
	ret0, _ := ret[0].(*database.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommunityByID indicates an expected call of GetCommunityByID.
func (mr *MockStoreMockRecorder) GetCommunityByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommunityByID", reflect.TypeOf((*MockStore)(nil).GetCommunityByID), arg0, arg1)
}

// GetPostByID mocks base method.
func (m *MockStore) GetPostByID(arg0 context.Context, arg1 int64) (*database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", arg0, arg1)
	ret0, _ := ret[0].(*database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockStoreMockRecorder) GetPostByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockStore)(nil).GetPostByID), arg0, arg1)
}

// ListAllCommunity mocks base method.
func (m *MockStore) ListAllCommunity(arg0 context.Context) ([]database.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllCommunity", arg0)
	ret0, _ := ret[0].([]database.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllCommunity indicates an expected call of ListAllCommunity.
func (mr *MockStoreMockRecorder) ListAllCommunity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllCommunity", reflect.TypeOf((*MockStore)(nil).ListAllCommunity), arg0)
}

// ListAllPosts mocks base method.
func (m *MockStore) ListAllPosts(arg0 context.Context) ([]database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllPosts", arg0)
	ret0, _ := ret[0].([]database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllPosts indicates an expected call of ListAllPosts.
func (mr *MockStoreMockRecorder) ListAllPosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllPosts", reflect.TypeOf((*MockStore)(nil).ListAllPosts), arg0)
}

// UpdatePostByID mocks base method.
func (m *MockStore) UpdatePostByID(arg0 context.Context, arg1 database.UpdatePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePostByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePostByID indicates an expected call of UpdatePostByID.
func (mr *MockStoreMockRecorder) UpdatePostByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePostByID", reflect.TypeOf((*MockStore)(nil).UpdatePostByID), arg0, arg1)
}
