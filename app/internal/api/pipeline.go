package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func triggerPipelineHandler(c *gin.Context) {
	var request TriggerPipelineRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	usecase.TriggerPipeline(request.toCommand())
	c.Writer.WriteHeader(http.StatusNoContent)
}

type TriggerPipelineRequest struct {
	Commands []string `json:"commands"`
}

func (r TriggerPipelineRequest) toCommand() usecase.TriggerPipelineCommand {
	return usecase.TriggerPipelineCommand{
		Commands: r.Commands,
	}
}
