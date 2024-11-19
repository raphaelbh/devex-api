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

	var command = usecase.TriggerPipelineCommand{
		PipelineID: c.Param("id"),
		Input:      request.Input,
	}
	var executionID, _ = usecase.TriggerPipeline(command)
	c.IndentedJSON(http.StatusOK, buildResponse(TriggerPipelineResponse{
		ExecutionID: executionID,
	}))
}

type TriggerPipelineRequest struct {
	Input map[string]string `json:"input"`
}

type TriggerPipelineResponse struct {
	ExecutionID string `json:"execution_id"`
}
