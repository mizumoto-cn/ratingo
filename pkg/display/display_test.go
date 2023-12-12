package display_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/enttest"
	"github.com/mizumoto-cn/ratingo/pkg/display"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestDisplay(t *testing.T) {
	// Create a mock client
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Set up Gin with the Display handler
	r := gin.Default()
	r.POST("/display", display.Display(client))

	// Create a request to send to the handler
	reqBody := display.DisplayRequest{Topic: 1}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/display", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	resp := httptest.NewRecorder()

	// Call the handler
	r.ServeHTTP(resp, req)

	// Check the response
	println(resp.Body.String())
	// There is no such topic
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	var respBody display.DisplayResponse
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	// Add topic
	_, err := client.Topic.Create().
		SetTopicName("test").
		SetID(1).
		Save(req.Context())
	assert.NoError(t, err)
	// Add analysis
	ana, err := client.Analysis.Create().
		SetAvgRating(3.0).
		SetID(1).
		SetTopicID(1).
		Save(req.Context())
	assert.NoError(t, err)
	fmt.Printf("%+v\n", ana)
	// Record the response
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/display", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	// Call the handler
	r.ServeHTTP(resp, req)

	// Check the response
	println(resp.Body.String())
	// There is such topic
	assert.Equal(t, http.StatusOK, resp.Code)

}
