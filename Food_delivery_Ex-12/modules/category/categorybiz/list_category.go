package categorybiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
)

type ListCategoryStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStore
}

func NewListCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(
	ctx context.Context,
	paging *common.Paging,
) ([]categorymodel.Category, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}
	return result, nil
}
