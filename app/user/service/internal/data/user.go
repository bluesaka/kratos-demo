package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	orderv1 "kratos-demo/api/order/service/v1"
	"kratos-demo/app/user/service/internal/biz"
	"time"
)

type User struct {
	Id    int64
	Phone string
	Age   int32
	//CreateTime string
	CreateTime time.Time
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func convertUser(u User) *biz.User {
	return &biz.User{
		Phone:      u.Phone,
		Age:        u.Age,
		CreateTime: u.CreateTime.Format("2006-01-02 15:04:05"),
	}
}

func (r *userRepo) Get(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	err := r.data.bunDB.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return &biz.User{}, err
	}

	orderReply, err := r.data.orderClient.GetOrder(ctx, &orderv1.GetOrderReq{
		Id: 1,
	})
	fmt.Printf("orderReply: %+v, err: %v\n", orderReply, err)

	return convertUser(user), nil
}

func (r *userRepo) List(ctx context.Context) ([]*biz.User, error) {
	var users []User
	result := r.data.db.WithContext(ctx).
		Where("1=1").
		Limit(10).
		Offset(0).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rv := make([]*biz.User, 0)
	for _, v := range users {
		rv = append(rv, convertUser(v))
	}

	//if true {
	//	return nil, userapiv1.ErrorUserNotFound("user %s not found", "test")
	//}
	//
	//if true {
	//	return nil, errors.New(500, "EMPTY_USER", "用户未找到")
	//}

	return rv, nil
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	panic("implement me")
}
