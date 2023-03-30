package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yosikez/item-api/model"
	"github.com/yosikez/item-api/service"
	"github.com/yosikez/item-api/structs"
	cusMess "github.com/yosikez/custom-error-message"
	"gorm.io/gorm"
)

type ItemController interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type itemController struct {
	service service.ItemService
}

func NewItemController(service service.ItemService) *itemController {
	return &itemController{
		service: service,
	}
}

func (controller *itemController) FindAll(c *gin.Context) {
	items, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var itemsResponse []structs.ItemResponse

	for _, it := range items {
		itemResponse := convertToResponseStruct(it)
		itemsResponse = append(itemsResponse, itemResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : itemsResponse,
	})
	
}

func (controller *itemController) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format id",
		})
		return
	}

	item, err := controller.service.FindByID(uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	default:
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data" : convertToResponseStruct(item),
	})

}

func (controller *itemController) Create(c *gin.Context) {
	var newItem structs.ItemRequest

	if err := c.ShouldBindJSON(&newItem); err != nil {
		errFields := cusMess.GetErrMess(err, newItem, nil)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors" : errFields,
		})
		return
	}

	item, err := controller.service.Create(newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponseStruct(item),
	})
}

func (controller *itemController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format id",
		})
		return
	}

	var updateItem structs.ItemRequest

	if err := c.ShouldBindJSON(&updateItem); err != nil {
		errFields := cusMess.GetErrMess(err, updateItem, nil)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors" : errFields,
		})
		return
	}

	item, err := controller.service.Update(uint(id), updateItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponseStruct(item),
	})
}

func (controller *itemController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format id",
		})
		return
	}

	err = controller.service.Delete(uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	default:
	}

	c.Status(http.StatusNoContent)
}


func convertToResponseStruct(i model.Item) structs.ItemResponse {
	return structs.ItemResponse{
		ID: i.Id,
		Name: i.Name,
		Code: i.Code,
		Quantity: i.Quantity,
	}
}