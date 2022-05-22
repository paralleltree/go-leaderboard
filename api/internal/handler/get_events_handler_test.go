package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	"github.com/paralleltree/go-leaderboard/internal/model"
	mock_usecase "github.com/paralleltree/go-leaderboard/mock/usecase"
)

type event struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

func TestGetEventsHandler(t *testing.T) {
	cases := []struct {
		Name       string
		QueryPage  string
		QueryCount string
		WantPage   int64
		WantCount  int64
	}{
		{
			Name:       "page not specified",
			QueryCount: "10",
			WantPage:   1,
			WantCount:  10,
		},
		{
			Name:      "count not specified",
			QueryPage: "10",
			WantPage:  10,
			WantCount: 20,
		},
		{
			Name:       "page and count specified",
			QueryPage:  "10",
			QueryCount: "30",
			WantPage:   10,
			WantCount:  30,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			getEventsUsecase := mock_usecase.NewMockGetEventsUsecase(mockCtrl)

			wantEventId := "1"
			wantEvent := model.Event{
				Name:    "test event",
				StartAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
				EndAt:   time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
			}
			wantEventRecord := model.Record[model.Event]{
				Id:   wantEventId,
				Item: wantEvent,
			}
			wantResponse := []event{
				{
					Id:      wantEventId,
					Name:    wantEvent.Name,
					StartAt: "2022-05-01T00:00:00Z",
					EndAt:   "2022-05-01T12:00:00Z",
				},
			}
			wantEventRecords := []model.Record[model.Event]{wantEventRecord}
			ctx := context.Background()
			wantStatus := http.StatusOK

			getEventsUsecase.EXPECT().
				GetEvents(ctx, tt.WantPage, tt.WantCount).
				Return([]model.Record[model.Event]{wantEventRecord}, nil)

			handler := handler.BuildGetEventsHandler(getEventsUsecase)

			u, err := url.Parse("/events")
			if err != nil {
				t.Fatalf("failed to parse url: %v", err)
			}

			params := u.Query()
			if tt.QueryPage != "" {
				params.Add("page", tt.QueryPage)
			}
			if tt.QueryCount != "" {
				params.Add("count", tt.QueryCount)
			}
			u.RawQuery = params.Encode()

			router, recorder := buildMockedServer()
			router.GET("/events", handler)
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
			res := []event{}
			if err := json.NewDecoder(recorder.Body).Decode(&res); err != nil {
				t.Fatalf("unexpected error while decoding response: %v", err)
			}
			if len(wantEventRecords) != len(res) {
				t.Fatalf("unexpected result count: expected %v, got %v", len(wantEventRecords), len(res))
			}
			for i, wantEvent := range wantResponse {
				gotEvent := res[i]
				if wantEvent != gotEvent {
					t.Fatalf("unexpected event: expected %v, but got %v", wantEvent, gotEvent)
				}
			}
		})
	}
}
