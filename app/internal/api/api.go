package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/raphaelbh/devex-api/internal/config"
)

var (
	conf = config.Default()
)

func Run() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(corsConfig()))

	// public rotes
	public := router.Group("/v1")
	public.GET("/health", healthHandler)
	public.GET("/github-events", githubEventsHandler)

	// private routes
	private := router.Group("/v1")
	private.Use(authenticationMiddleware())
	environments := private.Group("/environments")
	{
		environments.GET("", getAllEnvironmentsHandler)
		environments.GET(":id", getEnvironmentHandler)
	}
	pipelines := private.Group("/pipelines")
	{
		pipelines.POST(":id/trigger", triggerPipelineHandler)
	}

	router.Run(":" + conf.Server.Port)
}

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}
