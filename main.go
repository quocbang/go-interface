package main

import (
	"context"
	"fmt"
	"log"
	"quocbang/go-interface/impl"
	user "quocbang/go-interface/impl/handlers/Auth"
	"quocbang/go-interface/impl/service"
)

type AuthKey struct {
	Key    string
	Values string
}

func Service(in user.AuthKey) service.Auth {
	return user.NewService(in)
}

func main() {
	AuthKey := AuthKey{
		Key:    "producttion",
		Values: "this is key to access for all service",
	}

	s := Service(user.AuthKey(AuthKey))

	userInfo, err := s.GetUserInfo(context.Background(), impl.GetUserInfoRequest{
		UserID: "quocbang",
	})
	if err != nil {
		log.Fatalf("get user failed with error: %v", err.Error())
	}

	fmt.Printf("the user name is %s %s and the user age is %d", userInfo.FisrtName, userInfo.LastName, userInfo.UserAge)
}
