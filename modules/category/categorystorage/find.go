package categorystorage

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/modules/category/categorymodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*categorymodel.Category, error) {
	var result categorymodel.Category

	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}