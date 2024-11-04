package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func getEnvironmentHandler(c *gin.Context) {
	id := c.Param("id")
	data, _ := usecase.GetEnvironment(id)
	c.IndentedJSON(http.StatusOK, buildResponse(data))
}
