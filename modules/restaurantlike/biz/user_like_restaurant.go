package rstlikebiz

import (
	"context"
	"food_delivery_service/common"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
	"food_delivery_service/pubsub"
)

type UserLikeRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}
type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore,
	//incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, pubsub: pubsub}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
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
	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//	})
	//
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	return nil
}
