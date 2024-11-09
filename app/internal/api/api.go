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

	public := router.Group("/v1")
	public.GET("/health", healthHandler)

	// environments
	environments := public.Group("/environments")
	{
		environments.GET("", getAllEnvironmentsHandler)
		environments.GET(":id", getEnvironmentHandler)
	}

	// events
	events := public.Group("/events")
	{
		events.POST("", gitEventsHandler)
	}

	// pipelines
	pipelines := public.Group("/pipelines")
	{
		pipelines.POST(":id/execute", executePipelineHandler)
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
