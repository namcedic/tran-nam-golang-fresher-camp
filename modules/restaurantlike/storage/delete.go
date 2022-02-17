package restaurantlikestorage

import (
	"context"
	"food_delivery_service/common"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("restaurant_id =? and user_id =?", data.RestaurantId, data.UserId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
