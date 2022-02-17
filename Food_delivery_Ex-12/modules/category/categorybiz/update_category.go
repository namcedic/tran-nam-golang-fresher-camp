package categorybiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
)

type UpdateCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *categorymodel.CategoryUpdate,
	) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryBiz(store UpdateCategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
