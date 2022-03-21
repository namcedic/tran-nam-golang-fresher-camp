package subscriber

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/restaurant/restaurantstorage"
	"food_delivery_service/pubsub"
	"food_delivery_service/skio"
	"go.opencensus.io/trace"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			ctx1, span := trace.StartSpan(ctx, "pubsub.sub.RunIncreaseLikeCountAfterUserLikeRestaurant")
			defer span.End()

			return store.IncreaseLikeCount(ctx1, likeData.GetRestaurantId())
		},
	}
}

func EmitRealtimeAfterUserLikeRestaurant(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			return rtEngine.EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
		},
	}
}
