package userlikerstbiz

import (
	"context"
	"food_delivery_service/common"
	userlikemodel "food_delivery_service/modules/userlike/model"
)

type ListRestaurantLikeByUserStore interface {
	GetRestaurantLikeByUser(ctx context.Context,
		conditions map[string]interface{},
		filter *userlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleRestaurant, error)
}

type listRestaurantLikeByUserBiz struct {
	store ListRestaurantLikeByUserStore
}

func NewListRestaurantLikeByUserBiz(store ListRestaurantLikeByUserStore) *listRestaurantLikeByUserBiz {
	return &listRestaurantLikeByUserBiz{store: store}
}

func (biz *listRestaurantLikeByUserBiz) ListRestaurants(
	ctx context.Context,
	filter *userlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleRestaurant, error) {
	restaurants, err := biz.store.GetRestaurantLikeByUser(ctx, nil, filter, paging, "Restaurant")

	if err != nil {
		return nil, common.ErrCannotListEntity(userlikemodel.EntityName, err)
	}

	return restaurants, nil
}
