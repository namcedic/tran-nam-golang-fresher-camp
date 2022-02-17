package restaurantlikestorage

import (
	"context"
	"food_delivery_service/common"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUserLikedRestaurant(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*restaurantlikemodel.Like, error) {
	db := s.db.Table(restaurantlikemodel.Like{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var userLiked restaurantlikemodel.Like

	if err := db.Where(conditions).First(&userLiked).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &userLiked, nil
}
