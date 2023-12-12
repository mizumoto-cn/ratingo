package collector

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/topic"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/user"
)

type CollectRequest struct {
	Name    string  `json:"name"`
	Mail    string  `json:"mail"`
	Rating  float64 `json:"rating"`
	Comment string  `json:"comment"`
	Topic   string  `json:"topic"`
}

type CollectResponse struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
	TopicID int  `json:"topic_id"`
}

func Collect(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CollectRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the user exists, if not create a new user
		u, err := client.User.Query().Where(user.EmailEQ(req.Mail)).Only(c)
		if err != nil {
			u, err = client.User.Create().SetName(req.Name).SetEmail(req.Mail).Save(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// Check if the topic exists, if not create a new topic
		t, err := client.Topic.Query().Where(topic.TopicNameEQ(req.Topic)).Only(c)
		if err != nil {
			t, err = client.Topic.Create().SetTopicName(req.Topic).Save(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// Create a new rating in the database
		rating, err := client.Rating.
			Create().
			SetUserID(u.ID).
			SetTopicID(t.ID).
			SetRating(req.Rating).
			SetComment(req.Comment).
			Save(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, CollectResponse{Success: true, ID: rating.ID, TopicID: t.ID})

	}
}
