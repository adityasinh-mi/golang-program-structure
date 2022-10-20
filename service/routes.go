package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"golang-program-structure/common/middleware"
)

func AddRoutes(r *gin.Engine, config *Config) {
	requestAuth := middleware.NewRequestAuth(config.Logger)

	userGroup := r.Group("/v1")
	userGroup.Use(requestAuth.Middleware)
	userGroup.GET("/user", func(c *gin.Context) {
		c.JSON(200, "Sucess1")
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, "Sucess2")
	})
}

func SetupDefaultEndpoints(r *gin.Engine, config *Config) {
	r.GET("/", func(c *gin.Context) {
		htmlString := "<html><body>Welcome to Golang-Program-Structure!</body></html>"
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte(htmlString))
		if err != nil {
			return
		}
	})
	r.GET("/ping", func(c *gin.Context) {
		var msg string
		if config.Env == "production" {
			msg = fmt.Sprintf("Pong! I am %s. Version is %s.", config.ServiceName, config.Version)
		} else {
			msg = "pong"
		}
		c.JSON(200, gin.H{"message": msg})
	})
	if config.Env == "development" || config.Env == "staging" {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
