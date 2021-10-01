package web

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/database"
	"github.com/qianyuzhou97/danforum/internal/database/mockdb"
	"github.com/stretchr/testify/require"
)

func TestListAllCommunity(t *testing.T) {

	communities := []database.Community{{
		ID:           1234,
		Name:         "Go Lover",
		Introduction: "This is the place where Gopher could discuss and share opinions",
	}, {
		ID:           123,
		Name:         "Cooking",
		Introduction: "Learn how to cook from here!",
	}}

	tests := []struct {
		name          string
		mock          func(m *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					ListAllCommunity(gomock.Any()).
					Times(1).
					Return(communities, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCommunities(t, recorder.Body, communities)
			},
		},
		{
			name: "Error",
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					ListAllCommunity(gomock.Any()).
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()
			m := mockdb.NewMockStore(ctrl)

			tt.mock(m)

			srv := NewServer().SetDB(m).SetRoutes(true)

			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/community", nil)

			require.NoError(t, err)

			srv.mux.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}

}

func TestGetCommunityByID(t *testing.T) {

	communities := []database.Community{{
		ID:           1234,
		Name:         "Go Lover",
		Introduction: "This is the place where Gopher could discuss and share opinions",
	}, {
		ID:           123,
		Name:         "Cooking",
		Introduction: "Learn how to cook from here!",
	}}

	tests := []struct {
		name          string
		communityID   int64
		mock          func(m *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			communityID: 1234,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					GetCommunityByID(gomock.Any(), int64(1234)).
					Times(1).
					Return(&communities[0], nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCommunity(t, recorder.Body, communities[0])
			},
		},
		{
			name:        "Error",
			communityID: 12,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					GetCommunityByID(gomock.Any(), int64(12)).
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()
			m := mockdb.NewMockStore(ctrl)

			tt.mock(m)

			srv := NewServer().SetDB(m).SetRoutes(true)

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/community/%d", tt.communityID)

			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			srv.mux.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}

}

func TestCreateCommunity(t *testing.T) {

	input := map[string]interface{}{
		"name":         "Go Lover",
		"introduction": "This is the place where Gopher could discuss and share opinions",
	}

	nc := database.NewCommunity{
		Name:         "Go Lover",
		Introduction: "This is the place where Gopher could discuss and share opinions",
	}

	tests := []struct {
		name          string
		input         map[string]interface{}
		nc            database.NewCommunity
		mock          func(m *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			input: input,
			nc:    nc,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					CreateCommunity(gomock.Any(), nc).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "Error",
			input: input,
			nc:    nc,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					CreateCommunity(gomock.Any(), nc).
					Times(1).
					Return(errors.New("field validation error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()
			m := mockdb.NewMockStore(ctrl)

			tt.mock(m)

			srv := NewServer().SetDB(m).SetRoutes(true)

			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tt.input)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/community", bytes.NewReader(data))

			require.NoError(t, err)

			srv.mux.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchCommunities(t *testing.T, body *bytes.Buffer, communities []database.Community) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var got []database.Community
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, communities, got)
}

func requireBodyMatchCommunity(t *testing.T, body *bytes.Buffer, community database.Community) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var got database.Community
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, community, got)
}
