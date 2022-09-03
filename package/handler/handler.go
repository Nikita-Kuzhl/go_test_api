package handler

import (
	"github.com/Nikita-Kuzhl/go-rest-api/package/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h*Handler) InitRouter() *gin.Engine {
	router:=gin.New();

	auth:= router.Group("/auth")
	{
		auth.POST("/signup",h.singUp)
		auth.POST("/signin",h.singIn)
	}
	api:=router.Group("/api")
	{
		lists:=api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/",h.getAllLists)
			lists.GET("/:id",h.getListById)
			lists.PUT("/:id",h.updateList)
			lists.DELETE("/:id",h.deleteList)
			items:=lists.Group(":id/items")
			{
				items.POST("/",h.createItem)
				items.GET("/",h.getAllItems)
				items.GET("/:item_id",h.getItemById)
				items.PUT("/:item_id",h.updateItem)
				items.DELETE("/:item_id",h.deleteItem)
			}

		}
	}
	return router
}