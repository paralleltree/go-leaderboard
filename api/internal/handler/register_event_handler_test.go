package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	"github.com/paralleltree/go-leaderboard/internal/model"
	mock_usecase "github.com/paralleltree/go-leaderboard/mock/usecase"
)

func TestRegisterEventHandler_WhenParameterIsInvalid_ReturnsBadRequest(t *testing.T) {
	cases := []struct {
		Name   string
		Params map[string]string
	}{
		{
			Name: "name missing",
			Params: map[string]string{
				"start_at": "2022-05-01T12:00:00+09:00",
				"end_at":   "2022-05-01T15:00:00+09:00",
			},
		},
		{
			Name: "start_at missing",
			Params: map[string]string{
				"name":   "test event",
				"end_at": "2022-05-01T15:00:00+09:00",
			},
		},
		{
			Name: "end_at missing",
			Params: map[string]string{
				"name":     "test event",
				"start_at": "2022-05-01T12:00:00+09:00",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			wantStatus := http.StatusBadRequest

			usecase := mock_usecase.NewMockRegisterEventUsecase(mockCtrl)

			handler := handler.BuildRegisterEventHandler(usecase)
			router, recorder := buildMockedServer()
			router.POST("/", handler)
			payload, err := json.Marshal(tt.Params)
			if err != nil {
				t.Fatalf("failed to marshal params: %v", err)
			}
			req, err := http.NewRequest("POST", "/", bytes.NewBuffer(payload))
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

func TestRegisterEventHandler_WhenRequestIsValid_ReturnsCreated(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	wantName := "test"
	params := map[string]string{
		"name":     wantName,
		"start_at": "2022-05-01T00:00:00Z",
		"end_at":   "2022-05-01T15:00:00+09:00",
	}
	wantEvent := model.Event{
		Name:    wantName,
		StartAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2022, 5, 1, 6, 0, 0, 0, time.UTC),
	}
	wantStatus := http.StatusCreated
	wantEventId := "1"
	wantContent := fmt.Sprintf(`{"id":"%s"}`, wantEventId)

	usecase := mock_usecase.NewMockRegisterEventUsecase(mockCtrl)
	usecase.EXPECT().
		RegisterEvent(gomock.Any(), wantEvent).
		Return(wantEventId, nil)

	handler := handler.BuildRegisterEventHandler(usecase)
	router, recorder := buildMockedServer()
	router.POST("/", handler)
	payload, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("failed to marshal params: %v", err)
	}
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("unexpected error while requesting: %v", err)
	}
	router.ServeHTTP(recorder, req)

	gotStatus := recorder.Result().StatusCode
	if wantStatus != gotStatus {
		t.Fatalf("unexpected status code: expected %v, but got %v", wantStatus, gotStatus)
	}

	defer recorder.Result().Body.Close()
	gotContent, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatalf("unexpected error while reading body: %v", err)
	}
	if wantContent != string(gotContent) {
		t.Fatalf("unexpected content: expected %v, but got %v", wantContent, gotContent)
	}
}
