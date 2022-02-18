package rstlikebiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/component/asyncjob"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
)

type UserUnLikeRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}
type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikeRestaurantBiz struct {
	store    UserUnLikeRestaurantStore
	decStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, decStore DecreaseLikeCountStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store, decStore: decStore}
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

	go func() {
		defer common.AppRecover()

		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)

	}()

	return nil
}
