package foodbiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/food/foodmodel"
)

type DeleteFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*foodmodel.Food, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteFoodBiz struct {
	store DeleteFoodStore
}

func NewDeleteFoodBiz(store DeleteFoodStore) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) DeleteFood(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}
