package categorybiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
)

type GetCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type getCategoryBiz struct {
	store GetCategoryStore
}

func NewGetCategoryBiz(store GetCategoryStore) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategory(ctx context.Context, id int) (*categorymodel.Category, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	return data, err
}
