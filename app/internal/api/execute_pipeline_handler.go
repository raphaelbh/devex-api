package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func executePipelineHandler(c *gin.Context) {
	usecase.ExecutePipeline()
	c.Writer.WriteHeader(http.StatusNoContent)
}
