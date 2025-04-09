// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/specterops/bloodhound/src/services/agi (interfaces: AgiData)
//
// Generated by this command:
//
//	mockgen -copyright_file=../../../../../LICENSE.header -destination=./mocks/mock.go -package=mocks . AgiData
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/specterops/bloodhound/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAgiData is a mock of AgiData interface.
type MockAgiData struct {
	ctrl     *gomock.Controller
	recorder *MockAgiDataMockRecorder
	isgomock struct{}
}

// MockAgiDataMockRecorder is the mock recorder for MockAgiData.
type MockAgiDataMockRecorder struct {
	mock *MockAgiData
}

// NewMockAgiData creates a new mock instance.
func NewMockAgiData(ctrl *gomock.Controller) *MockAgiData {
	mock := &MockAgiData{ctrl: ctrl}
	mock.recorder = &MockAgiDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgiData) EXPECT() *MockAgiDataMockRecorder {
	return m.recorder
}

// CreateAssetGroupCollection mocks base method.
func (m *MockAgiData) CreateAssetGroupCollection(ctx context.Context, collection model.AssetGroupCollection, entries model.AssetGroupCollectionEntries) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAssetGroupCollection", ctx, collection, entries)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAssetGroupCollection indicates an expected call of CreateAssetGroupCollection.
func (mr *MockAgiDataMockRecorder) CreateAssetGroupCollection(ctx, collection, entries any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAssetGroupCollection", reflect.TypeOf((*MockAgiData)(nil).CreateAssetGroupCollection), ctx, collection, entries)
}

// GetAllAssetGroups mocks base method.
func (m *MockAgiData) GetAllAssetGroups(ctx context.Context, order string, filter model.SQLFilter) (model.AssetGroups, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAssetGroups", ctx, order, filter)
	ret0, _ := ret[0].(model.AssetGroups)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAssetGroups indicates an expected call of GetAllAssetGroups.
func (mr *MockAgiDataMockRecorder) GetAllAssetGroups(ctx, order, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAssetGroups", reflect.TypeOf((*MockAgiData)(nil).GetAllAssetGroups), ctx, order, filter)
}

// GetAssetGroup mocks base method.
func (m *MockAgiData) GetAssetGroup(ctx context.Context, id int32) (model.AssetGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetGroup", ctx, id)
	ret0, _ := ret[0].(model.AssetGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetGroup indicates an expected call of GetAssetGroup.
func (mr *MockAgiDataMockRecorder) GetAssetGroup(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetGroup", reflect.TypeOf((*MockAgiData)(nil).GetAssetGroup), ctx, id)
}
