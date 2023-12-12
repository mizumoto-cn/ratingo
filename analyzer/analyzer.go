package analyzer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/analysis"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/rating"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/topic"
)

type AnalyzeRequest struct {
	Topic int `json:"topic"`
}

type AnalyzeResponse struct {
	OK bool `json:"ok"`
}

func Analyze(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AnalyzeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tp, err := client.Topic.Query().Where(topic.IDEQ(req.Topic)).Only(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Calculate the average rating for the given topic
		// get the amount of ratings for the topic
		// if amount is 0, set avg to 0
		// else calculate the average rating
		ratings, err := client.Rating.Query().
			Where(rating.HasUnderTopicOfWith(topic.IDEQ(tp.ID))).
			All(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var avg float64

		if len(ratings) == 0 {
			avg = 0
		} else {
			var sum float64
			for _, r := range ratings {
				sum += r.Rating
			}
			avg = sum / float64(len(ratings))
		}

		// Save the average rating to the database
		// check if Analysis already exists
		a, err := client.Analysis.
			Query().
			Where(analysis.TopicIDEQ(tp.ID)). // filter by topic
			Only(c)

		if ent.IsNotFound(err) {
			// if not, create a new Analysis
			_, err = client.Analysis.
				Create().
				SetTopicID(req.Topic).
				SetAvgRating(avg).
				Save(c)
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			// else update the existing Analysis
			_, err = client.Analysis.
				UpdateOne(a).
				SetAvgRating(avg).
				Save(c)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, AnalyzeResponse{OK: true})
	}
}
