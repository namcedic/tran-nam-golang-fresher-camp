package categorybiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
)

type DeleteCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteCategoryBiz struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryBiz(store DeleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
