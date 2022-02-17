package userlikestorage

import (
	"context"
	"fmt"
	"food_delivery_service/common"
	userlikemodel "food_delivery_service/modules/userlike/model"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestaurantLikeByUser(ctx context.Context,
	conditions map[string]interface{},
	filter *userlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleRestaurant, error) {
	var result []userlikemodel.UserLike

	db := s.db

	db = db.Table(userlikemodel.UserLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("Restaurant")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	restaurants := make([]common.SimpleRestaurant, len(result))

	for i, item := range result {
		result[i].Restaurant.CreatedAt = item.CreatedAt
		result[i].Restaurant.UpdatedAt = nil
		restaurants[i] = *result[i].Restaurant

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}
	return restaurants, nil
}
