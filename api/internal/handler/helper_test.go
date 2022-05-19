package handler_test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func buildMockedServer() (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	_, router := gin.CreateTestContext(recorder)
	return router, recorder
}
