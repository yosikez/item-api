package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yosikez/item-api/config"
	"github.com/yosikez/item-api/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() ([]model.Item, error)
	FindByID(ID uint) (model.Item, error)
	Create(item model.Item) (model.Item, error)
	Update(item model.Item) (model.Item, error)
	Delete(ID uint) error
}

type itemRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewItemRepository() *itemRepository {
	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect to database")
	}

	redisClient := config.InitRedis()

	return &itemRepository{
		db:    db,
		cache: redisClient,
	}
}

func (repo *itemRepository) FindAll() ([]model.Item, error) {
	var items []model.Item
	var newData []byte

	data, err := repo.cache.Get(context.Background(), "items").Result()
	if err == nil {
		err = json.Unmarshal([]byte(data), &items)
		return items, err
	}

	err = repo.db.Find(&items).Error

	if err != nil {
		return items, err
	}
	
	newData, err = json.Marshal(items)
	if err == nil {
		err = repo.cache.Set(context.Background(), "items", string(newData), time.Minute*10).Err()
	}

	return items, err
}

func (repo *itemRepository) FindByID(ID uint) (model.Item, error) {
	var item model.Item
	var newData []byte

	data, err := repo.cache.Get(context.Background(), fmt.Sprintf("item:%v", ID)).Result()
	if err == nil {
		err = json.Unmarshal([]byte(data), &item)
		return item, err
	}

	err = repo.db.First(&item, ID).Error

	if err != nil {
		return item, err
	}

	newData, err = json.Marshal(item)
	if err == nil {
		err = repo.cache.Set(context.Background(), fmt.Sprintf("item:%d", ID), string(newData), time.Minute*10).Err()
	}

	return item, err
}

func (repo *itemRepository) Create(item model.Item) (model.Item, error) {
	err := repo.db.Create(&item).Error
	return item, err
}

func (repo *itemRepository) Update(item model.Item) (model.Item, error) {
	err := repo.db.Save(&item).Error
	return item, err
}
func (repo *itemRepository) Delete(ID uint) error {
	err := repo.db.Delete(&model.Item{}, ID).Error
	return err
}
