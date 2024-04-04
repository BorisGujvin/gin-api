package content

import (
    "fmt"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
    "os"
    "io"
    "net/textproto"
    "mime/multipart"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetContents_success(t *testing.T) {
    responseRecorder := httptest.NewRecorder()
    ctx, engine := gin.CreateTestContext(responseRecorder)
    engine.GET("/contents", GetContents)
    ctx.Request = httptest.NewRequest(http.MethodGet, "/contents?name=deepak", nil)
    engine.ServeHTTP(responseRecorder, ctx.Request)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)

    expectedBody := map[string]string{
        "name": "deepak",
    }

    var actualBody map[string]string
    err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualBody)
    if err != nil {
        t.FailNow()
    }

    assert.Equal(t, expectedBody, actualBody)
}

func TestGetContents_failure(t *testing.T) {
    responseRecorder := httptest.NewRecorder()
    ctx, engine := gin.CreateTestContext(responseRecorder)
    engine.GET("/contents", GetContents)
    ctx.Request = httptest.NewRequest(http.MethodGet, "/contents", nil)
    engine.ServeHTTP(responseRecorder, ctx.Request)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestPostContents_success(t *testing.T) {
    responseRecorder := httptest.NewRecorder()
    ctx, engine := gin.CreateTestContext(responseRecorder)
    engine.POST("/contents", PostContents)
    requestBody := `{"id":"some-random-id","name":"content name"}`
    ctx.Request = httptest.NewRequest(http.MethodPost, "/contents", bytes.NewBuffer([]byte(requestBody)))
    engine.ServeHTTP(responseRecorder, ctx.Request)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.Equal(t, requestBody, responseRecorder.Body.String())
}

func TestPostContents_failure(t *testing.T) {
    responseRecorder := httptest.NewRecorder()
    ctx, engine := gin.CreateTestContext(responseRecorder)
    engine.POST("/contents", PostContents)
    ctx.Request = httptest.NewRequest(http.MethodPost, "/contents", nil)
    engine.ServeHTTP(responseRecorder, ctx.Request)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConsumeFile_success(t *testing.T) {
    //Create multipart request
    body := new(bytes.Buffer)
    multipartWriter := multipart.NewWriter(body)
    //Create multipart header
    fileHeader := make(textproto.MIMEHeader)
    fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", "sample_success.csv"))
    fileHeader.Set("Content-Type", "text/plain")
    writer, err := multipartWriter.CreatePart(fileHeader)
    assert.Nil(t, err)
    //Copy file to file multipart writer
    file, err := os.Open("txt/sample_success.txt")
    assert.Nil(t, err)
    io.Copy(writer, file)
    // close the writer before making the request
    multipartWriter.Close()

    w := httptest.NewRecorder()
    ctx, r := gin.CreateTestContext(w)

    r.POST("/content/import", ConsumeFile)
    ctx.Request = httptest.NewRequest(http.MethodPost, "/content/import", body)
    ctx.Request.Header.Add("Content-Type", multipartWriter.FormDataContentType())
    r.ServeHTTP(w, ctx.Request)

    assert.Equal(t, http.StatusOK, w.Code)

}