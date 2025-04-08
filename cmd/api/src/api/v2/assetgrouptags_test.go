// Copyright 2025 Specter Ops, Inc.
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

package v2_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	uuid2 "github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/specterops/bloodhound/graphschema/ad"
	"github.com/specterops/bloodhound/headers"
	"github.com/specterops/bloodhound/mediatypes"
	"github.com/specterops/bloodhound/src/api"
	v2 "github.com/specterops/bloodhound/src/api/v2"
	"github.com/specterops/bloodhound/src/api/v2/apitest"
	"github.com/specterops/bloodhound/src/database"
	mocks_db "github.com/specterops/bloodhound/src/database/mocks"
	"github.com/specterops/bloodhound/src/database/types/null"
	"github.com/specterops/bloodhound/src/model"
	"github.com/specterops/bloodhound/src/queries"
	mocks_graph "github.com/specterops/bloodhound/src/queries/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestResources_CreateAssetGroupTagSelector(t *testing.T) {
	var (
		mockCtrl      = gomock.NewController(t)
		mockDB        = mocks_db.NewMockDatabase(mockCtrl)
		mockGraphDb   = mocks_graph.NewMockGraph(mockCtrl)
		resourcesInst = v2.Resources{
			DB:         mockDB,
			GraphQuery: mockGraphDb,
		}
		user    = setupUser()
		userCtx = setupUserCtx(user)
	)

	defer mockCtrl.Finish()

	apitest.
		NewHarness(t, resourcesInst.CreateAssetGroupTagSelector).
		Run([]apitest.Case{
			{
				Name: "BadRequest",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyString(input, `{"name":["BadRequest"]}`)
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
					apitest.BodyContains(output, api.ErrorResponsePayloadUnmarshalError)
				},
			},
			{
				Name: "MissingName",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
					apitest.BodyContains(output, "Name: property is required")
				},
			},
			{
				Name: "MissingSeeds",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
					apitest.BodyContains(output, "Seeds: property is required")
				},
			},
			{
				Name: "MissingUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "InvalidUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "non-numeric")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "NonExistentTagUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1234")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, database.ErrNotFound).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsResourceNotFound)
				},
			},
			{
				Name: "DatabaseError",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockGraphDb.EXPECT().
						PrepareCypherQuery(gomock.Any(), gomock.Any()).
						Return(queries.PreparedQuery{}, nil).Times(1)
					mockDB.EXPECT().
						CreateAssetGroupTagSelector(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, errors.New("failure")).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusInternalServerError)
					apitest.BodyContains(output, api.ErrorResponseDetailsInternalServerError)
				},
			},
			{
				Name: "InvalidCypher",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "cypher that's too complex"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockGraphDb.EXPECT().
						PrepareCypherQuery(gomock.Any(), gomock.Any()).
						Return(queries.PreparedQuery{}, queries.ErrCypherQueryTooComplex).Times(1)

				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
				},
			},
			{
				Name: "InvalidSeedType",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: 0, Value: ""},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
				},
			},
			{
				Name: "Success",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "MATCH (n:User) RETURN n LIMIT 1;"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().
						CreateAssetGroupTagSelector(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{Name: "TestSelector"}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)

					mockGraphDb.EXPECT().
						PrepareCypherQuery(gomock.Any(), gomock.Any()).
						Return(queries.PreparedQuery{}, nil).Times(1)

				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusCreated)
				},
			},
		})
}

