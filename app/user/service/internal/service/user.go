package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/user/service/v1"
	"kratos-demo/app/user/service/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return &v1.GetUserReply{}, err
	}

	return &v1.GetUserReply{
		User: &v1.GetUserReply_User{
			Phone:      user.Phone,
			Age:        user.Age,
			CreateTime: user.CreateTime,
		},
	}, nil
}

func (s *UserService) List(ctx context.Context, req *v1.UserListRequest) (*v1.UserListReply, error) {
	users, err := s.uc.List(ctx)
	if err != nil {
		return &v1.UserListReply{}, err
	}

	list := make([]*v1.UserListReply_User, 0)
	for _, v := range users {
		list = append(list, &v1.UserListReply_User{
			Phone:      v.Phone,
			Age:        v.Age,
			CreateTime: v.CreateTime,
		})
	}

	return &v1.UserListReply{
		List: list,
	}, nil
}

//Update(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
