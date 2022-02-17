package rstlikebiz

import (
	"context"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
)

type UserUnLikeRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	userliked, _ := biz.store.FindUserLikedRestaurant(ctx,
		map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId})

	if userliked == nil {
		return restaurantlikemodel.ErrUserNotLikeRestaurant
	}
	err := biz.store.Delete(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotUnLikeRestaurant(err)
	}
	return nil
}
