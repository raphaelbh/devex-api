package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/domain/application"
)

type GitPushRequest struct {
	Ref    string `json:"ref"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Repository struct {
		Name    string `json:"name"`
		HtmlUrl string `json:"html_url"`
	} `json:"repository"`
}

func gitEventsHandler(c *gin.Context) {
	var request GitPushRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event"})
		return
	}

	var command = application.TriggerPipelinesCommand{
		Branch: request.Ref,
		Pusher: struct {
			Name  string
			Email string
		}{
			Name:  request.Pusher.Name,
			Email: request.Pusher.Email,
		},
		Repository: struct {
			Name string
			Url  string
		}{
			Name: request.Repository.Name,
			Url:  request.Repository.HtmlUrl,
		},
	}

	if err := application.TriggerPipelines(command); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
