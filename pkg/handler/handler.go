package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/singin")
		auth.POST("/signup")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/")
			lists.POST("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.GET("/")
				items.POST("/")
				items.GET("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}

	return router
}
