package foodstorage

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/food/foodmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *foodmodel.FoodUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
