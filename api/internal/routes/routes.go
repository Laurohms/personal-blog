package routes

import (
	"github.com/Laurohms/blog-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Engine) error {
	api := c.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", handlers.GetAllUsers)
			users.GET("/name/:name")
			users.GET("/:id", handlers.GetUserById)
			users.POST("/", handlers.CreateUser)
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
