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
// Source: github.com/specterops/bloodhound/src/queries (interfaces: Graph)
//
// Generated by this command:
//
//	mockgen -copyright_file=../../../../LICENSE.header -destination=./mocks/graph.go -package=mocks . Graph
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	graph "github.com/specterops/bloodhound/dawgs/graph"
	model "github.com/specterops/bloodhound/src/model"
	queries "github.com/specterops/bloodhound/src/queries"
	agi "github.com/specterops/bloodhound/src/services/agi"
	gomock "go.uber.org/mock/gomock"
)

// MockGraph is a mock of Graph interface.
type MockGraph struct {
	ctrl     *gomock.Controller
	recorder *MockGraphMockRecorder
	isgomock struct{}
}

// MockGraphMockRecorder is the mock recorder for MockGraph.
type MockGraphMockRecorder struct {
	mock *MockGraph
}

// NewMockGraph creates a new mock instance.
func NewMockGraph(ctrl *gomock.Controller) *MockGraph {
	mock := &MockGraph{ctrl: ctrl}
	mock.recorder = &MockGraphMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraph) EXPECT() *MockGraphMockRecorder {
	return m.recorder
}

// BatchNodeUpdate mocks base method.
func (m *MockGraph) BatchNodeUpdate(ctx context.Context, nodeUpdate graph.NodeUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchNodeUpdate", ctx, nodeUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchNodeUpdate indicates an expected call of BatchNodeUpdate.
func (mr *MockGraphMockRecorder) BatchNodeUpdate(ctx, nodeUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchNodeUpdate", reflect.TypeOf((*MockGraph)(nil).BatchNodeUpdate), ctx, nodeUpdate)
}

// CountNodesByKind mocks base method.
func (m *MockGraph) CountNodesByKind(ctx context.Context, kinds ...graph.Kind) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range kinds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountNodesByKind", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountNodesByKind indicates an expected call of CountNodesByKind.
func (mr *MockGraphMockRecorder) CountNodesByKind(ctx any, kinds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, kinds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountNodesByKind", reflect.TypeOf((*MockGraph)(nil).CountNodesByKind), varargs...)
}

// FetchNodesByObjectIDs mocks base method.
func (m *MockGraph) FetchNodesByObjectIDs(ctx context.Context, objectIDs ...string) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range objectIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchNodesByObjectIDs", varargs...)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNodesByObjectIDs indicates an expected call of FetchNodesByObjectIDs.
func (mr *MockGraphMockRecorder) FetchNodesByObjectIDs(ctx any, objectIDs ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, objectIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNodesByObjectIDs", reflect.TypeOf((*MockGraph)(nil).FetchNodesByObjectIDs), varargs...)
}

// FetchNodesByObjectIDsAndKinds mocks base method.
func (m *MockGraph) FetchNodesByObjectIDsAndKinds(ctx context.Context, kinds graph.Kinds, objectIDs ...string) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, kinds}
	for _, a := range objectIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchNodesByObjectIDsAndKinds", varargs...)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNodesByObjectIDsAndKinds indicates an expected call of FetchNodesByObjectIDsAndKinds.
func (mr *MockGraphMockRecorder) FetchNodesByObjectIDsAndKinds(ctx, kinds any, objectIDs ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, kinds}, objectIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNodesByObjectIDsAndKinds", reflect.TypeOf((*MockGraph)(nil).FetchNodesByObjectIDsAndKinds), varargs...)
}

