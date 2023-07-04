package router

import (
	"github.com/gin-gonic/gin"
	"khl-meobot/actions"
	"khl-meobot/router/middlewares"
)

func InitRouter(e *gin.Engine) {
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")
	index := e.Group("/")
	{
		index.Any("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "This is kook bot app."})
		})
		index.GET("/actuator/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})
	}
	kook := e.Group("/kook")
	kook.Use(middlewares.CheckVerifyToken())
	{
		kook.POST("", actions.KookEntry)
	}
	spo := e.Group("/spotify")
	{
		spo.GET("/req", actions.ReqSpotifyAuth)
		spo.GET("/callback", actions.CompleteSpotifyAuth)
		spo.GET("/queue", actions.GetSpotifyQueueAPI)
	}
}
