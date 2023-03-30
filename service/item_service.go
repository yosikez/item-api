package service

import (
	"github.com/yosikez/item-api/model"
	"github.com/yosikez/item-api/repository"
	"github.com/yosikez/item-api/structs"
)

type ItemService interface {
	FindAll() ([]model.Item, error)
	FindByID(ID uint) (model.Item, error)
	Create(item structs.ItemRequest) (model.Item, error)
	Update(ID uint, item structs.ItemRequest) (model.Item, error)
	Delete(ID uint) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repository repository.ItemRepository) *itemService {
	return &itemService{
		repo: repository,
	}
}

func (service *itemService) FindAll() ([]model.Item, error) {
	return service.repo.FindAll()
}

func (service *itemService) FindByID(ID uint) (model.Item, error) {
	return service.repo.FindByID(ID)
}

func (service *itemService) Create(itemRequest structs.ItemRequest) (model.Item, error) {
	item := model.Item{
		Name: itemRequest.Name,
		Code: itemRequest.Code,
		Quantity: itemRequest.Quantity,
	}

	return service.repo.Create(item)
}

func (service *itemService) Update(ID uint, itemRequest structs.ItemRequest) (model.Item, error) {
	item, err := service.repo.FindByID(ID)
	if err != nil {
		return item, err
	}

	item.Name = itemRequest.Name
	item.Code = itemRequest.Code
	item.Quantity = itemRequest.Quantity
	
	return service.repo.Update(item)
}

func (service *itemService) Delete(ID uint) error {
	return service.repo.Delete(ID)
}