func TestDatabase_GetAssetGroupTag(t *testing.T) {
	var (
		mockCtrl    = gomock.NewController(t)
		mockDB      = mocks_db.NewMockDatabase(mockCtrl)
		mockGraphDb = mocks_graph.NewMockGraph(mockCtrl)
		resources   = v2.Resources{
			DB:         mockDB,
			GraphQuery: mockGraphDb,
		}
	)

	defer mockCtrl.Finish()

	userId, err := uuid2.NewV4()
	require.Nil(t, err)

	endpoint := "/api/v2/asset-group-tags/%s"
	assetGroupTagId := "5"

	t.Run("successfully got asset group tag", func(t *testing.T) {
		mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), int(5)).Return(model.AssetGroupTag{
			ID:             5,
			Name:           "test tag 5",
			Description:    "some description",
			RequireCertify: null.BoolFrom(true),
		}, nil)

		req, err := http.NewRequestWithContext(createContextWithOwnerId(userId), "GET", fmt.Sprintf(endpoint, assetGroupTagId), nil)
		require.Nil(t, err)

		req.Header.Set(headers.ContentType.String(), mediatypes.ApplicationJson.String())
		req = mux.SetURLVars(req, map[string]string{api.URIPathVariableAssetGroupTagID: assetGroupTagId})

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(resources.GetAssetGroupTag)

		handler.ServeHTTP(response, req)
		require.Equal(t, http.StatusOK, response.Code)

		bodyBytes, err := io.ReadAll(response.Body)
		require.Nil(t, err)

		var temp struct {
			Data struct {
				Tag model.AssetGroupTag `json:"tag"`
			} `json:"data"`
		}
		err = json.Unmarshal(bodyBytes, &temp)
		require.Nil(t, err)

		parsedTime, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
		require.Nil(t, err)

		require.Equal(t, model.AssetGroupTag{
			ID:             5,
			Type:           0,
			KindId:         0,
			Name:           "test tag 5",
			Description:    "some description",
			CreatedAt:      parsedTime,
			CreatedBy:      "",
			UpdatedAt:      parsedTime,
			UpdatedBy:      "",
			DeletedAt:      null.Time{},
			DeletedBy:      null.String{},
			Position:       null.Int32{},
			RequireCertify: null.BoolFrom(true),
		}, temp.Data.Tag)
	})

	t.Run("asset group tag doesn't exist error", func(t *testing.T) {
		mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).Return(model.AssetGroupTag{}, database.ErrNotFound)

		req, err := http.NewRequestWithContext(createContextWithOwnerId(userId), "GET", fmt.Sprintf(endpoint, assetGroupTagId), nil)
		require.Nil(t, err)

		req.Header.Set(headers.ContentType.String(), mediatypes.ApplicationJson.String())
		req = mux.SetURLVars(req, map[string]string{api.URIPathVariableAssetGroupTagID: assetGroupTagId})

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(resources.GetAssetGroupTag)

		handler.ServeHTTP(response, req)
		require.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("id malformed error", func(t *testing.T) {
		req, err := http.NewRequestWithContext(createContextWithOwnerId(userId), "GET", fmt.Sprintf(endpoint, assetGroupTagId), nil)
		require.Nil(t, err)

		req.Header.Set(headers.ContentType.String(), mediatypes.ApplicationJson.String())

		router := mux.NewRouter()
		router.HandleFunc("/api/v2/asset-group-tags/{5}", resources.GetAssetGroupTag).Methods("GET")

		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		require.Equal(t, http.StatusNotFound, response.Code)
		require.Contains(t, response.Body.String(), api.ErrorResponseDetailsIDMalformed)
	})
}

func TestResources_UpdateAssetGroupTagSelector(t *testing.T) {
	var (
		mockCtrl      = gomock.NewController(t)
		mockDB        = mocks_db.NewMockDatabase(mockCtrl)
		mockGraphDb   = mocks_graph.NewMockGraph(mockCtrl)
		resourcesInst = v2.Resources{
			DB:         mockDB,
			GraphQuery: mockGraphDb,
		}
		user    = setupUser()
		userCtx = setupUserCtx(user)
	)

	defer mockCtrl.Finish()

	apitest.
		NewHarness(t, resourcesInst.UpdateAssetGroupTagSelector).
		Run([]apitest.Case{
			{
				Name: "BadRequest",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyString(input, `{"name":["BadRequest"]}`)
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTagSelectorBySelectorId(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusBadRequest)
					apitest.BodyContains(output, api.ErrorResponsePayloadUnmarshalError)
				},
			},
			{
				Name: "MissingTagUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "InvalidTagUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "non-numeric")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "MissingSelectorUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "InvalidSelectorUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "non-numeric")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "NonExistentTagUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1234")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, database.ErrNotFound).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsResourceNotFound)
				},
			},
			{
				Name: "NonExistentSelectorUrlId",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1234")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTagSelectorBySelectorId(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, database.ErrNotFound).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsResourceNotFound)
				},
			},
			{
				Name: "CannotUpdateNameOnDefaultSelector",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name: "TestSelector",
					})
				},
				Setup: func() {
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTagSelectorBySelectorId(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{IsDefault: true}, nil).Times(1)

					mockGraphDb.EXPECT().
						PrepareCypherQuery(gomock.Any(), gomock.Any()).
						Return(queries.PreparedQuery{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusForbidden)
					apitest.BodyContains(output, "default selectors only support modifying auto_certify and disabled_at")
				},
			},
			{
				Name: "DatabaseError",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "this should be a string of cypher"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().UpdateAssetGroupTagSelector(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, errors.New("failure")).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTagSelectorBySelectorId(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusInternalServerError)
					apitest.BodyContains(output, api.ErrorResponseDetailsInternalServerError)
				},
			},
			{
				Name: "Success",
				Input: func(input *apitest.Input) {
					apitest.SetContext(input, userCtx)
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagSelectorID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "MATCH (n:User) RETURN n LIMIT 1;"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().
						UpdateAssetGroupTagSelector(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{Name: "TestSelector"}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTagSelectorBySelectorId(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelector{}, nil).Times(1)

					mockGraphDb.EXPECT().
						PrepareCypherQuery(gomock.Any(), gomock.Any()).
						Return(queries.PreparedQuery{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusOK)
				},
			},
		})
}

