package Auth

import (
	"context"
	"fmt"
	"quocbang/go-interface/impl"
	"testing"

	fake "github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_GetUserInfo(t *testing.T) {
	assertion := assert.New(t)
	missingAuthKey := NewService(AuthKey{})
	s := NewService(AuthKey{Key: "test_auth_key"})
	ctx := context.Background()

	// case: missing Auth key.
	{

		_, actual := missingAuthKey.GetUserInfo(ctx, impl.GetUserInfoRequest{
			UserID: fake.Name(),
		})
		assertion.Error(actual)
		expected := fmt.Errorf("missing auth key")

		assertion.Equal(expected, actual)
	}
	// case: missing request.
	{
		_, actual := s.GetUserInfo(ctx, impl.GetUserInfoRequest{})
		assertion.Error(actual)
		expected := fmt.Errorf("missing request")

		assertion.Equal(expected, actual)
	}
	// case: get user info successfully.
	{
		reply, err := s.GetUserInfo(ctx, impl.GetUserInfoRequest{
			UserID: fake.Name(),
		})
		assertion.NoError(err)
		expected := impl.GetUserInfoReply{
			LastName:  "Bang",
			FisrtName: "Quoc",
			UserAge:   23,
		}
		assertion.Equal(expected, reply)
	}
}
