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
