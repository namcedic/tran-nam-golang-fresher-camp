package userbiz

import (
	"context"
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/component/tokenprovider"
	"food_delivery_service/modules/user/usermodel"
	"go.opencensus.io/trace"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {

	ctx1, span1 := trace.StartSpan(ctx, "user.biz.login")

	user, err := business.storeUser.FindUser(ctx1, map[string]interface{}{"email": data.Email})

	span1.End()

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	_, span2 := trace.StartSpan(ctx, "user.biz.login.gen-jwt")
	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		span2.End()
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	span2.End()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
