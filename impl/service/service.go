package service

import (
	"context"
	"quocbang/go-interface/impl"
)

type Auth interface {
	GetUserInfo(context.Context, impl.GetUserInfoRequest) (impl.GetUserInfoReply, error)
}

type Service struct {
	auth Auth
}
