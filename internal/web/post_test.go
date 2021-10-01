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

func TestListAllPosts(t *testing.T) {

	posts := []database.Post{{
		ID:      1234,
		Title:   "Join us on gopher slack",
		Content: "Just as title suggests",
		Author:  1,
	}, {
		ID:      123,
		Title:   "Join us on c++ slack",
		Content: "Just as title suggests",
		Author:  2,
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
					ListAllPosts(gomock.Any()).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPosts(t, recorder.Body, posts)
			},
		},
		{
			name: "Error",
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					ListAllPosts(gomock.Any()).
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

			request, err := http.NewRequest(http.MethodGet, "/posts", nil)

			require.NoError(t, err)

			srv.mux.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}

}

func TestGetPostByID(t *testing.T) {

	posts := []database.Post{{
		ID:      1234,
		Title:   "Join us on gopher slack",
		Content: "Just as title suggests",
		Author:  1,
	}, {
		ID:      123,
		Title:   "Join us on c++ slack",
		Content: "Just as title suggests",
		Author:  2,
	}}

	tests := []struct {
		name          string
		postID        int64
		mock          func(m *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			postID: 1234,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					GetPostByID(gomock.Any(), int64(1234)).
					Times(1).
					Return(&posts[0], nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPost(t, recorder.Body, posts[0])
			},
		},
		{
			name:   "Error",
			postID: 12,
			mock: func(m *mockdb.MockStore) {
				m.EXPECT().
					GetPostByID(gomock.Any(), int64(12)).
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

			url := fmt.Sprintf("/posts/%d", tt.postID)

			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			srv.mux.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}

}

func TestCreatePost(t *testing.T) {

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

// func TestUpdatePostByID(t *testing.T) {

// 	input := map[string]interface{}{
// 		"name":         "Go Lover",
// 		"introduction": "This is the place where Gopher could discuss and share opinions",
// 	}

// 	nc := database.NewCommunity{
// 		Name:         "Go Lover",
// 		Introduction: "This is the place where Gopher could discuss and share opinions",
// 	}

// 	tests := []struct {
// 		name          string
// 		input         map[string]interface{}
// 		nc            database.NewCommunity
// 		mock          func(m *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:  "OK",
// 			input: input,
// 			nc:    nc,
// 			mock: func(m *mockdb.MockStore) {
// 				m.EXPECT().
// 					CreateCommunity(gomock.Any(), nc).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name:  "Error",
// 			input: input,
// 			nc:    nc,
// 			mock: func(m *mockdb.MockStore) {
// 				m.EXPECT().
// 					CreateCommunity(gomock.Any(), nc).
// 					Times(1).
// 					Return(errors.New("field validation error"))
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			defer ctrl.Finish()
// 			m := mockdb.NewMockStore(ctrl)

// 			tt.mock(m)

// 			srv := NewServer().SetDB(m).SetRoutes(true)

// 			recorder := httptest.NewRecorder()

// 			data, err := json.Marshal(tt.input)
// 			require.NoError(t, err)

// 			request, err := http.NewRequest(http.MethodPost, "/community", bytes.NewReader(data))

// 			require.NoError(t, err)

// 			srv.mux.ServeHTTP(recorder, request)
// 			tt.checkResponse(t, recorder)
// 		})
// 	}

// }

// func TestDeletePostByID(t *testing.T) {

// 	input := map[string]interface{}{
// 		"name":         "Go Lover",
// 		"introduction": "This is the place where Gopher could discuss and share opinions",
// 	}

// 	nc := database.NewCommunity{
// 		Name:         "Go Lover",
// 		Introduction: "This is the place where Gopher could discuss and share opinions",
// 	}

// 	tests := []struct {
// 		name          string
// 		input         map[string]interface{}
// 		nc            database.NewCommunity
// 		mock          func(m *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:  "OK",
// 			input: input,
// 			nc:    nc,
// 			mock: func(m *mockdb.MockStore) {
// 				m.EXPECT().
// 					CreateCommunity(gomock.Any(), nc).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name:  "Error",
// 			input: input,
// 			nc:    nc,
// 			mock: func(m *mockdb.MockStore) {
// 				m.EXPECT().
// 					CreateCommunity(gomock.Any(), nc).
// 					Times(1).
// 					Return(errors.New("field validation error"))
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			defer ctrl.Finish()
// 			m := mockdb.NewMockStore(ctrl)

// 			tt.mock(m)

// 			srv := NewServer().SetDB(m).SetRoutes(true)

// 			recorder := httptest.NewRecorder()

// 			data, err := json.Marshal(tt.input)
// 			require.NoError(t, err)

// 			request, err := http.NewRequest(http.MethodPost, "/community", bytes.NewReader(data))

// 			require.NoError(t, err)

// 			srv.mux.ServeHTTP(recorder, request)
// 			tt.checkResponse(t, recorder)
// 		})
// 	}

// }

func requireBodyMatchPosts(t *testing.T, body *bytes.Buffer, posts []database.Post) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var got []database.Post
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, posts, got)
}

func requireBodyMatchPost(t *testing.T, body *bytes.Buffer, post database.Post) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var got database.Post
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, post, got)
}
