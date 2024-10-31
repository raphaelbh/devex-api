package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":    "UP",
		"timestamp": time.Now(),
	})
}
