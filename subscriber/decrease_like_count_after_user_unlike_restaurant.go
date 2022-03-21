package subscriber

import (
	"context"
	"food_delivery_service/component"
	"food_delivery_service/modules/restaurant/restaurantstorage"
	"food_delivery_service/pubsub"
	"food_delivery_service/skio"
	"go.opencensus.io/trace"
)

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			ctx1, span := trace.StartSpan(ctx, "pubsub.sub.RunDecreaseLikeCountAfterUserUnlikeRestaurant")
			defer span.End()

			return store.DecreaseLikeCount(ctx1, likeData.GetRestaurantId())
		},
	}
}
