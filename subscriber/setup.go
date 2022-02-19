package subscriber

import (
	"context"
	"food_delivery_service/component"
)

func Setup(ctx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(ctx, context.Background())
}
