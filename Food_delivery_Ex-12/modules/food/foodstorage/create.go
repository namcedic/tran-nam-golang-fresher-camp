package foodstorage

import (
	"context"
	"food_delivery_service/modules/food/foodmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *foodmodel.FoodCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
