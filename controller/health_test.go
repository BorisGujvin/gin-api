package controller

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestIsHealthy_success(t *testing.T) {
    //Create a httptest recorder
    responseRecorder := httptest.NewRecorder()
    //Create a gin text context for the above recorder to get context and gin engine
    ctx, engine := gin.CreateTestContext(responseRecorder)
    //Register the healthcheck endpoint to the gin engine
    engine.GET("/healthcheck", IsHealthy)
    //Create a test request for the above registered endpoint
    ctx.Request = httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
    //Below line shows how to set headers if necessary
    ctx.Request.Header.Set("x-agentname", "agent name")

    //Test the endpoint
    engine.ServeHTTP(responseRecorder, ctx.Request)

    //Assert if response http status is as expected
    assert.Equal(t, http.StatusOK, responseRecorder.Code)
}