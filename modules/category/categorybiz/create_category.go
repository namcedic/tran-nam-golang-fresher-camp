package categorybiz

import (
	"context"
	"food_delivery_service/modules/category/categorymodel"
)

type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)
	return err
}
