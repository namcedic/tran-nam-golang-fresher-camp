package gincategory

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/category/categorybiz"
	"food_delivery_service/modules/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewGetCategoryBiz(store)

		data, err := biz.GetCategory(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
