package router

import (
	"github.com/Mario-Jimenez/gocompiler/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// newRouter configuration and endpoints
func newRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	monkeyHandler := new(handler.Monkey)

	monkeyRoutes := router.Group("/monkeycompiler")
	{
		monkeyRoutes.POST("/compile", monkeyHandler.Compile)
	}

	return router
}
