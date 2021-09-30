package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/qianyuzhou97/danforum/internal/database"
	"github.com/qianyuzhou97/danforum/internal/database/mockdb"
	"github.com/stretchr/testify/require"
)

func TestListAllCommunity(t *testing.T) {

	communities := []database.Community{{
		ID:           1234,
		Name:         "Go Lover",
		Introduction: "This is the place where Gopher could discuss and share opinions",
		Create_time:  time.Now(),
		Update_time:  time.Now(),
	}, {
		ID:           123,
		Name:         "Cooking",
		Introduction: "Learn how to cook from here!",
		Create_time:  time.Now(),
		Update_time:  time.Now(),
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
