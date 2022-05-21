package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	mock_usecase "github.com/paralleltree/go-leaderboard/mock/usecase"
)

func TestSetScoreHandler_WithoutParams_ReturnsBadRequest(t *testing.T) {
	cases := []struct {
		Name   string
		Params map[string]interface{}
	}{
		{
			Name: "user_id missing",
			Params: map[string]interface{}{
				"score": int64(42),
			},
		},
		{
			Name: "score missing",
			Params: map[string]interface{}{
				"user_id": "2",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			usecase := mock_usecase.NewMockSetScoreUsecase(mockCtrl)

			wantEventId := "1"
			wantStatus := http.StatusBadRequest

			handler := handler.BuildSetScoreHandler(usecase)
			router, recorder := buildMockedServer()
			router.PUT("/:id", handler)
			payload, err := json.Marshal(tt.Params)
			if err != nil {
				t.Fatalf("failed to marshal params: %v", err)
			}
			req, err := http.NewRequest("PUT", fmt.Sprintf("/%s", wantEventId), bytes.NewBuffer(payload))
			if err != nil {
				t.Fatalf("unexpected error while requesting: %v", err)
			}
			router.ServeHTTP(recorder, req)

			gotStatus := recorder.Result().StatusCode
			if wantStatus != gotStatus {
				t.Fatalf("unexpected status code: expected %v, but got %v", wantStatus, gotStatus)
			}
		})
	}
}

func TestSetScoreHandler_WithValidParams_ReturnsOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	usecase := mock_usecase.NewMockSetScoreUsecase(mockCtrl)

	wantEventId := "1"
	wantUserId := "2"
	wantScore := int64(42)
	params := map[string]interface{}{
		"user_id": wantUserId,
		"score":   wantScore,
	}
	wantStatus := http.StatusOK

	usecase.EXPECT().
		SetScore(gomock.Any(), wantEventId, wantUserId, wantScore).
		Return(nil)

	handler := handler.BuildSetScoreHandler(usecase)
	router, recorder := buildMockedServer()
	router.PUT("/:id", handler)
	payload, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("failed to marshal params: %v", err)
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("/%s", wantEventId), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("unexpected error while requesting: %v", err)
	}
	router.ServeHTTP(recorder, req)

	gotStatus := recorder.Result().StatusCode
	if wantStatus != gotStatus {
		t.Fatalf("unexpected status code: expected %v, but got %v", wantStatus, gotStatus)
	}
}
