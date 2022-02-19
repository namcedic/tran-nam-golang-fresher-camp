package rstlikebiz

import (
	"context"
	"food_delivery_service/common"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
	"food_delivery_service/pubsub"
)

type UserUnLikeRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

//type DecreaseLikeCountStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userUnLikeRestaurantBiz struct {
	store  UserUnLikeRestaurantStore
	pubsub pubsub.Pubsub
	//decStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, pubsub pubsub.Pubsub) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store, pubsub: pubsub}
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
	biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(data))
	return nil
}