// GetADEntityQueryResult mocks base method.
func (m *MockGraph) GetADEntityQueryResult(ctx context.Context, params queries.EntityQueryParameters, cacheEnabled bool) (any, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetADEntityQueryResult", ctx, params, cacheEnabled)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetADEntityQueryResult indicates an expected call of GetADEntityQueryResult.
func (mr *MockGraphMockRecorder) GetADEntityQueryResult(ctx, params, cacheEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetADEntityQueryResult", reflect.TypeOf((*MockGraph)(nil).GetADEntityQueryResult), ctx, params, cacheEnabled)
}

// GetAllShortestPaths mocks base method.
func (m *MockGraph) GetAllShortestPaths(ctx context.Context, startNodeID, endNodeID string, filter graph.Criteria) (graph.PathSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllShortestPaths", ctx, startNodeID, endNodeID, filter)
	ret0, _ := ret[0].(graph.PathSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllShortestPaths indicates an expected call of GetAllShortestPaths.
func (mr *MockGraphMockRecorder) GetAllShortestPaths(ctx, startNodeID, endNodeID, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllShortestPaths", reflect.TypeOf((*MockGraph)(nil).GetAllShortestPaths), ctx, startNodeID, endNodeID, filter)
}

// GetAssetGroupComboNode mocks base method.
func (m *MockGraph) GetAssetGroupComboNode(ctx context.Context, owningObjectID, assetGroupTag string) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetGroupComboNode", ctx, owningObjectID, assetGroupTag)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetGroupComboNode indicates an expected call of GetAssetGroupComboNode.
func (mr *MockGraphMockRecorder) GetAssetGroupComboNode(ctx, owningObjectID, assetGroupTag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetGroupComboNode", reflect.TypeOf((*MockGraph)(nil).GetAssetGroupComboNode), ctx, owningObjectID, assetGroupTag)
}

// GetAssetGroupNodes mocks base method.
func (m *MockGraph) GetAssetGroupNodes(ctx context.Context, assetGroupTag string, isSystemGroup bool) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetGroupNodes", ctx, assetGroupTag, isSystemGroup)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetGroupNodes indicates an expected call of GetAssetGroupNodes.
func (mr *MockGraphMockRecorder) GetAssetGroupNodes(ctx, assetGroupTag, isSystemGroup any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetGroupNodes", reflect.TypeOf((*MockGraph)(nil).GetAssetGroupNodes), ctx, assetGroupTag, isSystemGroup)
}

// GetEntityByObjectId mocks base method.
func (m *MockGraph) GetEntityByObjectId(ctx context.Context, objectID string, kinds ...graph.Kind) (*graph.Node, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, objectID}
	for _, a := range kinds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEntityByObjectId", varargs...)
	ret0, _ := ret[0].(*graph.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntityByObjectId indicates an expected call of GetEntityByObjectId.
func (mr *MockGraphMockRecorder) GetEntityByObjectId(ctx, objectID any, kinds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, objectID}, kinds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntityByObjectId", reflect.TypeOf((*MockGraph)(nil).GetEntityByObjectId), varargs...)
}

// GetEntityCountResults mocks base method.
func (m *MockGraph) GetEntityCountResults(ctx context.Context, node *graph.Node, delegates map[string]any) map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntityCountResults", ctx, node, delegates)
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// GetEntityCountResults indicates an expected call of GetEntityCountResults.
func (mr *MockGraphMockRecorder) GetEntityCountResults(ctx, node, delegates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntityCountResults", reflect.TypeOf((*MockGraph)(nil).GetEntityCountResults), ctx, node, delegates)
}

// GetFilteredAndSortedNodes mocks base method.
func (m *MockGraph) GetFilteredAndSortedNodes(orderCriteria model.OrderCriteria, filterCriteria graph.Criteria) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilteredAndSortedNodes", orderCriteria, filterCriteria)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilteredAndSortedNodes indicates an expected call of GetFilteredAndSortedNodes.
func (mr *MockGraphMockRecorder) GetFilteredAndSortedNodes(orderCriteria, filterCriteria any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilteredAndSortedNodes", reflect.TypeOf((*MockGraph)(nil).GetFilteredAndSortedNodes), orderCriteria, filterCriteria)
}

// GetNodesByKind mocks base method.
func (m *MockGraph) GetNodesByKind(ctx context.Context, kinds ...graph.Kind) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range kinds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNodesByKind", varargs...)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodesByKind indicates an expected call of GetNodesByKind.
func (mr *MockGraphMockRecorder) GetNodesByKind(ctx any, kinds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, kinds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodesByKind", reflect.TypeOf((*MockGraph)(nil).GetNodesByKind), varargs...)
}

// GetPrimaryNodeKindCounts mocks base method.
func (m *MockGraph) GetPrimaryNodeKindCounts(ctx context.Context, kinds ...graph.Kind) (map[string]int, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range kinds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPrimaryNodeKindCounts", varargs...)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrimaryNodeKindCounts indicates an expected call of GetPrimaryNodeKindCounts.
func (mr *MockGraphMockRecorder) GetPrimaryNodeKindCounts(ctx any, kinds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, kinds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryNodeKindCounts", reflect.TypeOf((*MockGraph)(nil).GetPrimaryNodeKindCounts), varargs...)
}

// PrepareCypherQuery mocks base method.
func (m *MockGraph) PrepareCypherQuery(rawCypher string, queryComplexityLimit int64) (queries.PreparedQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCypherQuery", rawCypher, queryComplexityLimit)
	ret0, _ := ret[0].(queries.PreparedQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCypherQuery indicates an expected call of PrepareCypherQuery.
func (mr *MockGraphMockRecorder) PrepareCypherQuery(rawCypher, queryComplexityLimit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCypherQuery", reflect.TypeOf((*MockGraph)(nil).PrepareCypherQuery), rawCypher, queryComplexityLimit)
}

// RawCypherQuery mocks base method.
func (m *MockGraph) RawCypherQuery(ctx context.Context, pQuery queries.PreparedQuery, includeProperties bool) (model.UnifiedGraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawCypherQuery", ctx, pQuery, includeProperties)
	ret0, _ := ret[0].(model.UnifiedGraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawCypherQuery indicates an expected call of RawCypherQuery.
func (mr *MockGraphMockRecorder) RawCypherQuery(ctx, pQuery, includeProperties any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawCypherQuery", reflect.TypeOf((*MockGraph)(nil).RawCypherQuery), ctx, pQuery, includeProperties)
}

// SearchByNameOrObjectID mocks base method.
func (m *MockGraph) SearchByNameOrObjectID(ctx context.Context, searchValue, searchType string) (graph.NodeSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByNameOrObjectID", ctx, searchValue, searchType)
	ret0, _ := ret[0].(graph.NodeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByNameOrObjectID indicates an expected call of SearchByNameOrObjectID.
func (mr *MockGraphMockRecorder) SearchByNameOrObjectID(ctx, searchValue, searchType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByNameOrObjectID", reflect.TypeOf((*MockGraph)(nil).SearchByNameOrObjectID), ctx, searchValue, searchType)
}

// SearchNodesByName mocks base method.
func (m *MockGraph) SearchNodesByName(ctx context.Context, nodeKinds graph.Kinds, nameQuery string, skip, limit int) ([]model.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodesByName", ctx, nodeKinds, nameQuery, skip, limit)
	ret0, _ := ret[0].([]model.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodesByName indicates an expected call of SearchNodesByName.
func (mr *MockGraphMockRecorder) SearchNodesByName(ctx, nodeKinds, nameQuery, skip, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodesByName", reflect.TypeOf((*MockGraph)(nil).SearchNodesByName), ctx, nodeKinds, nameQuery, skip, limit)
}

// UpdateSelectorTags mocks base method.
func (m *MockGraph) UpdateSelectorTags(ctx context.Context, db agi.AgiData, selectors model.UpdatedAssetGroupSelectors) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSelectorTags", ctx, db, selectors)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSelectorTags indicates an expected call of UpdateSelectorTags.
func (mr *MockGraphMockRecorder) UpdateSelectorTags(ctx, db, selectors any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSelectorTags", reflect.TypeOf((*MockGraph)(nil).UpdateSelectorTags), ctx, db, selectors)
}

// ValidateOUs mocks base method.
func (m *MockGraph) ValidateOUs(ctx context.Context, ous []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateOUs", ctx, ous)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateOUs indicates an expected call of ValidateOUs.
func (mr *MockGraphMockRecorder) ValidateOUs(ctx, ous any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateOUs", reflect.TypeOf((*MockGraph)(nil).ValidateOUs), ctx, ous)
}
