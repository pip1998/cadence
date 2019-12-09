// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: metadata.go

// Package cluster is a generated GoMock package.
package cluster

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	config "github.com/uber/cadence/common/service/config"
)

// MockMetadata is a mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataMockRecorder
}

// MockMetadataMockRecorder is the mock recorder for MockMetadata
type MockMetadataMockRecorder struct {
	mock *MockMetadata
}

// NewMockMetadata creates a new mock instance
func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &MockMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return m.recorder
}

// IsGlobalDomainEnabled mocks base method
func (m *MockMetadata) IsGlobalDomainEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsGlobalDomainEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsGlobalDomainEnabled indicates an expected call of IsGlobalDomainEnabled
func (mr *MockMetadataMockRecorder) IsGlobalDomainEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGlobalDomainEnabled", reflect.TypeOf((*MockMetadata)(nil).IsGlobalDomainEnabled))
}

// IsMasterCluster mocks base method
func (m *MockMetadata) IsMasterCluster() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMasterCluster")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMasterCluster indicates an expected call of IsMasterCluster
func (mr *MockMetadataMockRecorder) IsMasterCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMasterCluster", reflect.TypeOf((*MockMetadata)(nil).IsMasterCluster))
}

// GetNextFailoverVersion mocks base method
func (m *MockMetadata) GetNextFailoverVersion(arg0 string, arg1 int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextFailoverVersion", arg0, arg1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNextFailoverVersion indicates an expected call of GetNextFailoverVersion
func (mr *MockMetadataMockRecorder) GetNextFailoverVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextFailoverVersion", reflect.TypeOf((*MockMetadata)(nil).GetNextFailoverVersion), arg0, arg1)
}

// IsVersionFromSameCluster mocks base method
func (m *MockMetadata) IsVersionFromSameCluster(version1, version2 int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVersionFromSameCluster", version1, version2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVersionFromSameCluster indicates an expected call of IsVersionFromSameCluster
func (mr *MockMetadataMockRecorder) IsVersionFromSameCluster(version1, version2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVersionFromSameCluster", reflect.TypeOf((*MockMetadata)(nil).IsVersionFromSameCluster), version1, version2)
}

// GetMasterClusterName mocks base method
func (m *MockMetadata) GetMasterClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMasterClusterName indicates an expected call of GetMasterClusterName
func (mr *MockMetadataMockRecorder) GetMasterClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterClusterName", reflect.TypeOf((*MockMetadata)(nil).GetMasterClusterName))
}

// GetCurrentClusterName mocks base method
func (m *MockMetadata) GetCurrentClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCurrentClusterName indicates an expected call of GetCurrentClusterName
func (mr *MockMetadataMockRecorder) GetCurrentClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentClusterName", reflect.TypeOf((*MockMetadata)(nil).GetCurrentClusterName))
}

// GetAllClusterInfo mocks base method
func (m *MockMetadata) GetAllClusterInfo() map[string]config.ClusterInformation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllClusterInfo")
	ret0, _ := ret[0].(map[string]config.ClusterInformation)
	return ret0
}

// GetAllClusterInfo indicates an expected call of GetAllClusterInfo
func (mr *MockMetadataMockRecorder) GetAllClusterInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllClusterInfo", reflect.TypeOf((*MockMetadata)(nil).GetAllClusterInfo))
}

// ClusterNameForFailoverVersion mocks base method
func (m *MockMetadata) ClusterNameForFailoverVersion(failoverVersion int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNameForFailoverVersion", failoverVersion)
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterNameForFailoverVersion indicates an expected call of ClusterNameForFailoverVersion
func (mr *MockMetadataMockRecorder) ClusterNameForFailoverVersion(failoverVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNameForFailoverVersion", reflect.TypeOf((*MockMetadata)(nil).ClusterNameForFailoverVersion), failoverVersion)
}

// GetReplicationConsumerConfig mocks base method
func (m *MockMetadata) GetReplicationConsumerConfig() *config.ReplicationConsumerConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationConsumerConfig")
	ret0, _ := ret[0].(*config.ReplicationConsumerConfig)
	return ret0
}

// GetReplicationConsumerConfig indicates an expected call of GetReplicationConsumerConfig
func (mr *MockMetadataMockRecorder) GetReplicationConsumerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationConsumerConfig", reflect.TypeOf((*MockMetadata)(nil).GetReplicationConsumerConfig))
}