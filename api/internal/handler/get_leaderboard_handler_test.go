package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	"github.com/paralleltree/go-leaderboard/internal/model"
	mock_usecase "github.com/paralleltree/go-leaderboard/mock/usecase"
)

func TestGetLeaderboardHandler(t *testing.T) {
	cases := []struct {
		Name           string
		QueryStartRank string
		QueryEndRank   string
		WantStartRank  int64
		WantEndRank    int64
		WantRank       []model.UserRank
	}{
		{
			Name:           "start and end specified",
			QueryStartRank: "10",
			QueryEndRank:   "11",
			WantStartRank:  10,
			WantEndRank:    11,
			WantRank: []model.UserRank{
				{Rank: 10, UserId: "1", Score: 14},
				{Rank: 11, UserId: "2", Score: 42},
			},
		},
		{
			Name:           "start specified",
			QueryStartRank: "10",
			WantStartRank:  10,
			WantEndRank:    100,
			WantRank: []model.UserRank{
				{Rank: 10, UserId: "1", Score: 14},
			},
		},
		{
			Name:          "end specified",
			QueryEndRank:  "10",
			WantStartRank: 1,
			WantEndRank:   10,
			WantRank: []model.UserRank{
				{Rank: 10, UserId: "1", Score: 14},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			usecase := mock_usecase.NewMockGetLeaderboardUsecase(mockCtrl)

			wantEventId := "1"
			wantStatus := http.StatusOK

			usecase.EXPECT().
				GetLeaderboard(gomock.Any(), wantEventId, tt.WantStartRank, tt.WantEndRank).
				Return(tt.WantRank, true, nil)

			handler := handler.BuildGetLeaderboardHandler(usecase)

			u, err := url.Parse(fmt.Sprintf("/events/%s/leaderboard", wantEventId))
			if err != nil {
				t.Fatalf("failed to parse url: %v", err)
			}

			params := u.Query()
			if tt.QueryStartRank != "" {
				params.Add("start", tt.QueryStartRank)
			}
			if tt.QueryEndRank != "" {
				params.Add("end", tt.QueryEndRank)
			}
			u.RawQuery = params.Encode()

			router, recorder := buildMockedServer()
			router.GET("/events/:id/leaderboard", handler)
			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				t.Fatalf("unexpected error while requesting: %v", err)
			}
			router.ServeHTTP(recorder, req)

			gotStatus := recorder.Result().StatusCode
			if wantStatus != gotStatus {
				t.Fatalf("unexpected status code: expected %v, but got %v", wantStatus, gotStatus)
			}

			defer recorder.Result().Body.Close()
			res := []struct {
				Rank   int64  `json:"rank"`
				UserId string `json:"user_id"`
				Score  int64  `json:"score"`
			}{}
			if err := json.NewDecoder(recorder.Body).Decode(&res); err != nil {
				t.Fatalf("unexpected error while decoding response: %v", err)
			}
			if len(tt.WantRank) != len(res) {
				t.Fatalf("result count mismatch: expected count %v, but got %v", len(tt.WantRank), len(res))
			}
			for i, wantUserRank := range tt.WantRank {
				gotUserRank := model.UserRank(res[i])
				if wantUserRank != gotUserRank {
					t.Fatalf("result rank mismatch: expected %v, but got %v", wantUserRank, gotUserRank)
				}
			}
		})
	}
}
