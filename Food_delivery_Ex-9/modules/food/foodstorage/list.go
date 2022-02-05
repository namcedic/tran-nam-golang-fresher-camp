package foodstorage

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/food/foodmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]foodmodel.Food, error) {
	var result []foodmodel.Food

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(foodmodel.Food{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
