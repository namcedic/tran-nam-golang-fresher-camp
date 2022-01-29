package foodbiz

import (
	"context"
	"errors"
	"food_delivery_service/modules/food/foodmodel"
)

type CreateFoodStore interface {
	Create(ctx context.Context, data *foodmodel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error {
	if data.Name == "" {
		return errors.New("food name can not be blank")
	}
	err := biz.store.Create(ctx, data)

	return err
}
