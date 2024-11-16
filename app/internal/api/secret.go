package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func createSecretHandler(c *gin.Context) {
	var request CreateSecretRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	usecase.CreateSecret(request.toCommand())
	c.Writer.WriteHeader(http.StatusCreated)
}

type CreateSecretRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r CreateSecretRequest) toCommand() usecase.CreateSecretCommand {
	return usecase.CreateSecretCommand{
		Key:   r.Key,
		Value: r.Value,
	}
}
