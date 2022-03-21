package rstlikebiz

import (
	"context"
	"food_delivery_service/common"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
	"food_delivery_service/pubsub"
	"go.opencensus.io/trace"
)

type UserLikeRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store  UserLikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore,
	pubsub pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, pubsub: pubsub}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	_, span := trace.StartSpan(ctx, "restaurant.biz.like")
	span.AddAttributes(
		trace.Int64Attribute("restaurant_id", int64(data.RestaurantId)),
		trace.Int64Attribute("user_id", int64(data.UserId)),
	)

	defer span.End()
	userliked, _ := biz.store.FindUserLikedRestaurant(ctx,
		map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId})

	if userliked != nil {
		return restaurantlikemodel.ErrUserLikeRestaurant
	}
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))
	return nil
}
