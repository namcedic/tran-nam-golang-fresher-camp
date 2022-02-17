package ginuserlike

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	userlikerstbiz "food_delivery_service/modules/userlike/biz"
	userlikemodel "food_delivery_service/modules/userlike/model"
	userlikestorage "food_delivery_service/modules/userlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := userlikemodel.Filter{
			UserId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := userlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userlikerstbiz.NewListRestaurantLikeByUserBiz(store)

		result, err := biz.ListRestaurants(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
