package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(config))

	router.GET("/health", healthHandler)
	router.POST("/github/webhook", webhookHandler)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	router.Run(":" + port)

}

func healthHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":    "UP",
		"timestamp": time.Now(),
	})
}

func webhookHandler(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	println("Webhook called")
	println(payload)

	c.Writer.WriteHeader(http.StatusNoContent)
}
