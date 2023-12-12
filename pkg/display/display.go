package display

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/analysis"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/rating"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/topic"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/user"
)

type DisplayRequest struct {
	Topic int `json:"topic"`
}

type DisplayResponse struct {
	Topic   string  `json:"topic"`
	Avg     float64 `json:"avg"`
	Details []struct {
		Name    string  `json:"name"`
		Rating  float64 `json:"rating"`
		Comment string  `json:"comment"`
	} `json:"details"`
}

func Display(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DisplayRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tp, err := client.Topic.Query().Where(topic.IDEQ(req.Topic)).Only(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// read the average rating for the given topic
		analyze, err := client.Analysis.Query().
			Where(analysis.TopicIDEQ(tp.ID)).
			Only(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		avg := analyze.AvgRating

		// read the details for the given topic
		ratings, err := client.Rating.Query().
			Where(rating.TopicIDEQ(tp.ID)).
			All(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var details []struct {
			Name    string  `json:"name"`
			Rating  float64 `json:"rating"`
			Comment string  `json:"comment"`
		}

		for _, r := range ratings {
			u, err := client.User.Query().Where(user.IDEQ(r.UserID)).Only(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			details = append(details, struct {
				Name    string  `json:"name"`
				Rating  float64 `json:"rating"`
				Comment string  `json:"comment"`
			}{
				Name:    u.Name[0:1] + "****" + u.Name[len(u.Name)-1:],
				Rating:  r.Rating,
				Comment: r.Comment,
			})
		}

		c.JSON(http.StatusOK, DisplayResponse{
			Topic:   tp.TopicName,
			Avg:     avg,
			Details: details,
		})
	}
}
