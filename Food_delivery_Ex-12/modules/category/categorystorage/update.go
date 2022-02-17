package categorystorage

import (
	"context"

	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *categorymodel.CategoryUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
