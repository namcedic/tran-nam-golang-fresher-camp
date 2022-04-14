package foodmodel

type Filter struct {
	RestaurantId int `json:"restaurant_id,omitempty" form:"restaurant_id"`
}
