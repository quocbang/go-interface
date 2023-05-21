package Auth

import (
	"context"
	"fmt"
	"quocbang/go-interface/impl"
	"quocbang/go-interface/impl/service"
)

type AuthKey struct {
	Key    string
	Values string
}

func NewService(ak AuthKey) service.Auth {
	return ak
}

func (a AuthKey) GetUserInfo(ctx context.Context, req impl.GetUserInfoRequest) (impl.GetUserInfoReply, error) {
	if a.Key == "" {
		return impl.GetUserInfoReply{}, fmt.Errorf("missing auth key")
	}
	if req.UserID == "" {
		return impl.GetUserInfoReply{}, fmt.Errorf("missing request")
	}
	return impl.GetUserInfoReply{
		LastName:  "Bang",
		FisrtName: "Quoc",
		UserAge:   23,
	}, nil
}
