package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yosikez/item-api/repository"
	"github.com/yosikez/item-api/service"
	"github.com/yosikez/item-api/controller"
)

var (
	itemRepo = repository.NewItemRepository()
	itemService = service.NewItemService(itemRepo)
	itemController = controller.NewItemController(itemService)
)

func Api(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/items", itemController.FindAll)
		v1.GET("/items/:id", itemController.FindByID)
		v1.POST("/items", itemController.Create)
		v1.PUT("/items/:id", itemController.Update)
		v1.DELETE("/items/:id", itemController.Delete)
	}
}