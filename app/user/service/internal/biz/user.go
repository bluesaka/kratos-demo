package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 领域对象
type User struct {
	Phone      string
	Age        int32
	CreateTime string
}

type UserRepo interface {
	Get(ctx context.Context, id int64) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (user *User, err error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context) (user []*User, err error) {
	return uc.repo.List(ctx)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (user *User, err error) {
	return uc.repo.Update(ctx, u)
}
