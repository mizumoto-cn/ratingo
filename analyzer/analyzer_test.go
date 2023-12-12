package analyzer_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/analyzer"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/enttest"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestAnalyze(t *testing.T) {
	// Create a mock client
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Set up Gin with the Analyze handler
	r := gin.Default()
	r.POST("/analyze", analyzer.Analyze(client))

	// Create a request to send to the handler
	reqBody := analyzer.AnalyzeRequest{Topic: 1}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	resp := httptest.NewRecorder()

	// Call the handler
	r.ServeHTTP(resp, req)

	// Check the response
	fmt.Printf("%+v\n", resp.Body.String())
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	// Insert topic
	_, err := client.Topic.Create().
		SetTopicName("test").
		SetID(1).
		Save(req.Context())
	assert.NoError(t, err)

	// Record the response
	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	// Call the handler
	r.ServeHTTP(resp, req)
	// Check the response
	fmt.Printf("%+v\n", resp.Body.String())
	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody analyzer.AnalyzeResponse
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	// Add assertions based on your requirements
	assert.True(t, respBody.OK)
}
