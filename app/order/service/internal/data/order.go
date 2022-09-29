package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/app/order/service/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

type Order struct {
	Id     int64
	UserId int64
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "order-service/orderRepo")),
	}
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	order := Order{}
	result := r.data.db.WithContext(ctx).First(&order, id)
	return &biz.Order{
		Id:     order.Id,
		UserId: order.UserId,
	}, result.Error
}
