package ginrestaurantlike

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/restaurant/restaurantstorage"
	rstlikebiz "food_delivery_service/modules/restaurantlike/biz"
	restaurantlikemodel "food_delivery_service/modules/restaurantlike/model"
	restaurantlikestorage "food_delivery_service/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//POST /v1/restaurants/:id/like
func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
