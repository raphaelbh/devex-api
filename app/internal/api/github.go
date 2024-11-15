package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func githubEventsHandler(c *gin.Context) {
	var request GitEventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := usecase.GitEvent(request.toCommand()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

type GitEventRequest struct {
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

func (r GitEventRequest) toCommand() usecase.GitEventCommand {
	return usecase.GitEventCommand{
		Branch: r.Ref,
		Pusher: struct {
			Name  string
			Email string
		}{
			Name:  r.Pusher.Name,
			Email: r.Pusher.Email,
		},
		Repository: struct {
			Name string
			Url  string
		}{
			Name: r.Repository.Name,
			Url:  r.Repository.HtmlUrl,
		},
	}
}