func TestResources_GetAssetGroupTagSelectors(t *testing.T) {
	var (
		mockCtrl      = gomock.NewController(t)
		mockDB        = mocks_db.NewMockDatabase(mockCtrl)
		resourcesInst = v2.Resources{
			DB: mockDB,
		}
	)
	defer mockCtrl.Finish()

	apitest.
		NewHarness(t, resourcesInst.GetAssetGroupTagSelectors).
		Run([]apitest.Case{
			{
				Name: "Bad Request - Invalid Asset Group Tag ID",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "foo")
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "DB error - GetAssetGroupTag",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, database.ErrNotFound).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsResourceNotFound)
				},
			},
			{
				Name: "DB error - GetAssetGroupTagSelectorsByTagId",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTagSelectorsByTagId(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelectors{}, errors.New("some error")).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusInternalServerError)
				},
			},
			{
				Name: "Success",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
					apitest.BodyStruct(input, model.AssetGroupTagSelector{
						Name:        "TestSelector",
						Description: "Test selector description",
						Seeds: []model.SelectorSeed{
							{Type: model.SelectorTypeCypher, Value: "MATCH (n:User) RETURN n LIMIT 1;"},
						},
						IsDefault:   false,
						AutoCertify: null.BoolFrom(false),
					})
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTagSelectorsByTagId(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTagSelectors{{Name: "Test1", AssetGroupTagId: 1}}, nil).Times(1)
					mockDB.EXPECT().GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, nil).Times(1)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusOK)
					apitest.BodyContains(output, "Test1")
				},
			},
		})
}

func TestResources_GetAssetGroupTagSelectorsByTagId(t *testing.T) {
	var (
		mockCtrl    = gomock.NewController(t)
		mockDB      = mocks_db.NewMockDatabase(mockCtrl)
		mockGraphDb = mocks_graph.NewMockGraph(mockCtrl)
		resources   = v2.Resources{
			DB:         mockDB,
			GraphQuery: mockGraphDb,
		}
		assetGroupTag = model.AssetGroupTag{ID: 1, Name: "Tier Zero"}
	)
	defer mockCtrl.Finish()

	apitest.NewHarness(t, resources.GetAssetGroupTagMemberCountsByKind).
		Run([]apitest.Case{
			{
				Name: "InvalidAssetGroupTagID",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "invalid")
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusNotFound)
					apitest.BodyContains(output, api.ErrorResponseDetailsIDMalformed)
				},
			},
			{
				Name: "DatabaseGetAssetGroupTagError",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(model.AssetGroupTag{}, errors.New("GetAssetGroupTag fail"))
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusInternalServerError)
					apitest.BodyContains(output, api.ErrorResponseDetailsInternalServerError)
				},
			},
			{
				Name: "GraphDatabaseError",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(assetGroupTag, nil)
					mockGraphDb.EXPECT().
						GetPrimaryNodeKindCounts(gomock.Any(), gomock.Any()).
						Return(map[string]int{}, fmt.Errorf("GetAssetGroupTag Nodes fail"))
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusInternalServerError)
				},
			},
			{
				Name: "Success",
				Input: func(input *apitest.Input) {
					apitest.SetURLVar(input, api.URIPathVariableAssetGroupTagID, "1")
				},
				Setup: func() {
					mockDB.EXPECT().
						GetAssetGroupTag(gomock.Any(), gomock.Any()).
						Return(assetGroupTag, nil)
					mockGraphDb.EXPECT().
						GetPrimaryNodeKindCounts(gomock.Any(), gomock.Any()).
						Return(map[string]int{ad.Domain.String(): 2}, nil)
				},
				Test: func(output apitest.Output) {
					apitest.StatusCode(output, http.StatusOK)
					result := v2.GetAssetGroupTagMemberCountsResponse{}
					apitest.UnmarshalData(output, &result)
					require.Equal(t, 2, result.TotalCount)
					require.Equal(t, 2, result.Counts[ad.Domain.String()])
				},
			},
		})
}
