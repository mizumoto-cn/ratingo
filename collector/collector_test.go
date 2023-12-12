package collector_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/collector"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/enttest"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/rating"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/topic"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/user"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestCollect(t *testing.T) {
	// Create a mock client
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Set up Gin with the Collect handler
	r := gin.Default()
	r.POST("/collect", collector.Collect(client))

	// Create a request to send to the handler
	reqBody := collector.CollectRequest{
		Name:    "Test User",
		Mail:    "test@example.com",
		Rating:  4.5,
		Comment: "Great topic!",
		Topic:   "Test Topic",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/collect", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	fmt.Printf("%+v\n", req)
	// Record the response
	resp := httptest.NewRecorder()

	// Call the handler
	r.ServeHTTP(resp, req)

	// Check the response
	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody collector.CollectResponse
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)
	fmt.Printf("%+v\n", respBody)
	// check respBody.Success is true
	assert.True(t, respBody.Success)

	// check database for user
	u, err := client.User.Query().Where(user.EmailEQ(reqBody.Mail)).Only(req.Context())
	assert.NoError(t, err)
	assert.Equal(t, reqBody.Name, u.Name)
	fmt.Printf("%+v\n", u)

	// check database for topic
	to, err := client.Topic.Query().Where(topic.TopicNameEQ(reqBody.Topic)).Only(req.Context())
	assert.NoError(t, err)
	assert.Equal(t, reqBody.Topic, to.TopicName)
	fmt.Printf("%+v\n", to)

	// check database for rating
	ra, err := client.Rating.Query().Where(rating.IDEQ(respBody.ID)).Only(req.Context())
	assert.NoError(t, err)
	assert.Equal(t, reqBody.Rating, ra.Rating)
	assert.Equal(t, reqBody.Comment, ra.Comment)
	fmt.Printf("%+v\n", ra)
}
