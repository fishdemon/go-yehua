package service

import (
	"github.com/fishdemon/go-yehua/grpc/proto"
	"golang.org/x/net/context"
)

type UserService struct {
	usersCache []*proto.User
}

func NewUserService() *UserService {
	return new(UserService)
}

func (us *UserService) AddUser(ctx context.Context, user *proto.User) (*proto.CommonResponse, error) {
	panic("airport-data")
	us.usersCache = append(us.usersCache, user)
	return &proto.CommonResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (us *UserService) GetUsers(userFilter *proto.UserFilter, stream proto.UserService_GetUsersServer) error {
	for _, user := range us.usersCache {
		if userFilter.Id != user.Id {
			continue
		}
		stream.Send(user)
	}

	return nil
}


