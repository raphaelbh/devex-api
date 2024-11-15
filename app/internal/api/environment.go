package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/usecase"
)

func getAllEnvironmentsHandler(c *gin.Context) {
	var data, _ = usecase.GetAllEnvironments()
	c.IndentedJSON(http.StatusOK, buildResponse(data))
}

func getEnvironmentHandler(c *gin.Context) {
	id := c.Param("id")
	data, _ := usecase.GetEnvironment(id)
	c.IndentedJSON(http.StatusOK, buildResponse(data))
}
