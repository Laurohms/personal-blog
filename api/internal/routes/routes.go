package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(c *gin.Engine) error {
	api := c.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"ping": "kong",
				})
			})
			users.GET("/name/:name")
			users.GET("/:id")
			users.POST("/")
			users.PUT("/:id")
			users.DELETE("/:id")
		}

		blogPosts := api.Group("/posts")
		{
			blogPosts.GET("/")
			blogPosts.GET("/:id")
			blogPosts.GET("/title/:title")
			blogPosts.GET("/content/:content")
			blogPosts.POST("/")
			blogPosts.PUT("/:id")
			blogPosts.DELETE("/:id")
		}
	}
	return nil
}
